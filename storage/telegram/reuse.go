package telegram

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/tg"

	"github.com/krau/SaveAny-Bot/common/utils/tgutil"
	"github.com/krau/SaveAny-Bot/pkg/consts/tglimit"
	"github.com/krau/SaveAny-Bot/pkg/storagetypes"
	"github.com/krau/SaveAny-Bot/pkg/tfile"
)

func (t *Telegram) ReuseTelegramMedia() bool {
	return t.config.ReuseMedia
}

// SaveTelegramMedia reuses photo/document references in Telegram. No file bytes
// are read, uploaded, renamed, converted or split. Errors are returned directly
// so a failed reuse never silently starts a download.
func (t *Telegram) SaveTelegramMedia(ctx context.Context, items []storagetypes.TelegramItem, onSent func(int)) error {
	tctx := tgutil.ExtFromContext(ctx)
	if tctx == nil {
		return fmt.Errorf("failed to get telegram context")
	}
	media := make([]message.MultiMediaOption, len(items))
	chatIDs := make([]int64, len(items))
	for i, item := range items {
		filename, chatID := t.target(tctx, path.Clean(item.StoragePath))
		chatIDs[i] = chatID
		caption := sourceCaptionOverride(ctx)
		if source, ok := item.File.(tfile.TGFileMessage); ok && source.Message() != nil {
			text := source.Message().GetMessage()
			caption = &text
		}
		styled := mediaCaption(filename, caption)
		switch location := item.File.Location().(type) {
		case *tg.InputPhotoFileLocation:
			media[i] = message.Photo(location, styled...)
		case *tg.InputDocumentFileLocation:
			media[i] = message.Document(location, styled...)
		default:
			return fmt.Errorf("unsupported Telegram media reference: %T", location)
		}
	}

	started := time.Now()
	logger := log.FromContext(ctx)
	logger.Info("Reusing Telegram media", "items", len(items))
	for start := 0; start < len(items); {
		end := start + 1
		for end < len(items) && end-start < tglimit.MaxAlbumItems && chatIDs[end] == chatIDs[start] {
			end++
		}
		peer := tryGetInputPeer(tctx, chatIDs[start])
		if peer == nil || peer.Zero() {
			return fmt.Errorf("failed to get input peer for chat ID %d", chatIDs[start])
		}
		if err := t.limiter.Wait(ctx); err != nil {
			return fmt.Errorf("rate limit failed: %w", err)
		}
		// Album also handles a single item, issuing messages.sendMedia instead.
		if _, err := tctx.Sender.To(peer).Album(ctx, media[start], media[start+1:end]...); err != nil {
			return fmt.Errorf("failed to reuse Telegram media: %w", err)
		}
		if onSent != nil {
			for i := start; i < end; i++ {
				onSent(i)
			}
		}
		start = end
	}
	logger.Info("Telegram media reused", "items", len(items), "elapsed", time.Since(started))
	return nil
}
