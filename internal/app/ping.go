package app

import (
	"log"
	"net/http"
)

func (a *App) Ping(w http.ResponseWriter, _ *http.Request) {

	if a.DBReady {
		if err := a.DataBase.DB.Ping(); err != nil {
			log.Println("don't ping Database")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	log.Println("don't ping Database")
	w.WriteHeader(http.StatusInternalServerError)
}
