package event

import (
	"time"

	"github.com/google/uuid"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=Repository
type Repository interface {
	GetByID(id uuid.UUID) (*Event, error)
	GetByDate(dt time.Time) (*Event, error)
	GetAll() ([]Event, error)
	Add(e Event) error
	Delete(id uuid.UUID) error
	Update(e Event) error
}
