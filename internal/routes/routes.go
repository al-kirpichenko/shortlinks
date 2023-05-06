package routes

import (
	"github.com/al-kirpichenko/shortlinks/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Route() http.Handler {

	router := chi.NewRouter()
	router.Get("/{id}", handlers.GetOriginalURL)
	router.Post("/", handlers.GetShortURL)
	return router

}
