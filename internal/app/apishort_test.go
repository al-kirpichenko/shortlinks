package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
)

func Test_APIGetShortURL(t *testing.T) {

	type want struct {
		code        int
		contentType string
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
			body:   `{"url": "https://yandex.ru"}`,
			want: want{
				code:        201,
				contentType: "application/json",
			},
		},
		{
			name:   "test#-fail",
			method: http.MethodPost,
			body:   "sdfqwed",
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
	app.ConfigureStorage()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.method, "https://localhost:8080/api/shorten", strings.NewReader(test.body))

			w := httptest.NewRecorder()

			app.APIGetShortURL(w, r)

			assert.Equal(t, test.want.code, w.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, test.want.contentType, w.Header().Get("Content-Type"), "Тип контента не совпадает с ожидаемым")
		})
	}

}
