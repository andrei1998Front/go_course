package event

import (
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	GetByID(id uuid.UUID) (*Event, error)
	GetByDate(dt time.Time) (*Event, error)
	GetAll() ([]Event, error)
	Add(e Event) error
	Delete(e Event) error
	Update(e Event) error
}
