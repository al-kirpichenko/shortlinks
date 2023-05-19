package app

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *App) GetOriginalURL(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	url, err := a.Storage.GetURL(id)

	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
