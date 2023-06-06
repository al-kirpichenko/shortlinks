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

	if err := newApp.ConfigureDB(); err != nil {
		log.Println("Don't configure Database!")
		log.Println(err)
	}

	data, err := storage.LoadFromFile(conf.FilePATH)

	if err != nil {
		log.Println("filePath is empty!")
		log.Println(err)
	}

	newApp.Storage.Load(data)

	router := routes.Router(newApp)

	log.Fatal(http.ListenAndServe(conf.Host, router))

}
