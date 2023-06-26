package cookies

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/services/jwtstringbuilder"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

func Cookies(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token, err := r.Cookie("token")

		if err != nil {
			zap.L().Error("cookies not found")
			setCookie(w)
		} else if _, err := userid.GetUserID(token.Value); err != nil {
			zap.L().Error("user id not found")
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if !userid.ValidationToken(token.Value) {
			zap.L().Error("token is not valid")
			setCookie(w)
		}

		//ctx := context.WithValue(r.Context(), app.Token, token.Value)
		h.ServeHTTP(w, r)

	})
}

func setCookie(w http.ResponseWriter) {
	userID := uuid.New().String()

	cookieString, err := jwtstringbuilder.BuildJWTSting(userID)

	if err != nil {
		zap.L().Error("Don't create cookie string")
	}

	newCookie := &http.Cookie{
		Name:     "token",
		Value:    cookieString,
		MaxAge:   10800,
		HttpOnly: true,
	}

	log.Println(newCookie.String())

	http.SetCookie(w, newCookie)
}
