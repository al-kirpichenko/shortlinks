package app

import (
	"encoding/json"
	"net/http"

	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
)

type Req struct {
	ID  string `json:"correlation_id"`
	URL string `json:"original_url"`
}

type Resp struct {
	ID       string `json:"correlation_id"`
	Short    string `json:"short_url"`
	Original string `json:"-"`
}

func (a *App) APIBatch(w http.ResponseWriter, r *http.Request) {

	var (
		originals []Req
		shorts    []Resp
		links     []entities.Link
	)

	err := json.NewDecoder(r.Body).Decode(&originals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, val := range originals {
		short := keygen.KeyGenerate()
		resp := Resp{
			ID:    val.ID,
			Short: short,
		}
		link := entities.Link{
			Short:    short,
			Original: val.URL,
		}
		shorts = append(shorts, resp)
		links = append(links, link)
	}

}
