package memory

import (
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/google/uuid"
)

type Repo struct {
	events map[string]event.Event
}

func NewRepo() Repo {
	events := make(map[string]event.Event)
	return Repo{events: events}
}

func (r Repo) GetByID(id uuid.UUID) (*event.Event, error) {
	e, ok := r.events[id.String()]

	if !ok {
		return &event.Event{}, event.ErrNonExistentEvent
	}

	return &e, nil
}

func (r Repo) GetAll() ([]event.Event, error) {
	var eventSlice []event.Event

	for _, val := range r.events {
		eventSlice = append(eventSlice, val)
	}

	return eventSlice, nil
}

func (r Repo) GetByDate(dt time.Time) (*event.Event, error) {
	for _, val := range r.events {
		if val.Date == dt {
			return &val, nil
		}
	}

	return &event.Event{}, event.ErrNonExistentDate
}

func (r Repo) Add(e event.Event) error {
	_, err := r.GetByID(e.ID)

	if err == nil {
		return event.ErrExistentID
	}

	_, err = r.GetByDate(e.Date)

	if err == nil {
		return event.ErrDateBusy
	}

	r.events[e.ID.String()] = e

	return nil
}

func (r Repo) Delete(e event.Event) error {
	_, err := r.GetByID(e.ID)

	if err != nil {
		return err
	}

	delete(r.events, e.ID.String())

	return nil
}

func (r Repo) Update(e event.Event) error {
	_, err := r.GetByID(e.ID)

	if err != nil {
		return err
	}

	_, err = r.GetByDate(e.Date)

	if err == nil {
		return event.ErrDateBusy
	}

	r.events[e.ID.String()] = e

	return nil
}
