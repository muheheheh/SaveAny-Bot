package handlers

import (
	"cmp"
	"errors"
	"slices"
	"sync"
	"time"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/ext"
	"github.com/charmbracelet/log"
	"github.com/gotd/td/tg"
	"github.com/krau/SaveAny-Bot/client/bot/handlers/utils/mediautil"
	"github.com/krau/SaveAny-Bot/config"
	"github.com/krau/SaveAny-Bot/database"
	"github.com/krau/SaveAny-Bot/pkg/tfile"
)

// mediaGroupKey uniquely identifies a media group by chat, sender, and group
// ID so files from different users in the same chat can never be mixed.
type mediaGroupKey struct {
	chatID  int64
	userID  int64
	groupID int64
}

type MediaGroupHandler struct {
	groups    map[mediaGroupKey][]tfile.TGFileMessage
	timers    map[mediaGroupKey]*time.Timer
	mu        sync.Mutex
	timeout   time.Duration
	setupOnce sync.Once
}

func (m *MediaGroupHandler) SetupTimeout(timeoutSec int) {
	m.setupOnce.Do(func() {
		if timeoutSec < 1 {
			timeoutSec = 1
		}
		m.timeout = time.Duration(timeoutSec) * time.Second
	})
}

var (
	mediaGroupHandler = &MediaGroupHandler{
		groups: make(map[mediaGroupKey][]tfile.TGFileMessage),
		timers: make(map[mediaGroupKey]*time.Timer),
		mu:     sync.Mutex{},
	}
)

func handleGroupMediaMessage(ctx *ext.Context, update *ext.Update, message *tg.Message, groupID int64) error {
	mediaGroupHandler.SetupTimeout(max(config.C().Telegram.MediaGroupTimeout, 1))
	media := message.Media
	supported := mediautil.IsSupported(media)
	if !supported {
		return dispatcher.EndGroups
	}
	userId := update.GetUserChat().GetID()
	userDB, err := database.GetUserByChatID(ctx, userId)
	if err != nil {
		return err
	}
	tfOpts := mediautil.TfileOptions(ctx, userDB, message)
	file, err := tfile.FromMediaMessage(media, ctx.Raw, message, tfOpts...)
	if err != nil {
		return reportRelayError(ctx, update, 0, err)
	}
	mediaGroupHandler.mu.Lock()
	defer mediaGroupHandler.mu.Unlock()
	key := mediaGroupKey{
		chatID:  update.EffectiveChat().GetID(),
		userID:  userId,
		groupID: groupID,
	}
	if mediaGroupHandler.groups[key] == nil {
		mediaGroupHandler.groups[key] = make([]tfile.TGFileMessage, 0)
	}
	mediaGroupHandler.groups[key] = append(mediaGroupHandler.groups[key], file)

	if timer, exists := mediaGroupHandler.timers[key]; exists {
		timer.Stop()
	}
	mediaGroupHandler.timers[key] = time.AfterFunc(mediaGroupHandler.timeout, func() {
		processMediaGroup(ctx, update, key)
	})
	return dispatcher.EndGroups
}

func processMediaGroup(ctx *ext.Context, update *ext.Update, key mediaGroupKey) {
	logger := log.FromContext(ctx)
	mediaGroupHandler.mu.Lock()
	items := mediaGroupHandler.groups[key]
	delete(mediaGroupHandler.groups, key)
	delete(mediaGroupHandler.timers, key)
	mediaGroupHandler.mu.Unlock()
	if len(items) == 0 {
		logger.Warn("No media items to process for group", "groupID", key.groupID)
		return
	}
	logger.Debugf("Processing media group %d with %d items", key.groupID, len(items))

	slices.SortFunc(items, func(a, b tfile.TGFileMessage) int {
		return cmp.Compare(a.Message().GetID(), b.Message().GetID())
	})
	if err := processTelegramFiles(ctx, update, items, 0, nil); err != nil && !errors.Is(err, dispatcher.EndGroups) {
		logger.Errorf("Failed to process media group: %v", err)
	}
}
