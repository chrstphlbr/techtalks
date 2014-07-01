package data

import (
	"log"
	"testing"
)

func TestImportVideo(t *testing.T) {
	const dir = "/Users/Christoph/tmp/techtalks_data/"

	i := NewFileImporter(dir)
	i.SetImportFunc(func(d interface{}) {
		switch d.(type) {
		case fileVideo:
			log.Println("fileVideo")
		case fileEvent:
			log.Println("fileEvent")
		case fileSpeaker:
			log.Println("fileSpeaker")
		case fileEventSeries:
			log.Println("fileEventSeries")
		default:
			log.Println("unknown type")
		}
	})
	i.ImportData()
}
