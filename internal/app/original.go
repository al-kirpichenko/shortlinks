package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

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
