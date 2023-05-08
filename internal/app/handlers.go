package app

import (
	"fmt"
	"github.com/al-kirpichenko/shortlinks/internal/shortlinkgen"
	"github.com/go-chi/chi/v5"
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

	id := shortlinkgen.GenerateID()

	a.storage.SetURL(id, url)

	response := fmt.Sprintf(a.cfg.ResultURL+"/%s", id)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = io.WriteString(w, response)
	if err != nil {
		return
	}
}

func (a *App) GetOriginalURL(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	url, err := a.storage.GetURL(id)

	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
