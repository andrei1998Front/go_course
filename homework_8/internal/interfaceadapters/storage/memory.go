package memory

import (
	"fmt"
	"sync"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/google/uuid"
)

type Repo struct {
	mutex  sync.RWMutex
	events map[string]event.Event
}

func NewRepo() *Repo {
	events := make(map[string]event.Event)
	return &Repo{events: events}
}

func (r *Repo) GetByID(id uuid.UUID) (*event.Event, error) {
	const op = "interfaceadapters.storage.GetByID"

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	e, ok := r.events[id.String()]

	if !ok {
		return &event.Event{}, fmt.Errorf("%s: %w", op, event.ErrNonExistentEvent)
	}

	return &e, nil
}

func (r *Repo) GetAll() ([]event.Event, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var eventSlice []event.Event

	for _, val := range r.events {
		eventSlice = append(eventSlice, val)
	}

	return eventSlice, nil
}

func (r *Repo) GetByDate(dt time.Time) (*event.Event, error) {
	const op = "interfaceadapters.storage.GetByDate"

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, val := range r.events {
		if val.Date == dt {
			return &val, nil
		}
	}

	return &event.Event{}, fmt.Errorf("%s: %w", op, event.ErrNonExistentDate)
}

func (r *Repo) Add(e event.Event) error {
	const op = "interfaceadapters.storage.Add"

	_, err := r.GetByID(e.ID)

	if err == nil {
		return fmt.Errorf("%s: %w", op, event.ErrExistentID)
	}

	_, err = r.GetByDate(e.Date)

	if err == nil {
		return fmt.Errorf("%s: %w", op, event.ErrDateBusy)
	}

	r.mutex.Lock()
	r.events[e.ID.String()] = e
	r.mutex.Unlock()

	return nil
}

func (r *Repo) Delete(id uuid.UUID) error {
	_, err := r.GetByID(id)

	if err != nil {
		return err
	}

	r.mutex.Lock()
	delete(r.events, id.String())
	r.mutex.Unlock()

	return nil
}

func (r *Repo) Update(e event.Event) error {
	const op = "interfaceadapters.storage.Update"

	_, err := r.GetByID(e.ID)

	if err != nil {
		return err
	}

	if eByDate, err := r.GetByDate(e.Date); eByDate.ID != e.ID && err == nil {
		return fmt.Errorf("%s: %w", op, event.ErrDateBusy)
	}

	r.mutex.Lock()
	r.events[e.ID.String()] = e
	r.mutex.Unlock()

	return nil
}
