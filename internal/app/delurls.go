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

	for _, url := range shorts {
		wg.Add(1)
		go func(url string, userID string) {
			if err := a.Storage.DelURL(url, userID); err != nil {
				log.Println(url + " не был удален!")
			}
			wg.Done()
		}(url, userID)
	}
	wg.Wait()

	w.WriteHeader(http.StatusAccepted)

}
