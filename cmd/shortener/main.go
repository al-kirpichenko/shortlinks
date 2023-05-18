package main

import (
	"flag"
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/routes"
	"log"
	"net/http"
)

func main() {

	conf := config.NewCfg()
	newApp := app.NewApp(conf)

	router := routes.GetRouter(newApp)
	flag.Parse()
	log.Fatal(http.ListenAndServe(conf.Host, router))

}
