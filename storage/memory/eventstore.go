package memory

import (
	"github.com/dmashuda/calendar/model"
)

type EventRepo struct {
	events       map[string]model.Event
	participants map[string]model.Participant
}

func (r EventRepo) Store(m model.Event) error {
	return nil
}
