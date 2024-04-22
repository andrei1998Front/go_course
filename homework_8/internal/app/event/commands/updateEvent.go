package commands

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/google/uuid"
)

type UpdateEventRequest struct {
	ID    string
	Title string
	Date  string
}

type UpadateEventRequestHandler interface {
	Handle(query UpdateEventRequest) error
	setupUpdatableEvent(req *UpdateEventRequest) (event.Event, error)
}

type updateEventRequestHandler struct {
	log  *slog.Logger
	repo event.Repository
}

func NewUpdateEventRequestHandler(log *slog.Logger, repo event.Repository) UpadateEventRequestHandler {
	return updateEventRequestHandler{
		log:  log,
		repo: repo,
	}
}

func (h updateEventRequestHandler) Handle(query UpdateEventRequest) error {
	const op = "app.comands.updateEvent"

	log := h.log.With(slog.String("op", op))

	if query.Title == "" && query.Date == "" {
		return ErrEmptyQuery
	}

	updatetableEvent, err := h.setupUpdatableEvent(&query)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	if err := h.repo.Update(updatetableEvent); err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (h updateEventRequestHandler) setupUpdatableEvent(req *UpdateEventRequest) (event.Event, error) {
	var ev event.Event

	eventID, err := uuid.Parse(req.ID)

	if err != nil {
		return event.Event{}, ErrInvalidUUID
	}

	ev.ID = eventID

	var evByID *event.Event
	if req.Title == "" || req.Date == "" {
		evByID, err = h.repo.GetByID(eventID)

		if err != nil {
			return ev, err
		}
	}

	if err := setEmptyEventFields(&ev, evByID, req); err != nil {
		return ev, err
	}

	return ev, nil
}

func setEmptyEventFields(ev *event.Event, evByID *event.Event, req *UpdateEventRequest) error {
	if req.Title == "" && evByID.Title == "" {
		return ErrInvalidTitle
	}

	if req.Title != "" {
		ev.Title = req.Title
	} else {
		ev.Title = evByID.Title
	}

	if req.Date == "" && evByID.Date.IsZero() {
		return ErrInvalidDate
	}

	if req.Date != "" {
		evDate, err := time.Parse("2006-01-02", req.Date)

		if err != nil {
			return fmt.Errorf("%w: %w", ErrInvalidDate, err)
		}

		ev.Date = evDate
	} else {
		ev.Date = evByID.Date
	}

	return nil
}
