package app

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

func (a *App) GetShortURL(w http.ResponseWriter, r *http.Request) {

	var (
		status = http.StatusCreated
		userID string
	)

	cook, err := r.Cookie("token")

	if err != nil {
		userID = ""
	} else {
		userID, err = userid.GetUserID(cook.Value)
		if err != nil {
			userID = ""
		}
	}

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
		UserID:   userID,
	}

	if err = a.Storage.Insert(link); err != nil {
		if errors.Is(err, storage.ErrConflict) {
			link, err = a.Storage.GetShort(link.Original)
			if err != nil {
				zap.L().Error("Don't get short URL", zap.Error(err))
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			status = http.StatusConflict
		} else {
			zap.L().Error("Don't insert URL", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	response := strings.TrimSpace(fmt.Sprintf(a.cfg.ResultURL+"/%s", link.Short))
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)

	_, err = io.WriteString(w, response)
	if err != nil {
		zap.L().Error("Don't write response", zap.Error(err))
		return
	}
}
