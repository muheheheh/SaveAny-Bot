package batchtfile

import (
	"context"

	"github.com/krau/SaveAny-Bot/pkg/storagetypes"
	"github.com/krau/SaveAny-Bot/storage"
)

func (t *Task) reuseMedia(ctx context.Context, reuser storage.StorageTelegramReuser, elems []*TaskElement) error {
	items := make([]storagetypes.TelegramItem, len(elems))
	for i, elem := range elems {
		items[i] = storagetypes.TelegramItem{File: elem.File, StoragePath: elem.Path}
		t.updateItem(elem.ID, func(item *itemProgressState) {
			item.phase = ItemPhaseReusing
		})
	}
	t.notifyStateChange(ctx)
	err := reuser.SaveTelegramMedia(ctx, items, func(index int) {
		t.markItemCompleted(elems[index].ID)
	})
	if err != nil {
		for _, elem := range elems {
			// markItemFailed preserves already confirmed items on partial failure.
			t.markItemFailed(elem.ID, FailureStageConfirm, err)
		}
	}
	t.notifyStateChange(ctx)
	return err
}
