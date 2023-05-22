package middleware

import (
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type (
	// берём структуру для хранения сведений об ответе
	responseData struct {
		status int
		size   int
	}

	// добавляем реализацию http.ResponseWriter
	loggingResponseWriter struct {
		http.ResponseWriter // встраиваем оригинальный http.ResponseWriter
		responseData        *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size // захватываем размер
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	// записываем код статуса, используя оригинальный http.ResponseWriter
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode // захватываем код статуса
}

// Logger добавляет дополнительный код для регистрации сведений о запросе
// и возвращает новый http.Handler.

func Logger(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger, err := zap.NewDevelopment()
		if err != nil {
			// вызываем панику, если ошибка
			log.Fatal(err)
		}

		defer func() {
			err = logger.Sync()
		}()

		// делаем регистратор SugaredLogger
		sugar := *logger.Sugar()

		start := time.Now()

		responseData := &responseData{}

		lw := loggingResponseWriter{
			ResponseWriter: w, // встраиваем оригинальный http.ResponseWriter
			responseData:   responseData,
		}
		h.ServeHTTP(&lw, r) // внедряем реализацию http.ResponseWriter

		sugar.Infoln(
			"uri", r.RequestURI,
			"method", r.Method,
			"status", responseData.status, // получаем перехваченный код статуса ответа
			"duration", time.Since(start),
			"size", responseData.size, // получаем перехваченный размер ответа
			"loc", w.Header().Get("Location"),
		)
	})
}
