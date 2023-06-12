package main

import (
	"log"
	"net/http"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/al-kirpichenko/shortlinks/internal/routes"
)

func main() {

	conf := config.NewAppConfig()

	logger.InitLogger()

	newApp := app.NewApp(conf)

	newApp.DB = pg.InitDB(conf.DataBaseString)
	defer newApp.DB.Close()

	newApp.ConfigureStorage()

	router := routes.Router(newApp)

	log.Fatal(http.ListenAndServe(conf.Host, router))
}
