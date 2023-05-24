package main

import (
	"flag"
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/al-kirpichenko/shortlinks/internal/routes"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
	"log"
	"net/http"
)

func init() {
	flag.Parse()
}

func main() {

	conf := config.NewAppConfig()

	logger.InitLogger()

	newApp := app.NewApp(conf)

	data := storage.LoadFromFile(conf.FilePATH)

	newApp.Storage.Load(data)

	router := routes.Router(newApp)

	log.Println(conf.FilePATH)

	log.Fatal(http.ListenAndServe(conf.Host, router))

}
