package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/al-kirpichenko/shortlinks/internal/middleware/cookies"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

type RespURLs struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

// APIGetUserURLs в теле запроса принимает список идентификаторов сокращённых URL для асинхронного удаления:
// формат: ["6qxTVvsy", "RTfd56hn", "Jlfd67ds"]
func (a *App) APIGetUserURLs(w http.ResponseWriter, r *http.Request) {

	var links []RespURLs

	token := r.Context().Value(cookies.ContextUserKey).(string)

	userID, err := userid.GetUserID(token)
	if err != nil {
		userID = ""
	}

	userURLs, err := a.Storage.GetAllByUserID(userID)

	if err != nil {
		http.Error(w, "ошибка получения урл", http.StatusNoContent)
		return
	}

	if len(userURLs) == 0 || userID == "" {
		http.Error(w, "урл нет или юзер не найден", http.StatusNoContent)
		return
	}

	for _, val := range userURLs {
		resp := RespURLs{
			Short:    strings.TrimSpace(fmt.Sprintf(a.cfg.ResultURL+"/%s", val.Short)),
			Original: strings.TrimSpace(val.Original),
		}

		links = append(links, resp)
		log.Println(resp)
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
