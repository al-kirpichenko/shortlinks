package app

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
	"github.com/al-kirpichenko/shortlinks/internal/services/sqlerror"
)

func (a *App) GetShortURL(w http.ResponseWriter, r *http.Request) {

	var status = http.StatusCreated

	responseData, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, fmt.Sprintf("cannot read request body: %s", err), http.StatusBadRequest)
		return
	}
	if string(responseData) == "" {
		http.Error(w, "Empty POST request body!", http.StatusBadRequest)
		return
	}

	link := &entities.Link{
		Short:    keygen.GenerateKey(),
		Original: string(responseData),
	}

	if link, err = a.Storage.Insert(link); err != nil {
		log.Println("Don't insert url!")
		log.Println(err)

		if sqlerror.GetSQLState(err) == "23505" {
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

	response := fmt.Sprintf(a.cfg.ResultURL+"/%s", link.Short)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	log.Println(response)

	_, err = io.WriteString(w, response)
	if err != nil {
		log.Println(err)
		return
	}
}
