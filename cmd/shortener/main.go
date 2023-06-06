package main

import (
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/al-kirpichenko/shortlinks/internal/routes"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
	"log"
	"net/http"
)

func main() {

	conf := config.NewAppConfig()

	logger.InitLogger()

	newApp := app.NewApp(conf)

	run(newApp)

}

func run(app *app.App) {

	if err := app.ConfigureDB(); err != nil {
		log.Println("Don't configure Database!")
		log.Println(err)
	}

	if app.GetConfig().FilePATH != "" {

		data, err := storage.LoadFromFile(app.GetConfig().FilePATH)

		if err != nil {
			log.Println("Don't load from file!")
			log.Println(err)
		}

		app.Storage.Load(data)
	}
	router := routes.Router(app)

	log.Fatal(http.ListenAndServe(app.GetConfig().Host, router))
}
