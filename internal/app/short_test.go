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

	//очищаем флаги командной строки
	//flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	type want struct {
		code        int
		contentType string
		location    string
	}
	testURL := "https://yandex.ru"

	type testInf struct {
		method string
		url    string
		want   want
	}

	test := testInf{
		method: http.MethodPost,
		url:    testURL,
		want: want{
			code:        201,
			contentType: "text/plain",
		},
	}

	conf := &config.AppConfig{
		Host:      "localhost:8080",
		ResultURL: "http://localhost:8080",
		FilePATH:  "/tmp/short-url-db.json",
	}
	app := NewApp(conf)

	r := httptest.NewRequest(test.method, "https://localhost:8080", strings.NewReader(test.url))

	w := httptest.NewRecorder()

	app.GetShortURL(w, r)

	assert.Equal(t, test.want.code, w.Code, "Код ответа не совпадает с ожидаемым")
	assert.Equal(t, test.want.contentType, w.Header().Get("Content-Type"), "Тип контента не совпадает с ожидаемым")

}
