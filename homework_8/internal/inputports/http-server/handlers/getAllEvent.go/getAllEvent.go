package getalleventgo

import (
	"log/slog"
	"net/http"

	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/queries"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/api/responce"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Responce struct {
	responce.Responce
	Events []*queries.GetAllEventResponce `json:"events,omitempty"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=EventsGetter
type EventsGetter interface {
	Handle() ([]*queries.GetAllEventResponce, error)
}

func New(log *slog.Logger, eventsGetter EventsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.getAllEvents.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		allEvents, err := eventsGetter.Handle()

		if err != nil {
			log.Error("failed to get events", sl.Err(err))

			render.JSON(w, r, responce.Error("failed to get events"))
		}

		log.Info("all events geted")

		render.JSON(w, r, Responce{
			Responce: responce.OK(),
			Events:   allEvents,
		})

		return
	}
}
