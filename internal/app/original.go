package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GetOriginalURL - Эндпоинт с методом GET и путём /{id}, где id — идентификатор сокращённого URL
// (например, /EwHXdJfB). В случае успешной обработки запроса сервер возвращает ответ с кодом 307
// и оригинальным URL в HTTP-заголовке Location.
func (a *App) GetOriginalURL(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	link, err := a.Storage.GetOriginal(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	if link.Deleted {
		w.WriteHeader(http.StatusGone)
		return
	}

	w.Header().Set("Location", link.Original)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
