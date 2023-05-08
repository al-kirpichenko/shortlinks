package app

import (
	"context"
	"flag"
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_GetOriginalURL(t *testing.T) {

	//очищаем флаги командной строки
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	conf := config.GetCfg()
	app := NewApp(conf)

	testID := "vRFgdzs"

	testURL := "https://yandex.ru"

	servAddr := "http://localhost:8080"

	app.storage.SetURL(testID, testURL)

	r := httptest.NewRequest(http.MethodGet, servAddr+"/"+testID, nil)

	w := httptest.NewRecorder()

	router := chi.NewRouteContext()

	router.URLParams.Add("id", testID)

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, router))

	app.GetOriginalURL(w, r)

	assert.Equal(t, 307, w.Code, "Код ответа (307) не совпадает с ожидаемым")
	assert.Equal(t, testURL, w.Header().Get("Location"), "Location не совпадает с ожидаемым")

}
