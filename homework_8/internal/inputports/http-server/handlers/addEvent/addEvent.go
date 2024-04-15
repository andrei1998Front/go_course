package addevent

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/commands"
	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/api/responce"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Title string `json:"title" validate:"required"`
	Date  string `json:"date" validate:"required,datetime=2006-01-02"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=EventAdder
type EventAdder interface {
	Handle(query *commands.AddEventRequest) error
}

func New(log *slog.Logger, eventAdder EventAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.addEvent.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")

			render.JSON(w, r, responce.Error("empty request"))

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

		dateEvent, err := time.Parse("2006-01-02", req.Date)

		if err != nil {
			log.Error("failed to convert dt", sl.Err(err))

			render.JSON(w, r, responce.Error("internal error"))

			return
		}

		addEventRequest := commands.NewAddEventRequest(req.Title, dateEvent)

		err = eventAdder.Handle(addEventRequest)

		if errors.Is(err, event.ErrExistentID) {
			log.Info("event id already exists")
			render.JSON(w, r, responce.Error("event already exists"))

			return
		} else if errors.Is(err, event.ErrDateBusy) {
			log.Error("event date is busy", slog.String("date", req.Date))

			render.JSON(w, r, responce.Error("event date is busy"))

			return
		}

		if err != nil {
			log.Error("failed to add event", sl.Err(err))

			render.JSON(w, r, responce.Error("failed to add event"))

			return
		}

		log.Info("event added")

		render.JSON(w, r, responce.OK())
	}
}
