package app

import (
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

type App struct {
	cfg      *config.AppConfig
	Storage  *storage.InMemoryStorage
	Fstorage *storage.FileStorage
}

func NewApp(cfg *config.AppConfig) *App {

	return &App{
		cfg:      cfg,
		Storage:  storage.NewInMemoryStorage(),
		Fstorage: storage.NewFileStorage(),
	}
}
