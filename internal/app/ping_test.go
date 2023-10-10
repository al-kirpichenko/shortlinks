package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
)

func Test_Ping(t *testing.T) {

	type want struct {
		code int
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "test#1-ok",
			want: want{
				code: 200,
			},
		},
	}

	conf := config.NewAppConfig()

	app := NewApp(conf)

	app.ConfigureStorage()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := httptest.NewRequest(http.MethodGet, "https://localhost:8080/ping", nil)

			w := httptest.NewRecorder()

			app.Ping(w, r)

			assert.Equal(t, test.want.code, w.Code, "Код ответа не совпадает с ожидаемым")

		})
	}

}
