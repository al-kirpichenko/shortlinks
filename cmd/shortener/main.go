package main

import (
	"io"
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
		if isValidUrl(longUrl) {

			writer.Header().Set("content-type", "text/plain")

			shortUrl := baseUrl + "/" + shorting()

			linkListShorts[shortUrl] = longUrl

			io.WriteString(writer, shortUrl)
		}
	} else if request.Method == http.MethodGet {

		shortUrl := baseUrl + request.URL.String()

		for k, v := range linkListShorts {
			if k == shortUrl {
				writer.Header().Set("Location", v)
				writer.WriteHeader(http.StatusTemporaryRedirect)
				return
			}
		}

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
