package app

import (
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_GetShortURL(t *testing.T) {

	type want struct {
		code        int
		contentType string
		location    string
	}

	tests := []struct {
		name   string
		method string
		body   string
		want   want
	}{
		{
			name:   "test#1-ok",
			method: http.MethodPost,
			body:   "https://yandex.ru",
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
		{
			name:   "test#2-fail",
			method: http.MethodPost,
			body:   "",
			want: want{
				code:        400,
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	conf := &config.AppConfig{
		Host:      "localhost:8080",
		ResultURL: "http://localhost:8080",
		FilePATH:  "/tmp/short-url-db.json",
	}

	app := NewApp(conf)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.method, "https://localhost:8080", strings.NewReader(test.body))

			w := httptest.NewRecorder()

			app.GetShortURL(w, r)

			assert.Equal(t, test.want.code, w.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, test.want.contentType, w.Header().Get("Content-Type"), "Тип контента не совпадает с ожидаемым")
		})
	}
}
