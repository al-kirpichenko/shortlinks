package app

import (
	"log"
	"net/http"
)

func (a *App) Ping(w http.ResponseWriter, _ *http.Request) {

	_, err := confDB(a.GetConfig().DataBaseString)

	if err != nil {
		log.Println("don't ping Database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
