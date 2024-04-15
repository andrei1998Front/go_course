package queries

import (
	"log/slog"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/google/uuid"
)

type GetEventRequest struct {
	EventID uuid.UUID
}

type GetEventResponce struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

type GetEventRequestHandler interface {
	Handle(query GetEventRequest) (*GetEventResponce, error)
}

type getEventRequestHandler struct {
	log  *slog.Logger
	repo event.Repository
}

func NewGetRequestHandler(log *slog.Logger, repo event.Repository) GetEventRequestHandler {
	return getEventRequestHandler{log: log, repo: repo}
}

func (h getEventRequestHandler) Handle(query GetEventRequest) (*GetEventResponce, error) {
	const op = "app.queries.getEvent"

	log := h.log.With(slog.String("op", op))

	event, err := h.repo.GetByID(query.EventID)

	if err != nil {
		log.Error(err.Error())
		return &GetEventResponce{}, nil
	}

	return &GetEventResponce{
		ID:    event.ID,
		Title: event.Title,
		Date:  event.Date,
	}, nil
}
