package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// generateTestURL

func Test_shortenURL(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	testUrl := "https://yandex.ru"

	tests := []struct {
		name   string
		method string
		url    string
		want   want
	}{
		{
			name: "test post #1",
			url:  testUrl,
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/status", nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			shortenURL(w, r)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, res.StatusCode, test.want.code)
			// получаем и проверяем тело запроса
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(res.Body)
			//resBody, err := io.ReadAll(res.Body)

			//require.NoError(t, err)
			//assert.JSONEq(t, string(resBody), test.want.response)
			assert.Equal(t, res.Header.Get("Content-Type"), test.want.contentType)
		})
	}
}
