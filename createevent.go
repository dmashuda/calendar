package calendar

import (
	"github.com/dmashuda/calendar/model"
	"github.com/dmashuda/calendar/pattern"
	"time"
)

var _ pattern.Command = CreateEvent{}

type EventRepo interface {
	Store(model.Event) error
}

type CreateEventRequest struct {
	Start, End   time.Time
	Color, Title string
}

type CreateEvent struct {
	repo    EventRepo
	request *CreateEventRequest
}

func NewCreateEvent(ev EventRepo, req *CreateEventRequest) CreateEvent {
	return CreateEvent{
		repo:    ev,
		request: req,
	}
}

func (c CreateEvent) Execute() error {
	return nil
}
