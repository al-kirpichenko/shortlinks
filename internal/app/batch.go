package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/middleware/cookies"
	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

// Req - структура запроса
type Req struct {
	ID  string `json:"correlation_id"`
	URL string `json:"original_url"`
}

// Resp - структура ответа
type Resp struct {
	ID    string `json:"correlation_id"`
	Short string `json:"short_url"`
}

// APIBatch - хендлер, принимающий в теле запроса множество URL для сокращения в формате:
// [
//
//	{
//	    "correlation_id": "<строковый идентификатор>",
//	    "original_url": "<URL для сокращения>"
//	},
//
// ]
// В качестве ответа хендлер должен возвращает данные в формате:
// [
// {
// "correlation_id": "<строковый идентификатор из объекта запроса>",
// "short_url": "<результирующий сокращённый URL>"
// },
// ]
func (a *App) APIBatch(w http.ResponseWriter, r *http.Request) {

	var (
		originals []Req
		shorts    []Resp
		links     []*models.Link
		userID    string
	)

	token := r.Context().Value(cookies.ContextUserKey).(string)

	userID, err := userid.GetUserID(token)
	if err != nil {
		userID = ""
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
			Short: fmt.Sprintf(a.Cfg.ResultURL+"/%s", key),
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
