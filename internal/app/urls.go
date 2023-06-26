package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

type RespURLs struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

func (a *App) APIGetUserURLs(w http.ResponseWriter, r *http.Request) {

	var links []RespURLs

	token := r.Context().Value(Token).(string)

	userID, err := userid.GetUserID(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userURLs, err := a.Storage.GetAllByUserID(userID)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	for _, val := range userURLs {
		resp := RespURLs{
			Short:    fmt.Sprintf(a.cfg.ResultURL+"/%s", val.Short),
			Original: val.Original,
		}

		links = append(links, resp)
	}

	response, err := json.Marshal(links)

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
