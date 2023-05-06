package main

import (
	"flag"
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/al-kirpichenko/shortlinks/internal/routes"
	"log"
	"net/http"
)

func main() {

	router := routes.Route()

	flag.Parse()

	log.Fatal(http.ListenAndServe(config.AppConfig.Host, router))
}
