package routes

import (
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/log"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Route(cfg *config.AppConfig) http.Handler {

	newApp := app.NewApp(cfg)
	router := chi.NewRouter()
	router.Get("/{id}", log.WithLogging(newApp.GetOriginalURL))
	router.Post("/", log.WithLogging(newApp.GetShortURL))
	return router
}
