package app

import (
	"flag"
	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_APIGetShortURL(t *testing.T) {

	//очищаем флаги командной строки
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	type (
		want struct {
			code        int
			contentType string
			location    string
		}

		testInf struct {
			method string
			want   want
		}
	)

	test := testInf{
		method: http.MethodPost,
		want: want{
			code:        201,
			contentType: "application/json",
		},
	}

	conf := config.NewAppConfig()
	app := NewApp(conf)

	r := httptest.NewRequest(test.method, "https://localhost:8080", strings.NewReader(`{"url": "https://yandex.ru"}`))

	w := httptest.NewRecorder()

	app.APIGetShortURL(w, r)

	assert.Equal(t, test.want.code, w.Code, "Код ответа не совпадает с ожидаемым")
	assert.Equal(t, test.want.contentType, w.Header().Get("Content-Type"), "Тип контента не совпадает с ожидаемым")

}
