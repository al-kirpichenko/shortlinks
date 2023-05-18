package main

import (
	"flag"
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/al-kirpichenko/shortlinks/internal/routes"
	"log"
	"net/http"
)

func main() {

	conf := config.NewCfg()
	router := routes.Route(conf)
	flag.Parse()
	log.Fatal(http.ListenAndServe(conf.Host, router))

}
