package app

import (
	"encoding/json"
	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
	"log"
	"net/http"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

func (a *App) APIGetShortURL(w http.ResponseWriter, r *http.Request) {

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link := entities.Link{
		Short:    keygen.KeyGenerate(),
		Original: req.URL,
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

	result := Response{
		Result: a.cfg.ResultURL + "/" + link.Short,
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(response)
	if err != nil {
		return
	}

}
