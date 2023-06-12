package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgerrcode"

	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
	"github.com/al-kirpichenko/shortlinks/internal/services/sqlerror"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

func (a *App) APIGetShortURL(w http.ResponseWriter, r *http.Request) {

	var req Request
	var status = http.StatusCreated

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link := &entities.Link{
		Short:    keygen.GenerateKey(),
		Original: req.URL,
	}

	// "можно сделать меньше if-else: нашли ошибку - вернули. Продолжаем писать код зная что ошибка обработана" -
	// нельзя! в случае ошибки тут мы проверяем какая именно ошибка! если ошибка нам говорит о том, что такое значение уже есть,
	// то тогда делаем запрос для получения короткой ссылки из бд
	if link, err = a.Storage.Insert(link); err != nil {
		log.Println("Don't insert url!")
		log.Println(err)
		if sqlerror.GetSQLState(err) == pgerrcode.UniqueViolation {
			link, err = a.Storage.GetShort(link.Original)
			if err != nil {
				log.Println("Don't read data from table")
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			status = http.StatusConflict
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	result := Response{
		Result: a.cfg.ResultURL + "/" + link.Short,
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(response)
	if err != nil {
		return
	}

}
