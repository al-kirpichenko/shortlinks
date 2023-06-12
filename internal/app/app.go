package app

import (
	"log"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/services/file"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

type App struct {
	cfg     *config.AppConfig
	DB      *pg.PG
	Storage storage.Storage
}

func NewApp(cfg *config.AppConfig) *App {

	return &App{
		cfg: cfg,
	}
}

func (a *App) GetConfig() *config.AppConfig {
	return a.cfg
}

func (a *App) confDB() (*pg.PG, error) {
	db := pg.NewDB(a.cfg.DataBaseString)
	if err := db.Open(); err != nil {
		return nil, err
	}
	if err := db.PingDB(); err != nil {
		return nil, err
	}
	a.DB = db
	return db, nil
}

func (a *App) ConfigureStorage() {

	switch {
	case a.cfg.DataBaseString != "":
		db, err := a.confDB()
		if err == nil {
			a.Storage = &storage.Link{
				Store: db,
			}
		} else {
			log.Fatal(err)
		}

	case a.cfg.FilePATH != "":
		store := storage.NewFileStorage(a.cfg.FilePATH)

		data, err := file.LoadFromFile(a.cfg.FilePATH)

		store.Load(data)

		if err != nil {
			log.Println("Don't load from file!")
			log.Fatal(err)
		}

		store.Load(data)

		a.Storage = store

	default:
		a.Storage = storage.NewInMemoryStorage()

	}
}
