package app

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/services/ipcheker"
)

// Stat -
type Stat struct {
	URLs  int `json:"urls"`
	Users int `json:"users"`
}

// APIStats - возвращает статистику сервиса
// в виде json объекта:
//
//	{
//	 "urls": <int>, // количество сокращённых URL в сервисе
//	 "users": <int> // количество пользователей в сервисе
//	}
func (a *App) APIStats(w http.ResponseWriter, r *http.Request) {

	var (
		statURLs, statUsers int
		err                 error
	)

	if !ipcheker.CheckIP(r, a.Cfg.TrustedSubnet) {
		http.Error(w, "Досут из недоверенной сети", http.StatusForbidden)
		return
	}

	statURLs, err = a.Storage.GetCountURLs()

	if err != nil {
		http.Error(w, "Ошибка получения статистики URL", http.StatusInternalServerError)
		return
	}

	statUsers, err = a.Storage.GetCountUsers()

	if err != nil {
		http.Error(w, "Ошибка получения статистики Users", http.StatusInternalServerError)
		return
	}

	stat := &Stat{
		URLs:  statURLs,
		Users: statUsers,
	}

	response, err := json.Marshal(stat)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(response)
	if err != nil {
		zap.L().Error("error", zap.Error(err))
		return
	}

}
