package bootstrap

import (
	"log/slog"

	"github.com/FelipeWoo/newgo/internal/config"
)

type App struct {
	Config config.AppConfig
	Logger *slog.Logger
}

func Boot(logName string) (App, error) {
	cfg, err := config.Load()
	if err != nil {
		return App{}, err
	}

	logger, err := applog.New(logName, cfg.LogLevel, cfg.Root)
	if err != nil {
		return App{}, err
	}

	logger.Info("system initialized", "env", cfg.Env, "root", cfg.Root)

	return App{
		Config: cfg,
		Logger: logger,
	}, nil
}
