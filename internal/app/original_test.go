package app

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/cookies"
	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/jwtstringbuilder"

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
	app.ConfigureStorage()

	userID := uuid.New().String()

	cookieString, err := jwtstringbuilder.BuildJWTSting(userID)
	if err != nil {
		log.Println("Don't create cookie string")
	}

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

			link := &models.Link{
				Short:    test.body,
				Original: resultURL,
			}

			if err := app.Storage.Insert(link); err != nil {
				log.Println(err)
			}

			r := httptest.NewRequest(test.method, "http://localhost:8080/"+test.body, nil)

			ctx := context.WithValue(r.Context(), cookies.ContextUserKey, cookieString)

			w := httptest.NewRecorder()

			router := chi.NewRouteContext()

			router.URLParams.Add("id", test.body)

			r = r.WithContext(context.WithValue(ctx, chi.RouteCtxKey, router))

			app.GetOriginalURL(w, r)

			assert.Equal(t, test.want.code, w.Code, "Код ответа (307) не совпадает с ожидаемым")
			assert.Equal(t, test.want.location, w.Header().Get("Location"), "Location не совпадает с ожидаемым")
		})
	}
}
