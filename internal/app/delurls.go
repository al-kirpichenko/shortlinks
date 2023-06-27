package app

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/al-kirpichenko/shortlinks/internal/middleware/cookies"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

func (a *App) APIDelUserURLs(w http.ResponseWriter, r *http.Request) {

	var shorts []string

	token := r.Context().Value(cookies.ContextUserKey).(string)

	userID, err := userid.GetUserID(token)
	if err != nil {
		userID = ""
	}

	err = json.NewDecoder(r.Body).Decode(&shorts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func(urls []string, userID string) {
		if err := a.Storage.DelURL(urls, userID); err != nil {
			log.Println(err)
		}
		wg.Done()
	}(shorts, userID)
	wg.Wait()

	w.WriteHeader(http.StatusAccepted)

}
