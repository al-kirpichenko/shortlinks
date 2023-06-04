package app

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func (a *App) GetOriginalURL(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var url string
	var err error

	if a.DBReady {
		url, err = a.DataBase.Select(id)
		if err != nil {
			log.Println("Don't read data from table")
			log.Println(err)
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}
	} else {
		url, err = a.Storage.GetURL(id)

		if err != nil {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
