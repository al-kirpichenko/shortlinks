package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

type RespURLs struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

func (a *App) APIGetUserURLs(w http.ResponseWriter, r *http.Request) {

	var links []RespURLs
	var userID string

	cook, err := r.Cookie("token")

	if err != nil {
		userID = ""
	} else {
		userID, err = userid.GetUserID(cook.Value)
		if err != nil {
			userID = ""
		}
	}

	log.Println(userID)

	userURLs, err := a.Storage.GetAllByUserID(userID)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if len(userURLs) == 0 || userID == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	for _, val := range userURLs {
		resp := RespURLs{
			Short:    strings.TrimSpace(fmt.Sprintf(a.cfg.ResultURL+"/%s", val.Short)),
			Original: strings.TrimSpace(val.Original),
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
