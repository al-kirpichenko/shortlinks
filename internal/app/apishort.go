package app

import (
	"encoding/json"
	"github.com/al-kirpichenko/shortlinks/internal/shortlinkgen"
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

	id := shortlinkgen.GenerateID()

	a.Storage.SetURL(id, req.URL)

	a.Fstorage.AddURL(id, req.URL)

	err = a.Fstorage.SaveToFile(a.cfg.FilePATH)
	if err != nil {
		log.Fatal(err)
	}

	result := Response{
		Result: a.cfg.ResultURL + "/" + id,
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(response)
	if err != nil {
		return
	}

}
