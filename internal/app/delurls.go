package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/al-kirpichenko/shortlinks/internal/middleware/cookies"
	"github.com/al-kirpichenko/shortlinks/internal/services/delurls"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

func (a *App) APIDelUserURLs(w http.ResponseWriter, r *http.Request) {

	var shorts []string

	token := r.Context().Value(cookies.ContextUserKey).(string)

	userID, err := userid.GetUserID(token)
	if err != nil {
		http.Error(w, "user id not found in cookie", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	if <-r.Context().Done(); true {
		log.Println("Получен сигнал стоп")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&shorts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queue := delurls.NewQueue(a.Channel)

	queue.Push(&delurls.Task{
		UserID: userID,
		URLs:   shorts,
	})
}
