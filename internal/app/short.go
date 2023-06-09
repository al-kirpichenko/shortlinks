package app

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
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

	link := entities.Link{
		Short:    keygen.KeyGenerate(),
		Original: string(responseData),
	}

	if err := a.Storage.Insert(&link); err != nil {
		log.Println("Don't insert url!")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//if a.DBReady {
	//
	//	linkModel := models.Link{
	//		Store: a.DataBase,
	//	}
	//
	//	err = linkModel.Insert(&link)
	//	if err != nil {
	//		log.Println("Don't insert url!")
	//		log.Println(err)
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//} else {
	//	if a.cfg.FilePATH != "" {
	//
	//		err = storage.SaveToFile(&link, a.cfg.FilePATH)
	//
	//		if err != nil {
	//			log.Println(err)
	//			http.Error(w, err.Error(), http.StatusInternalServerError)
	//			return
	//		}
	//		//a.Storage.SetURL(link.Short, link.Original)
	//		err = a.Storage.Insert(&link)
	//		if err != nil {
	//			log.Println(err)
	//			http.Error(w, err.Error(), http.StatusInternalServerError)
	//			return
	//		}
	//	}
	//	//a.Storage.SetURL(link.Short, link.Original)
	//	err = a.Storage.Insert(&link)
	//	if err != nil {
	//		log.Println(err)
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//}

	response := fmt.Sprintf(a.cfg.ResultURL+"/%s", link.Short)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = io.WriteString(w, response)
	if err != nil {
		return
	}
}
