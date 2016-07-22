package memory

import (
	"github.com/dmashuda/calendar/model"
	"strconv"
	"sync/atomic"
)

var lastId uint64 = 0

type EventRepo struct {
	events       map[string]model.Event
	participants map[string]model.Participant
}

func NewEventRepo() EventRepo {
	return EventRepo{
		events:       map[string]model.Event{},
		participants: map[string]model.Participant{},
	}
}

func (r EventRepo) Store(m model.Event) (string, error) {
	if m.ID == "" {
		m.ID = strconv.Itoa(int(lastId))
		atomic.AddUint64(&lastId, 1)
	}
	r.events[m.ID] = m
	return m.ID, nil
}

func (r EventRepo) Destroy(s string) (string, error) {
	delete(r.events, s)
	return s, nil
}

func (r EventRepo) Load(s string) (model.Event, error) {
	return r.events[s], nil
}
