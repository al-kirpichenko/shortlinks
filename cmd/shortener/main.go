package main

import (
	"flag"
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/routes"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
	"log"
	"net/http"
)

func main() {

	flag.Parse()

	conf := config.NewCfg()

	newApp := app.NewApp(conf)

	data := storage.LoadInFile(conf.FilePATH)

	newApp.Storage.Load(data)

	router := routes.Router(newApp)

	log.Fatal(http.ListenAndServe(conf.Host, router))

}
