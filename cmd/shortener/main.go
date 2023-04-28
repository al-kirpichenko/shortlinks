package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const baseUrl = "http://localhost:8080"

var (
	linkListShorts map[string]string
)

func mainPage(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		err := request.ParseForm()
		if err != nil {
			return
		}
		longUrl := request.FormValue("")
		if longUrl == "" {
			writer.WriteHeader(http.StatusCreated)
		}
		if isValidUrl(longUrl) {

			shortUrl := baseUrl + "/" + shorting()
			linkListShorts[shortUrl] = longUrl

			writer.WriteHeader(http.StatusCreated)
			fmt.Fprintf(writer, shortUrl)

			return
		}
	} else if request.Method == http.MethodGet {

		shortUrl := baseUrl + request.URL.String()
		if shortUrl == "" {
			writer.WriteHeader(http.StatusTemporaryRedirect)
		}

		writer.Header().Set("Location", linkListShorts[shortUrl])
		writer.WriteHeader(http.StatusTemporaryRedirect)

	} else {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

}

func shorting() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func isValidUrl(token string) bool {
	_, err := url.ParseRequestURI(token)
	if err != nil {
		return false
	}
	u, err := url.Parse(token)
	if err != nil || u.Host == "" {
		return false
	}
	return true
}

func main() {

	linkListShorts = map[string]string{}

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
