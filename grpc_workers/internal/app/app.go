package app

import (
	"log/slog"

	grpcapp "github.com/andrei1998Front/grpc_workers_server/internal/app/grpc"
	wrkrsService "github.com/andrei1998Front/grpc_workers_server/internal/services"
	"github.com/andrei1998Front/grpc_workers_server/internal/storage"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	connStr string,
) *App {
	storage := storage.New(connStr)

	wrkrsService := wrkrsService.New(log, storage)

	grpcApp := grpcapp.New(log, *wrkrsService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
