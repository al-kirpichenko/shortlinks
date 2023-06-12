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

	run(newApp)

}

func run(app *app.App) {

	router := routes.Router(app)

	log.Fatal(http.ListenAndServe(app.GetConfig().Host, router))
}
