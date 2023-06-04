package app

import (
	"fmt"
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

	if a.DBReady {

		rows, err := a.DataBase.DB.Query("SELECT * FROM links")
		if err != nil {
			log.Println("table not found!")
			if err := a.DataBase.CreateTable(); err != nil {
				log.Println("table don't created!")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Println("the table has been created!")
		}
		defer rows.Close()

		if err := a.DataBase.Insert(id, url); err != nil {
			log.Println("Don't insert url!")
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	a.Storage.SetURL(id, url)

	fileStorage := storage.NewFileStorage()

	fileStorage.Short = id
	fileStorage.Original = url

	err = storage.SaveToFile(fileStorage, a.cfg.FilePATH)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf(a.cfg.ResultURL+"/%s", id)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = io.WriteString(w, response)
	if err != nil {
		return
	}
}
