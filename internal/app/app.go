package app

import (
	"log"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/file"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

type App struct {
	cfg     *config.AppConfig
	DB      *pg.PG
	Storage storage.Storage
}

func NewApp(cfg *config.AppConfig) *App {

	app := &App{
		cfg: cfg,
	}

	return app
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
	return db, nil
}

func (a *App) ConfigureStorage() {
	if a.cfg.DataBaseString != "" {
		log.Println("start configure db")
		db, err := a.confDB()
		if err == nil {
			a.Storage = &models.Link{
				Store: db,
			}
			a.DB = db
		} else {
			log.Fatal(err)
		}

	} else if a.cfg.FilePATH != "" {
		store := storage.NewFileStorage(a.cfg.FilePATH)

		data, err := file.LoadFromFile(a.cfg.FilePATH)

		store.Load(data)

		if err != nil {
			log.Println("Don't load from file!")
			log.Fatal(err)
		}

		store.Load(data)

		a.Storage = store
	} else {
		a.Storage = storage.NewInMemoryStorage()
	}
}
