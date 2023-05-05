package main

import (
	"flag"
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {

	router := chi.NewRouter()

	router.Get("/{id}", app.GetOriginalURL)
	router.Post("/", app.GetShortURL)

	flag.Parse()

	log.Fatal(http.ListenAndServe(config.AppConfig.Host, router))
}
