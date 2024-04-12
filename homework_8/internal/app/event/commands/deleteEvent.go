package commands

import (
	"log/slog"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/google/uuid"
)

type DeleteEventRequest struct {
	ID string
}

type DeleteEventRequestHandler interface {
	Handle(query DeleteEventRequest) error
}

type deleteEventRequestHandler struct {
	log  *slog.Logger
	repo event.Repository
}

func NewEventRequestHandler(log *slog.Logger, repo event.Repository) DeleteEventRequestHandler {
	return deleteEventRequestHandler{
		log:  log,
		repo: repo,
	}
}

func (h deleteEventRequestHandler) Handle(query DeleteEventRequest) error {
	const op = "app.commands.delete"

	log := h.log.With(slog.String("op", op))

	eventID, err := uuid.Parse(query.ID)

	if err != nil {
		log.Error(ErrInvalidUUID.Error() + " - " + err.Error())
		return ErrInvalidUUID
	}

	if err := h.repo.Delete(eventID); err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
