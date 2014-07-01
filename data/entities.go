package data

import (
	"time"
)

type Video struct {
	id         string
	Name       string
	Url        string
	Speakers   []Speaker
	Event      Event
	RecordedAt time.Time
	Duration   time.Duration
	Poster     string
	Lang       string
}

type Speaker struct {
	id        string
	FirstName string
	LastName  string
	Videos    []Video
}

type Event struct {
	id          string
	Name        string
	EventSeries EventSeries
	StartDate   time.Time
	EndDate     time.Time
	Videos      []Video
}

type EventSeries struct {
	id     string
	Name   string
	Events []Event
}
