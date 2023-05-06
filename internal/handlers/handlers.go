package handlers

import (
	"fmt"
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/al-kirpichenko/shortlinks/internal/random"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
	"io"
	"net/http"
)

func GetShortURL(w http.ResponseWriter, r *http.Request) {

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

	id := random.GenerateID()

	storage.SetURL(url, id)

	response := fmt.Sprintf(config.AppConfig.ResultURL+"/%s", id)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = io.WriteString(w, response)
	if err != nil {
		return
	}
}

func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]

	//id := chi.URLParam(r, "id")

	//url, ok := urls[id]

	url, err := storage.GetURL(id)

	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
