package main

import (
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {

	router := chi.NewRouter()

	router.Get("/{id}", app.GetOriginalURL)
	router.Post("/", app.GetShortURL)

	log.Fatal(http.ListenAndServe(":8080", router))
}
