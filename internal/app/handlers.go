package app

import (
	"fmt"
	"github.com/al-kirpichenko/shortlinks/internal/random"
	"io"
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

	id := random.GenerateID()

	a.storage.Urls[id] = url

	response := fmt.Sprintf(a.cfg.ResultURL+"/%s", id)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = io.WriteString(w, response)
	if err != nil {
		return
	}
}

func (a *App) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]

	url, ok := a.storage.Urls[id]

	if !ok {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
