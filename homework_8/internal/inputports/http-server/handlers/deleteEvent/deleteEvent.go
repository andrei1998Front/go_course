package deleteevent

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/commands"
	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/api/responce"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	ID string `json:"id" validate:"required,uuid"`
}

type EventDeleter interface {
	Handle(query commands.DeleteEventRequest) error
}

func New(log *slog.Logger, eventDeleter EventDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.deleteEvent.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if errors.Is(err, io.EOF) {
			log.Error("request body is empty", sl.Err(err))

			render.JSON(w, r, responce.Error("request body is empty"))

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

		err = eventDeleter.Handle(commands.DeleteEventRequest{ID: req.ID})

		if errors.Is(err, commands.ErrInvalidUUID) {
			log.Error("invalid id", sl.Err(err))

			render.JSON(w, r, responce.Error("invalid id"))

			return
		} else if errors.Is(err, event.ErrNonExistentEvent) {
			log.Error("non existent id", sl.Err(err))

			render.JSON(w, r, responce.Error("non existent id"))

			return
		}

		if err != nil {
			log.Error("failed to delete event", sl.Err(err))

			render.JSON(w, r, responce.Error("failed to delete event"))

			return
		}

		log.Info("event deleted")

		render.JSON(w, r, responce.OK())
	}
}
