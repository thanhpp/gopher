package gcal

import "time"

type Event struct {
	ID    string
	CalID string
	Name  string
	Start time.Time
	End   time.Time
	Desc  string
}

type Calendar struct {
	ID   string
	Name string
}
