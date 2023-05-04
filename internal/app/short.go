package app

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

var urls map[string]string

func generateID() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetShortURL(w http.ResponseWriter, r *http.Request) {

	urls = make(map[string]string)
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
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := generateID()
	urls[id] = url
	response := fmt.Sprintf("http://localhost:8080/%s", id)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = io.WriteString(w, response)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	return
}

func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	url, ok := urls[id]
	if !ok {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
	return
}
