package storagetypes

import "github.com/krau/SaveAny-Bot/pkg/tfile"

// TelegramItem retains the source media reference instead of a file reader.
type TelegramItem struct {
	File        tfile.TGFile
	StoragePath string
}
