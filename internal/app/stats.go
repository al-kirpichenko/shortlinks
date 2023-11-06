package app

import "net/http"

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

func (a *App) APIStats(w http.ResponseWriter, r *http.Request) {

}
