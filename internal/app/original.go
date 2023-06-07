package app

import (
	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func (a *App) GetOriginalURL(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	link := &entities.Link{
		Short: id,
	}
	linkModel := &models.Link{
		Store: a.DataBase,
	}
	var err error
	if a.DBReady {
		link, err = linkModel.GetOriginal(link)
		if err != nil {
			log.Println("Don't read data from table")
			log.Println(err)
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}
	}
	if link.Original == "" {
		link.Original, err = a.Storage.GetURL(id)
		if err != nil {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}
	}

	w.Header().Set("Location", link.Original)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
