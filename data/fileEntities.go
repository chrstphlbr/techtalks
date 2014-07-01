package data

type fileVideo struct {
	Id         string
	Name       string
	Url        string
	Speakers   []string
	Tags       []string
	Event      string
	RecordedAt string `json:"recorded_at"`
	Duration   string
	UploadedAt string `json:"uploaded_at"`
	Poster     string
	Lang       string
}

type fileSpeaker struct {
	Id        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type fileEvent struct {
	Id          string
	Name        string
	EventSeries string `json:"event_series"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type fileEventSeries struct {
	Id   string
	Name string
}
