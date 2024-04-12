package app

import (
	"log/slog"

	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/commands"
	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/queries"
	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
)

type Queries struct {
	GetAllEventsHandler queries.GetAllEventRequestHandler
	GetEventHandler     queries.GetEventRequestHandler
}

type Commands struct {
	AddEventHandler    commands.AddEventRequestHandler
	UpdateEventHandler commands.UpadateEventRequestHandler
	DeleteEventHandler commands.DeleteEventRequestHandler
}

type EventService struct {
	Queries
	Commands
}

type Service struct {
	EventService
}

func NewServices(eventRepo event.Repository, log *slog.Logger) Service {
	return Service{
		EventService{
			Queries: Queries{
				GetAllEventsHandler: queries.NewGetAllEventRequestHandler(log, eventRepo),
				GetEventHandler:     queries.NewGetRequestHandler(log, eventRepo),
			},
			Commands: Commands{
				AddEventHandler:    commands.NewAddEventRequestHandler(log, eventRepo),
				UpdateEventHandler: commands.NewUpdateEventRequestHandler(log, eventRepo),
				DeleteEventHandler: commands.NewEventRequestHandler(log, eventRepo),
			},
		},
	}
}
