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

	return &App{
		cfg:     cfg,
		Storage: ConfigureStorage(cfg),
	}
}

func (a *App) GetConfig() *config.AppConfig {
	return a.cfg
}

func confDB(conn string) (*pg.PG, error) {
	db := pg.NewDB(conn)
	if err := db.Open(); err != nil {
		return nil, err
	}
	if err := db.PingDB(); err != nil {
		return nil, err
	}
	return db, nil
}

func ConfigureStorage(conf *config.AppConfig) storage.Storage {
	if conf.DataBaseString != "" {
		db, err := confDB(conf.DataBaseString)
		if err == nil {
			return &models.Link{
				Store: db,
			}
		} else {
			log.Fatal(err)
			return nil
		}

	} else if conf.FilePATH != "" {
		store := storage.NewFileStorage(conf.FilePATH)

		data, err := file.LoadFromFile(conf.FilePATH)

		store.Load(data)

		if err != nil {
			log.Println("Don't load from file!")
			log.Fatal(err)
			return nil
		}

		store.Load(data)

		return store
	}
	return storage.NewInMemoryStorage()
}
