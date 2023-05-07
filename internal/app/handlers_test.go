package app

import (
	"context"
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testUrls map[string]string

func Test_GetShortURL(t *testing.T) {

	testUrls = make(map[string]string)

	type want struct {
		code        int
		contentType string
		location    string
	}
	testURL := "https://yandex.ru"

	tests := []struct {
		name   string
		method string
		url    string
		want   want
	}{
		{
			name:   "test post #1",
			url:    testURL,
			method: http.MethodPost,
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
	}

	conf := config.GetCfg()
	app := NewApp(conf)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := httptest.NewRequest(test.method, "https://localhost:8080", strings.NewReader(test.url))

			w := httptest.NewRecorder()

			app.GetShortURL(w, r)

			url := w.Body.String()

			testUrls[testURL] = url

			assert.Equal(t, test.want.code, w.Code, "Код ответа не совпадает с ожидаемым")

			r2 := httptest.NewRequest(http.MethodGet, url, nil)

			w2 := httptest.NewRecorder()

			router := chi.NewRouteContext()

			slice := strings.Split(url, "/")

			router.URLParams.Add("id", slice[:][3])

			r2 = r2.WithContext(context.WithValue(r2.Context(), chi.RouteCtxKey, router))

			app.GetOriginalURL(w2, r2)

			assert.Equal(t, 307, w2.Code, "Код ответа (307) не совпадает с ожидаемым")
			assert.Equal(t, testURL, w2.Header().Get("Location"), "Location не совпадает с ожидаемым")

		})
	}
}
