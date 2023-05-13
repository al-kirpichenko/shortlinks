package app

import (
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

type App struct {
	cfg     *config.AppConfig
	storage *storage.Storage
}

func NewApp(cfg *config.AppConfig) *App {

	return &App{
		cfg:     cfg,
		storage: storage.NewStorage(),
	}
}
