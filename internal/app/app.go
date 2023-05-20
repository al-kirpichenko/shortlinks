package app

import (
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/fs"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

type App struct {
	cfg      *config.AppConfig
	Storage  *storage.Storage
	Fstorage *fs.Fstorage
}

func NewApp(cfg *config.AppConfig) *App {

	return &App{
		cfg:      cfg,
		Storage:  storage.NewStorage(),
		Fstorage: fs.NewFstorage(),
	}
}
