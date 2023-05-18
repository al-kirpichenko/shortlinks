package routes

import (
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetRouter(app *app.App) http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Gzip)

	router.Get("/{id}", app.GetOriginalURL)
	router.Post("/", app.GetShortURL)
	router.Post("/api/shorten", app.APIGetShortURL)

	return router

}
