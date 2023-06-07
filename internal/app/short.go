package app

import (
	"fmt"
	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
	"io"
	"log"
	"net/http"
)

func (a *App) GetShortURL(w http.ResponseWriter, r *http.Request) {

	responseData, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, fmt.Sprintf("cannot read request body: %s", err), http.StatusBadRequest)
		return
	}
	if string(responseData) == "" {
		http.Error(w, "Empty POST request body!", http.StatusBadRequest)
		return
	}
	url := string(responseData)

	id := keygen.KeyGenerate()

	link := entities.Link{
		Short:    id,
		Original: url,
	}

	linkModel := models.Link{
		Store: a.DataBase,
	}

	if a.DBReady {
		_, err = linkModel.Insert(&link)
		if err != nil {
			log.Println("Don't insert url!")
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if a.cfg.FilePATH != "" {
		fileStorage := storage.NewFileStorage()

		fileStorage.Short = link.Short
		fileStorage.Original = link.Original

		err = storage.SaveToFile(fileStorage, a.cfg.FilePATH)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	a.Storage.SetURL(link.Short, link.Original)

	response := fmt.Sprintf(a.cfg.ResultURL+"/%s", link.Short)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = io.WriteString(w, response)
	if err != nil {
		return
	}
}
