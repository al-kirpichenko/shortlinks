package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var urls map[string]string

func main() {
	urls = make(map[string]string)
	http.HandleFunc("/", shortenURL)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			return
		}

		url := r.FormValue("")

		id := generateID()
		urls[id] = url
		response := fmt.Sprintf("http://localhost:8080/%s", id)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(response))
		if err != nil {
			return
		}

	} else if r.Method == http.MethodGet {
		id := r.URL.Path[1:]
		url, ok := urls[id]
		if !ok {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		//fmt.Println(url)
		w.Header().Add("Location", url)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func generateID() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
