package app

import (
	"net/http"

	"go.uber.org/zap"
)

func (a *App) Ping(w http.ResponseWriter, _ *http.Request) {

	err := a.DB.PingDB()

	if err != nil {
		zap.L().Error("Don't ping Database", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
