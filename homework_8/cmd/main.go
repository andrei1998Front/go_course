package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/andrei1998Front/go_course/homework_8/internal/app"
	"github.com/andrei1998Front/go_course/homework_8/internal/config"
	"github.com/andrei1998Front/go_course/homework_8/internal/interfaceadapters"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	if _, err := os.Stat(cfg.LogFile); os.IsNotExist(err) {
		log.Fatal("Log file is not exists")
	}

	file, err := os.OpenFile(cfg.LogFile, os.O_APPEND, os.ModeAppend)

	if err != nil {
		log.Fatal("error opening log file")
	}

	log := setupLogger(cfg.Env, file)
	log = log.With(slog.String("env", cfg.Env))

	interfaceadaptersServices := interfaceadapters.NewService()
	app := app.NewServices(interfaceadaptersServices.Repo, log)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
}

func setupLogger(env string, logFile *os.File) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
