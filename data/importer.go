package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type Importer interface {
	SetStoreFunc(fn StoreFunc)
	ImportData()
}

type typeDir int

func (d typeDir) String() string {
	switch d {
	case VideoDir:
		return "videos"
	case SpeakerDir:
		return "speakers"
	case EventDir:
		return "events"
	case EventSeriesDir:
		return "event_series"
	}
	return ""
}

const (
	VideoDir typeDir = iota
	SpeakerDir
	EventDir
	EventSeriesDir
)

func typeDirPath(basePath string, td typeDir) (p string) {
	p = basePath
	if strings.HasSuffix(basePath, "/") {
		p = p + td.String()
	} else {
		p = p + "/" + td.String()
	}
	return
}

func imp(r io.Reader, td typeDir, importFn StoreFunc) (err error) {
	dec := json.NewDecoder(r)

	var value interface{}

	switch td {
	case VideoDir:
		var v fileVideo
		err = dec.Decode(&v)
		value = v
	case SpeakerDir:
		var v fileSpeaker
		err = dec.Decode(&v)
		value = v
	case EventDir:
		var v fileEvent
		err = dec.Decode(&v)
		value = v
	case EventSeriesDir:
		var v fileEventSeries
		err = dec.Decode(&v)
		value = v
	default:
		value = nil
		err = fmt.Errorf("unknown typeDir")
	}

	if err == nil {
		// log.Println(v)
		importFn(value)
	} else {
		log.Printf("decode err: %s\n", err)
	}
	return
}
