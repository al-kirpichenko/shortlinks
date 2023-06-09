package app

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/entities"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_GetOriginalURL(t *testing.T) {

	conf := &config.AppConfig{
		Host:      "localhost:8080",
		ResultURL: "http://localhost:8080",
		FilePATH:  "/tmp/short-url-db.json",
	}

	app := NewApp(conf)

	resultURL := "https://yandex.ru"

	type want struct {
		code     int
		location string
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
			body:   "vRFgdzs",
			want: want{
				code:     307,
				location: resultURL,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			link := entities.Link{
				Short:    test.body,
				Original: resultURL,
			}

			if err := app.Storage.Insert(link); err != nil {
				log.Println(err)
			}

			r := httptest.NewRequest(test.method, "http://localhost:8080/"+test.body, nil)

			w := httptest.NewRecorder()

			router := chi.NewRouteContext()

			router.URLParams.Add("id", test.body)

			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, router))

			app.GetOriginalURL(w, r)

			assert.Equal(t, test.want.code, w.Code, "Код ответа (307) не совпадает с ожидаемым")
			assert.Equal(t, test.want.location, w.Header().Get("Location"), "Location не совпадает с ожидаемым")
		})
	}
}
