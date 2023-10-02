package app

import (
	"net/http"

	"go.uber.org/zap"
)

// Ping - эндпоинт, который при запросе проверяет соединение с базой данных.
// При успешной проверке хендлер должен вернуть HTTP-статус 200 OK, при неуспешной — 500 Internal Server Error
func (a *App) Ping(w http.ResponseWriter, _ *http.Request) {

	err := a.DB.PingDB()

	if err != nil {
		zap.L().Error("Don't ping Database", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
