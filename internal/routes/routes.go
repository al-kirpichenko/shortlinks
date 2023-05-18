package routes

import (
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/middleware"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/encoding"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Route(cfg *config.AppConfig) http.Handler {

	newApp := app.NewApp(cfg)
	router := chi.NewRouter()

	//router.Get("/{id}", log.WithLogging(newApp.GetOriginalURL))
	//router.Post("/", log.WithLogging(newApp.GetShortURL))
	//router.Post("/api/shorten", log.WithLogging(newApp.APIGetShortURL))

	router.Get("/{id}", middleware.Conveyor(newApp.GetOriginalURL, logger.WithLogging, encoding.GzipMiddleware))
	router.Post("/", middleware.Conveyor(newApp.GetShortURL, logger.WithLogging, encoding.GzipMiddleware))
	router.Post("/api/shorten", middleware.Conveyor(newApp.APIGetShortURL, logger.WithLogging, encoding.GzipMiddleware))

	return router

}
