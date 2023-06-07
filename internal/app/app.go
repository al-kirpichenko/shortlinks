package app

import (
	"errors"
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
	"log"
)

type App struct {
	cfg      *config.AppConfig
	Storage  *storage.InMemoryStorage
	DataBase *pg.PG
	DBReady  bool
}

func NewApp(cfg *config.AppConfig) *App {

	return &App{
		cfg:     cfg,
		Storage: storage.NewInMemoryStorage(),
		DBReady: false,
	}
}

func (a *App) GetConfig() *config.AppConfig {
	return a.cfg
}

func (a *App) ConfigureDB() error {

	if a.cfg.DataBaseString != "" {
		db := pg.NewDB(a.cfg.DataBaseString)
		if err := db.Open(); err != nil {
			log.Println(err)
			return err
		}
		a.DataBase = db
		a.DBReady = true
		return nil
	}
	return errors.New("dataBaseString is empty")
}
