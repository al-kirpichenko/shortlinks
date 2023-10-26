package app

import (
	"encoding/json"
	"net/http"

	"github.com/al-kirpichenko/shortlinks/internal/middleware/cookies"
	"github.com/al-kirpichenko/shortlinks/internal/services/delurls"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

// APIDelUserURLs в теле запроса принимает список идентификаторов сокращённых URL для асинхронного удаления:
// формат: ["6qxTVvsy", "RTfd56hn", "Jlfd67ds"]
func (a *App) APIDelUserURLs(w http.ResponseWriter, r *http.Request) {

	var shorts []string

	token := r.Context().Value(cookies.ContextUserKey).(string)

	userID, err := userid.GetUserID(token)
	if err != nil {
		http.Error(w, "user id not found in cookie", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	err = json.NewDecoder(r.Body).Decode(&shorts)
	if err != nil {
		return
	}

	a.Channel <- &delurls.Task{
		UserID: userID,
		URLs:   shorts,
	}

}
