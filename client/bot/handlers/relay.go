package handlers

import (
	"errors"
	"fmt"
	"time"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/ext"
	"github.com/charmbracelet/log"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/tg"

	"github.com/krau/SaveAny-Bot/client/bot/handlers/utils/dirutil"
	"github.com/krau/SaveAny-Bot/client/bot/handlers/utils/msgelem"
	"github.com/krau/SaveAny-Bot/client/bot/handlers/utils/shortcut"
	"github.com/krau/SaveAny-Bot/common/i18n"
	"github.com/krau/SaveAny-Bot/common/i18n/i18nk"
	"github.com/krau/SaveAny-Bot/common/utils/tgutil"
	"github.com/krau/SaveAny-Bot/pkg/consts/tglimit"
	"github.com/krau/SaveAny-Bot/pkg/tcbdata"
	"github.com/krau/SaveAny-Bot/pkg/tfile"
	"github.com/krau/SaveAny-Bot/storage"
)

// Storage destinations never determine the relay target: it is this request's chat.
func processTelegramFiles(ctx *ext.Context, update *ext.Update, files []tfile.TGFileMessage, progressID int, fetchErr error) error {
	chatID := update.EffectiveChat().GetID()
	relayErr := errors.Join(fetchErr, relayTelegramFiles(ctx, chatID, files))
	stors := storage.GetUserStorages(ctx, update.GetUserChat().GetID())
	if relayErr != nil {
		errorMessageID := 0
		if len(stors) == 0 {
			errorMessageID = progressID
		}
		if err := reportRelayError(ctx, update, errorMessageID, relayErr); err != nil && !errors.Is(err, dispatcher.EndGroups) {
			log.FromContext(ctx).Errorf("Failed to report relay error: %v", err)
		}
	}
	if len(stors) == 0 {
		if relayErr == nil && progressID != 0 {
			if err := ctx.DeleteMessages(chatID, []int{progressID}); err != nil {
				return fmt.Errorf("delete relay progress message: %w", err)
			}
		}
		return dispatcher.EndGroups
	}
	if progressID == 0 {
		msg, err := ctx.Reply(update, ext.ReplyTextString(i18n.T(i18nk.BotMsgMediaGroupInfoSavingFiles)), nil)
		if err != nil {
			return fmt.Errorf("create save progress message: %w", err)
		}
		progressID = msg.ID
	}
	selectStorage := func(ctx *ext.Context, update *ext.Update) error {
		if len(files) == 1 {
			req, err := msgelem.BuildAddOneSelectStorageMessage(ctx, stors, files[0], progressID)
			if err != nil {
				return fmt.Errorf("build storage selection: %w", err)
			}
			if _, err := ctx.EditMessage(chatID, req); err != nil {
				return fmt.Errorf("show storage selection: %w", err)
			}
			return dispatcher.EndGroups
		}
		markup, err := msgelem.BuildAddSelectStorageKeyboard(stors, tcbdata.Add{Files: files, AsBatch: true})
		if err != nil {
			return fmt.Errorf("build storage selection: %w", err)
		}
		_, err = ctx.EditMessage(chatID, &tg.MessagesEditMessageRequest{
			ID:          progressID,
			Message:     i18n.T(i18nk.BotMsgCommonInfoFoundFilesSelectStorage, map[string]any{"Count": len(files)}),
			ReplyMarkup: markup,
		})
		if err != nil {
			return fmt.Errorf("show storage selection: %w", err)
		}
		return dispatcher.EndGroups
	}
	save := func(ctx *ext.Context, update *ext.Update) error {
		userID := update.GetUserChat().GetID()
		stor := storage.FromContext(ctx)
		dir := dirutil.PathFromContext(ctx)
		if len(files) == 1 {
			return shortcut.CreateAndAddTGFileTaskWithEdit(ctx, userID, stor, dir, files[0], progressID)
		}
		return shortcut.CreateAndAddBatchTGFileTaskWithEdit(ctx, userID, stor, dir, files, progressID)
	}
	return handleSilentMode(selectStorage, save)(ctx, update)
}

func reportRelayError(ctx *ext.Context, update *ext.Update, progressID int, relayErr error) error {
	log.FromContext(ctx).Errorf("Telegram relay failed: %v", relayErr)
	text := i18n.T(i18nk.BotMsgRelayErrorFailed, map[string]any{"Error": relayErr.Error()})
	var err error
	if progressID != 0 {
		_, err = ctx.EditMessage(update.EffectiveChat().GetID(), &tg.MessagesEditMessageRequest{ID: progressID, Message: text})
	} else {
		_, err = ctx.Reply(update, ext.ReplyTextString(text), nil)
	}
	if err != nil {
		return fmt.Errorf("send relay error: %w", err)
	}
	return dispatcher.EndGroups
}

func relayTelegramFiles(ctx *ext.Context, chatID int64, files []tfile.TGFileMessage) error {
	peer := ctx.PeerStorage.GetInputPeerById(chatID)
	if peer == nil || peer.Zero() {
		return fmt.Errorf("resolve relay chat %d", chatID)
	}
	media := make([]message.MultiMediaOption, len(files))
	for i, file := range files {
		caption := styling.Plain(file.Message().GetMessage())
		switch location := file.Location().(type) {
		case *tg.InputPhotoFileLocation:
			media[i] = message.Photo(location, caption)
		case *tg.InputDocumentFileLocation:
			media[i] = message.Document(location, caption)
		default:
			return fmt.Errorf("unsupported Telegram media reference: %T", location)
		}
	}
	started := time.Now()
	for start := 0; start < len(files); {
		if err := ctx.Err(); err != nil {
			return err
		}
		end := start + 1
		groupID := files[start].Message().GroupedID
		sourceChat := tgutil.ChatIdFromPeer(files[start].Message().PeerID)
		for groupID != 0 && end < len(files) && end-start < tglimit.MaxAlbumItems &&
			files[end].Message().GroupedID == groupID && tgutil.ChatIdFromPeer(files[end].Message().PeerID) == sourceChat {
			end++
		}
		if _, err := ctx.Sender.To(peer).Album(ctx, media[start], media[start+1:end]...); err != nil {
			return fmt.Errorf("relay media %d-%d: %w", start+1, end, err)
		}
		start = end
	}
	log.FromContext(ctx).Info("Telegram media relayed", "items", len(files), "elapsed", time.Since(started))
	return nil
}
