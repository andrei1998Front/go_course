package queries

import (
	"log/slog"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/google/uuid"
)

type GetAllEventResponce struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Title string    `json:"title,omitempty"`
	Date  time.Time `json:"date,omitempty"`
}

type GetAllEventRequestHandler interface {
	Handle() ([]*GetAllEventResponce, error)
}

type getAllEventRequestHandler struct {
	log  *slog.Logger
	repo event.Repository
}

func NewGetAllEventRequestHandler(log *slog.Logger, repo event.Repository) GetAllEventRequestHandler {
	return getAllEventRequestHandler{log: log, repo: repo}
}

func (h getAllEventRequestHandler) Handle() ([]*GetAllEventResponce, error) {
	const op = "app.queries.getAllEvent"

	log := h.log.With(slog.String("op", op))

	events, err := h.repo.GetAll()

	if err != nil {
		log.Error(err.Error())

		return []*GetAllEventResponce{}, err
	}

	var allEvents []*GetAllEventResponce

	for _, event := range events {
		allEvents = append(allEvents, &GetAllEventResponce{
			ID:    event.ID,
			Title: event.Title,
			Date:  event.Date,
		})
	}

	return allEvents, nil
}
