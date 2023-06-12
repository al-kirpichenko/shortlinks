package app

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
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

	link := &models.Link{
		Short:    keygen.GenerateKey(),
		Original: string(responseData),
	}

	if err = a.Storage.Insert(link); err != nil {
		if errors.Is(err, storage.ErrConflict) {
			link, err = a.Storage.GetShort(link.Original)
			if err != nil {
				log.Println("Don't read data from table")
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			status = http.StatusConflict
		} else {
			log.Println("Don't insert url!")
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	response := strings.TrimSpace(fmt.Sprintf(a.cfg.ResultURL+"/%s", link.Short))
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)

	_, err = io.WriteString(w, response)
	if err != nil {
		log.Println(err)
		return
	}
}
