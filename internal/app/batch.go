package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
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
		links     []entities.Link
	)

	err := json.NewDecoder(r.Body).Decode(&originals)
	if err != nil {
		log.Println("don't decode body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, val := range originals {

		log.Println("получен url " + val.URL)
		key := keygen.KeyGenerate()
		log.Println("сгенерирована короткая ссылка " + key)
		resp := Resp{
			ID:    val.ID,
			Short: fmt.Sprintf(a.cfg.ResultURL+"/%s", key),
		}
		link := entities.Link{
			Short:    key,
			Original: val.URL,
		}

		shorts = append(shorts, resp)
		links = append(links, link)

	}

	if err := a.Storage.InsertLinks(links); err != nil {
		log.Println("Don't insert to table")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//проверка
	for _, v := range links {
		link2, err := a.Storage.GetOriginal(v.Short)
		if err != nil {
			log.Println(err)
		}
		log.Println("делаем запрос на получение оригинальных")
		log.Println(link2.Short + " ***** " + link2.Original)

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
