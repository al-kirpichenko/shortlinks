package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testUrls map[string]string

func Test_shortenURL(t *testing.T) {
	urls = make(map[string]string)

	testUrls = make(map[string]string)
	type want struct {
		code        int
		contentType string
		location    string
	}
	testUrl := "https://yandex.ru"

	tests := []struct {
		name   string
		method string
		url    string
		want   want
	}{
		{
			name:   "test post #1",
			url:    testUrl,
			method: http.MethodPost,
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.method, "https://localhost:8080", strings.NewReader(test.url))

			w := httptest.NewRecorder()

			shortenURL(w, r)

			url := w.Body.String()

			testUrls[testUrl] = url

			assert.Equal(t, test.want.code, w.Code, "Код ответа не совпадает с ожидаемым")

			t.Log("w.Body" + url)

			r2 := httptest.NewRequest(http.MethodGet, url, nil)

			w2 := httptest.NewRecorder()

			shortenURL(w2, r2)

			//t.Log("w2.Location" + w2.Header().Get("Location"))

			assert.Equal(t, 307, w2.Code, "Код ответа (307) не совпадает с ожидаемым")
			assert.Equal(t, testUrl, w2.Header().Get("Location"), "Location не совпадает с ожидаемым")

		})
	}
}
