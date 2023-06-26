package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

type Req struct {
	ID  string `json:"correlation_id"`
	URL string `json:"original_url"`
}

type Resp struct {
	ID    string `json:"correlation_id"`
	Short string `json:"short_url"`
}

func (a *App) APIBatch(w http.ResponseWriter, r *http.Request) {

	var (
		originals []Req
		shorts    []Resp
		links     []*models.Link
		userID    string
	)

	//token := r.Context().Value(Token).(string)
	//userID, err := userid.GetUserID(token)
	//
	//if err != nil {
	//	zap.L().Info("token is not found")
	//}

	cook, err := r.Cookie("token")

	if err != nil {
		userID = ""
	} else {
		userID, err = userid.GetUserID(cook.String())
		if err != nil {
			userID = ""
		}
	}

	err = json.NewDecoder(r.Body).Decode(&originals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, val := range originals {

		key := keygen.GenerateKey()
		resp := Resp{
			ID:    val.ID,
			Short: fmt.Sprintf(a.cfg.ResultURL+"/%s", key),
		}
		link := &models.Link{
			Short:    key,
			Original: val.URL,
			UserID:   userID,
		}

		shorts = append(shorts, resp)
		links = append(links, link)

	}

	if err := a.Storage.InsertLinks(links); err != nil {
		zap.L().Error("Don't insert URLs", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(shorts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(response)
	if err != nil {
		return
	}

}
