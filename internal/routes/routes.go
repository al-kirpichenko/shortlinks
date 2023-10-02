package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	middle "github.com/go-chi/chi/v5/middleware"

	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/middleware"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/cookies"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
)

func Router(app *app.App) http.Handler {

	router := chi.NewRouter()

	router.Use(logger.Logger)
	router.Use(middleware.Gzip)
	router.Use(cookies.Cookies)

	router.Get("/{id}", app.GetOriginalURL)
	router.Get("/ping", app.Ping)
	router.Post("/", app.GetShortURL)
	router.Post("/api/shorten", app.APIGetShortURL)
	router.Post("/api/shorten/batch", app.APIBatch)
	router.Get("/api/user/urls", app.APIGetUserURLs)
	router.Delete("/api/user/urls", app.APIDelUserURLs)
	router.Mount("/debug", middle.Profiler())

	return router

}
