package app

import (
	"flag"
	"github.com/al-kirpichenko/shortlinks/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_GetShortURL(t *testing.T) {

	//очищаем флаги командной строки
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

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

	conf := config.NewCfg()
	app := NewApp(conf)

	r := httptest.NewRequest(test.method, "https://localhost:8080", strings.NewReader(test.url))

	w := httptest.NewRecorder()

	app.GetShortURL(w, r)

	assert.Equal(t, test.want.code, w.Code, "Код ответа не совпадает с ожидаемым")
	assert.Equal(t, test.want.contentType, w.Header().Get("Content-Type"), "Тип контента не совпадает с ожидаемым")

}
