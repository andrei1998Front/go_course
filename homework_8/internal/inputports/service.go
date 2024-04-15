package inputports

import (
	"log/slog"

	"github.com/andrei1998Front/go_course/homework_8/internal/app"
	"github.com/andrei1998Front/go_course/homework_8/internal/config"
	server "github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server"
)

type Service struct {
	Server *server.Server
}

func New(log *slog.Logger, appService app.Service, cfg config.HTTPServer) *Service {
	return &Service{Server: server.New(appService, log, cfg)}
}
