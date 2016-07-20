package model

import "time"

type Event struct {
	ID, Title, Color   string
	StartDate, EndDate time.Time
	Participants       []Participant
}
