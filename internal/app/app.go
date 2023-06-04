package app

import (
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/database"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

type App struct {
	cfg      *config.AppConfig
	Storage  *storage.InMemoryStorage
	DataBase *database.DBStore
	DBReady  bool
}

func NewApp(cfg *config.AppConfig) *App {

	return &App{
		cfg:     cfg,
		Storage: storage.NewInMemoryStorage(),
	}
}

func (a *App) ConfigureDB() error {
	db := database.NewDB(a.cfg.DataBaseString)
	if err := db.Open(); err != nil {
		return err
	}
	a.DataBase = db
	a.DBReady = true
	return nil
}
