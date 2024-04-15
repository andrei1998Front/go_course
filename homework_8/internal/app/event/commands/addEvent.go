package commands

import (
	"log/slog"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/google/uuid"
)

type AddEventRequest struct {
	Title string
	Date  time.Time
}

func NewAddEventRequest(title string, date time.Time) *AddEventRequest {
	return &AddEventRequest{Title: title, Date: date}
}

type AddEventRequestHandler interface {
	Handle(query *AddEventRequest) error
}

type addEventRequestHandler struct {
	log  *slog.Logger
	repo event.Repository
}

func NewAddEventRequestHandler(log *slog.Logger, repo event.Repository) AddEventRequestHandler {
	return addEventRequestHandler{log: log, repo: repo}
}

func (h addEventRequestHandler) Handle(query *AddEventRequest) error {
	op := "app.queries.addEvent"

	newEvent := event.Event{
		ID:    uuid.New(),
		Title: query.Title,
		Date:  query.Date,
	}

	log := h.log.With(slog.String("op", op))

	if err := h.repo.Add(newEvent); err != nil {
		log.Error(err.Error())

		return err
	}

	return nil
}
