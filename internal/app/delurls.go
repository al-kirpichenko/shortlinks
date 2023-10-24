package app

import (
	"net/http"
)

func (a *App) APIDelUserURLs(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusAccepted)

	//var shorts []string
	//
	//token := r.Context().Value(cookies.ContextUserKey).(string)
	//
	//userID, err := userid.GetUserID(token)
	//if err != nil {
	//	http.Error(w, "user id not found in cookie", http.StatusUnauthorized)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusAccepted)
	//
	//err = json.NewDecoder(r.Body).Decode(&shorts)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//queue := delurls.NewQueue(a.Channel)
	//
	//queue.Push(&delurls.Task{
	//	UserID: userID,
	//	URLs:   shorts,
	//})

}
