package server

import (
	"log/slog"
	"net/http"

	"github.com/andrei1998Front/go_course/homework_8/internal/app"
	"github.com/andrei1998Front/go_course/homework_8/internal/config"
	addevent "github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/addEvent"
	deleteevent "github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/deleteEvent"
	getalleventgo "github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/getAllEvent"
	getevent "github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/getEvent"
	updateevent "github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/updateEvent"
	mwlogger "github.com/andrei1998Front/go_course/homework_8/internal/lib/logger/mwLogger"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	AppService app.Service
	Router     *chi.Mux
	Cfg        config.HTTPServer
}

func New(appService app.Service, log *slog.Logger, cfg config.HTTPServer) *Server {
	router := setupRouter(chi.NewRouter(), log, &appService)

	return &Server{
		AppService: appService,
		Router:     router,
		Cfg:        cfg,
	}
}

func (s *Server) NewHTTPServer() *http.Server {
	return &http.Server{
		Addr:         s.Cfg.Address,
		Handler:      s.Router,
		ReadTimeout:  s.Cfg.Timeout,
		WriteTimeout: s.Cfg.Timeout,
		IdleTimeout:  s.Cfg.IdleTimeout,
	}
}

func setupRouter(r *chi.Mux, log *slog.Logger, appService *app.Service) *chi.Mux {
	r.Use(middleware.RequestID)
	r.Use(mwlogger.New(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Post("/events/", addevent.New(log, appService.AddEventHandler))
	r.Get("/events/", getalleventgo.New(log, appService.GetAllEventsHandler))
	r.Get("/events/{event_id}", getevent.New(log, appService.GetEventHandler))
	r.Patch("/events/{event_id}", updateevent.New(log, appService.UpdateEventHandler))
	r.Delete("/events/", deleteevent.New(log, appService.DeleteEventHandler))
	return r
}
