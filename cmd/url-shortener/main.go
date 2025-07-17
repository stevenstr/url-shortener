package main

import (
	"log/slog"
	"os"

	"github.com/stevenstr/url-shortener/internal/config"
	"github.com/stevenstr/url-shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// Done: init config: cleanenv
	cfg := config.MustLoad()

	// Done: init logger: slog (log/slog)
	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener...", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled...")

	// TODO: init  storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	_ = storage

	// TODO: init router: chi, "chi render"

	// TODO: run server

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		// текстовый
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		// текстовый
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		// а вот тут уже json для grafana kibana
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
