package commands

import (
	"log/slog"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/google/uuid"
)

type UpdateEventRequest struct {
	ID    string
	Title string
	Date  time.Time
}

type UpadateEventRequestHandler interface {
	Handle(query UpdateEventRequest) error
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

	eventID, err := uuid.Parse(query.ID)

	if err != nil {
		log.Error(ErrInvalidUUID.Error() + " - " + err.Error())
		return ErrInvalidUUID
	}

	updatetableEvent := event.Event{
		ID:    eventID,
		Title: query.Title,
		Date:  query.Date,
	}

	if err := h.repo.Update(updatetableEvent); err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
