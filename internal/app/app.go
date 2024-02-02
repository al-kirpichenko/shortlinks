package app

import (
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/services/delurls"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

// App основная структура приложения
type App struct {
	Cfg     *config.AppConfig
	DB      *pg.PG
	Storage storage.Storage
	Worker  *delurls.Worker
}

// NewApp конструктор
func NewApp(cfg *config.AppConfig) *App {

	return &App{
		Cfg: cfg,
		DB:  pg.InitDB(cfg.DataBaseString),
	}
}

// ConfigureStorage - конфигуратор хранилища
func (a *App) ConfigureStorage() {

	switch {
	case a.Cfg.DataBaseString != "":
		a.Storage = storage.NewLinkStorage(a.DB)
	case a.Cfg.FilePATH != "":
		a.Storage = storage.NewFileStorage(a.Cfg.FilePATH)
	default:
		a.Storage = storage.NewInMemoryStorage()
	}
}
