package app

import (
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

type contextKey string

const Token contextKey = "token"

type App struct {
	cfg     *config.AppConfig
	DB      *pg.PG
	Storage storage.Storage
}

func NewApp(cfg *config.AppConfig) *App {

	return &App{
		cfg: cfg,
		DB:  pg.InitDB(cfg.DataBaseString),
	}
}

func (a *App) ConfigureStorage() {

	switch {
	case a.cfg.DataBaseString != "":
		a.Storage = storage.NewLinkStorage(a.DB)
	case a.cfg.FilePATH != "":
		a.Storage = storage.NewFileStorage(a.cfg.FilePATH)
	default:
		a.Storage = storage.NewInMemoryStorage()
	}
}
