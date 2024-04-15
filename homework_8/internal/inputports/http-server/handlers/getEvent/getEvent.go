package getevent

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/queries"
	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/api/responce"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type Responce struct {
	responce.Responce
	Event *queries.GetEventResponce `json:"event"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=EventGetter
type EventGetter interface {
	Handle(query queries.GetEventRequest) (*queries.GetEventResponce, error)
}

func New(log *slog.Logger, eventGetter EventGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.getEvent.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		eventID := chi.URLParam(r, "event_id")

		eventUUID, err := uuid.Parse(eventID)

		if err != nil {
			log.Error("invalid event id", sl.Err(err))

			render.JSON(w, r, responce.Error("invalid event id"))

			return
		}

		ev, err := eventGetter.Handle(queries.GetEventRequest{EventID: eventUUID})

		if errors.Is(err, event.ErrNonExistentEvent) {
			log.Error("non existent event")

			render.JSON(w, r, responce.Error("non existent event"))

			return
		}

		render.JSON(w, r, Responce{
			Responce: responce.OK(),
			Event:    ev,
		})

		return
	}
}
