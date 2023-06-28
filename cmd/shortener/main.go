package main

import (
	"log"
	"net/http"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/al-kirpichenko/shortlinks/internal/routes"
)

func main() {

	conf := config.NewAppConfig()

	logger.InitLogger()

	newApp := app.NewApp(conf)

	newApp.ConfigureStorage()

	router := routes.Router(newApp)

	defer newApp.DB.Close()

	log.Fatal(http.ListenAndServe(conf.Host, router))
}
