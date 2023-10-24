package app

import (
	"encoding/json"
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
		userID = ""
	}

	err = json.NewDecoder(r.Body).Decode(&shorts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//go func(urls []string, userID string) {
	//	if err := a.Storage.DelURL(urls, userID); err != nil {
	//		logger.ZapLogger.Error("don't delete urls", zap.Error(err))
	//	}
	//}(shorts, userID)
	queue := delurls.NewQueue(a.Channel)

	queue.Push(&delurls.Task{
		UserID: userID,
		URLs:   shorts,
	})

	w.WriteHeader(http.StatusAccepted)

}
