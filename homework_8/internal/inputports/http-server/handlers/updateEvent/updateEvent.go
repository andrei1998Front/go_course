package updateevent

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/commands"
	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/api/responce"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Title string `json:"title,omitempty"`
	Date  string `json:"date,omitempty" validate:"omitempty,datetime=2006-01-02"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=EventUpdater
type EventUpdater interface {
	Handle(query commands.UpdateEventRequest) error
}

func New(log *slog.Logger, eventUpdater EventUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.updateEvent.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		eventID := chi.URLParam(r, "event_id")

		if eventID == "" {
			log.Error("empty event id")

			render.JSON(w, r, responce.Error("empty id"))

			return
		}

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")

			render.JSON(w, r, responce.Error("empty reuest"))

			return
		}

		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, responce.Error("failed to decode request body"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, responce.ValidationError(validateErr))

			return
		}

		err = eventUpdater.Handle(commands.UpdateEventRequest{
			ID:    eventID,
			Title: req.Title,
			Date:  req.Date,
		})

		if errors.Is(err, event.ErrDateBusy) {
			log.Error("event date is busy", slog.String("date", req.Date))

			render.JSON(w, r, responce.Error("event date is busy"))

			return
		} else if errors.Is(err, commands.ErrInvalidUUID) {
			log.Error("invalid event id", sl.Err(err))

			render.JSON(w, r, responce.Error("invalid event id"))

			return
		}

		if err != nil {
			log.Error("failed to update event", sl.Err(err))

			render.JSON(w, r, responce.Error("failed to update event"))

			return
		}

		render.JSON(w, r, responce.OK())
	}
}
