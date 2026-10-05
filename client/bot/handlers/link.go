package handlers

import (
	"github.com/celestix/gotgproto/ext"

	"github.com/krau/SaveAny-Bot/client/bot/handlers/utils/shortcut"
)

func handleMessageLink(ctx *ext.Context, update *ext.Update) error {
	replied, files, _, err := shortcut.GetFilesFromUpdateLinkMessageWithReplyEdit(ctx, update)
	if len(files) == 0 {
		return err
	}
	return processTelegramFiles(ctx, update, files, replied.ID, err)
}
