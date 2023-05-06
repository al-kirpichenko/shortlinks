package routes

import (
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Route() http.Handler {

	router := chi.NewRouter()
	router.Get("/{id}", app.GetOriginalURL)
	router.Post("/", app.GetShortURL)
	return router

}
