package data

import (
	"time"
)

type VideoId string

func (v VideoId) String() string {
	return string(v)
}

type SpeakerId string

func (s SpeakerId) String() string {
	return string(s)
}

type EventId string

func (e EventId) String() string {
	return string(e)
}

type EventSeriesId string

func (es EventSeriesId) String() string {
	return string(es)
}

type Video struct {
	id         VideoId
	Name       string
	Url        string
	Speakers   []SpeakerId
	Event      EventId
	RecordedAt time.Time
	Duration   time.Duration
	Poster     string
	Lang       string
}

type Speaker struct {
	id        SpeakerId
	FirstName string
	LastName  string
	Videos    []VideoId
}

type Event struct {
	id          string
	Name        string
	EventSeries EventSeriesId
	StartDate   time.Time
	EndDate     time.Time
	Videos      []VideoId
}

type EventSeries struct {
	id     string
	Name   string
	Events []EventId
}
