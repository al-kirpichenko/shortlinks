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

	resultURL := "https://yandex.ru"

	type want struct {
		code     int
		location string
	}

	type testInf struct {
		method string
		url    string
		testID string
		want   want
	}

	test := testInf{
		method: http.MethodGet,
		url:    "http://localhost:8080/vRFgdzs",
		testID: "vRFgdzs",
		want: want{
			code:     307,
			location: resultURL,
		},
	}

	app.storage.SetURL(test.testID, resultURL)

	r := httptest.NewRequest(test.method, test.url, nil)

	w := httptest.NewRecorder()

	router := chi.NewRouteContext()

	router.URLParams.Add("id", test.testID)

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, router))

	app.GetOriginalURL(w, r)

	assert.Equal(t, test.want.code, w.Code, "Код ответа (307) не совпадает с ожидаемым")
	assert.Equal(t, test.want.location, w.Header().Get("Location"), "Location не совпадает с ожидаемым")

}
