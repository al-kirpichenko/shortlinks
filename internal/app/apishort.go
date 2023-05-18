package app

import (
	"bytes"
	"encoding/json"
	"github.com/al-kirpichenko/shortlinks/internal/shortlinkgen"
	"net/http"
)

type Request struct {
	Url string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

func (a *App) ApiGetShortURL(w http.ResponseWriter, r *http.Request) {

	var req Request
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// десериализуем JSON
	if err = json.Unmarshal(buf.Bytes(), &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := shortlinkgen.GenerateID()

	a.storage.SetURL(id, req.Url)

	result := Response{
		Result: a.cfg.ResultURL + "/" + id,
	}

	response, err := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(response)
	if err != nil {
		return
	}

}
