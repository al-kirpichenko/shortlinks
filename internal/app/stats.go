package app

import (
	"encoding/json"
	"net/http"
)

type Stat struct {
	URLs  int `json:"urls"`
	Users int `json:"users"`
}

// APIStats - возвращает статистику сервиса
// в виде json объекта:
// {
//  "urls": <int>, // количество сокращённых URL в сервисе
//  "users": <int> // количество пользователей в сервисе
// }

func (a *App) APIStats(w http.ResponseWriter, _ *http.Request) {

	var (
		statURLs, statUsers int
		err                 error
	)

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
		return
	}

}
