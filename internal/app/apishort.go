package app

import (
	"encoding/json"
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

	id := keygen.KeyGenerate()

	if a.DBReady {

		if _, err := a.DataBase.DB.Query("SELECT * FROM links"); err != nil {
			log.Println("table not found!")
			log.Println(err)
			if err := a.DataBase.CreateTable(); err != nil {
				log.Println("table don't created!")
				log.Println(err)
			}
			log.Println("the table has been created!")
		}

		if err := a.DataBase.Insert(id, req.URL); err != nil {
			log.Println("Don't insert url!")
			log.Println(err)
		}

	}

	a.Storage.SetURL(id, req.URL)

	fileStorage := storage.NewFileStorage()

	fileStorage.Short = id
	fileStorage.Original = req.URL

	err = storage.SaveToFile(fileStorage, a.cfg.FilePATH)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := Response{
		Result: a.cfg.ResultURL + "/" + id,
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
