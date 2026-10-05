package handlers

import (
	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/ext"

	"github.com/krau/SaveAny-Bot/client/bot/handlers/utils/mediautil"
	"github.com/krau/SaveAny-Bot/database"
	"github.com/krau/SaveAny-Bot/pkg/tfile"
)

func handleMediaMessage(ctx *ext.Context, update *ext.Update) error {
	message := update.EffectiveMessage.Message
	if !mediautil.IsSupported(message.Media) {
		return dispatcher.EndGroups
	}
	if groupID, grouped := message.GetGroupedID(); grouped && groupID != 0 {
		return handleGroupMediaMessage(ctx, update, message, groupID)
	}
	user, err := database.GetUserByChatID(ctx, update.GetUserChat().GetID())
	if err != nil {
		return err
	}
	file, err := tfile.FromMediaMessage(message.Media, ctx.Raw, message, mediautil.TfileOptions(ctx, user, message)...)
	if err != nil {
		return reportRelayError(ctx, update, 0, err)
	}
	return processTelegramFiles(ctx, update, []tfile.TGFileMessage{file}, 0, nil)
}
