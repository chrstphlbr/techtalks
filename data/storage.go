package data

import (
	"log"
)

type Storer interface {
	Store(data interface{})
}

type StoreFunc func(data interface{})

func (sf StoreFunc) Store(data interface{}) {
	sf(data)
}

type InMemoryStorer struct {
	videos      map[VideoId]*Video
	speakers    map[SpeakerId]*Speaker
	events      map[EventId]*Event
	eventSeries map[EventSeriesId]*EventSeries
}

func NewInMemoryStorer() *InMemoryStorer {
	return &InMemoryStorer{
		videos:      make(map[VideoId]*Video),
		speakers:    make(map[SpeakerId]*Speaker),
		events:      make(map[EventId]*Event),
		eventSeries: make(map[EventSeriesId]*EventSeries),
	}
}

func (m *InMemoryStorer) Store(data interface{}) {
	//TODO: convert data from file* to actual entities
	switch data.(type) {
	case fileVideo:
	case fileEvent:
		log.Println("fileEvent")
	case fileSpeaker:
		log.Println("fileSpeaker")
	case fileEventSeries:
		log.Println("fileEventSeries")
	default:
		log.Println("unknown type")
	}
}
