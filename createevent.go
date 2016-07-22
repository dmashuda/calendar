package calendar

import (
	"errors"
	"github.com/dmashuda/calendar/model"
	"github.com/dmashuda/calendar/pattern"
	"time"
)

var _ pattern.Command = CreateEvent{}

type EventRepo interface {
	Store(model.Event) (string, error)
	Destroy(string) (string, error)
	Load(string) (model.Event, error)
}

type CreateEventRequest struct {
	Start, End   time.Time
	Color, Title string
}
type CreateEventResponse struct {
	ID string
}

type CreateEvent struct {
	repo     EventRepo
	request  CreateEventRequest
	response *CreateEventResponse
	eventId  *string
}

func NewCreateEvent(ev EventRepo, req CreateEventRequest, resp *CreateEventResponse) CreateEvent {
	return CreateEvent{
		repo:     ev,
		request:  req,
		response: resp,
	}
}

func (c CreateEvent) Execute() error {
	req := c.request

	if req.Start.After(req.End) {
		return errors.New("Start time must be before end time")
	}
	if req.Title == "" {
		return errors.New("Event must have title")
	}

	id, err := c.repo.Store(model.Event{
		Title:        req.Title,
		Color:        req.Color,
		StartDate:    req.Start,
		EndDate:      req.End,
		Participants: []model.Participant{},
	})

	if err != nil {
		return err
	}

	c.eventId = &id

	c.response.ID = id

	return nil
}

func (c CreateEvent) Undo() error {
	return nil
}
