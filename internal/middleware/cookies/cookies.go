package cookies

import (
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/al-kirpichenko/shortlinks/internal/services/jwtstringbuilder"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

type ContextKey string

const ContextUserKey ContextKey = "token"

func createToken() (string, error) {

	userID := uuid.New().String()
	token, err := jwtstringbuilder.BuildJWTSting(userID)
	if err != nil {
		return "", err
	}
	return token, nil

}

func Cookies(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var token string

		userCookie, err := r.Cookie("token")

		if err != nil {

			token, err = createToken()

			if err != nil {
				http.Error(w, "Try later..", http.StatusInternalServerError)
				return
			}
			userCookie = setCookie(w, token)

		}

		if _, err = userid.GetUserID(userCookie.Value); err != nil {

			http.Error(w, "user id not found in cookie", http.StatusUnauthorized)
			return

		}

		if !userid.ValidationToken(userCookie.Value) {

			logger.ZapLogger.Error("token is not valid")
			token, err = createToken()

			if err != nil {
				http.Error(w, "Try later..", http.StatusInternalServerError)
				return
			}
			setCookie(w, token)
			userCookie = setCookie(w, token)

		}

		ctx := context.WithValue(r.Context(), ContextUserKey, userCookie.Value)
		h.ServeHTTP(w, r.WithContext(ctx))

	})
}

func setCookie(w http.ResponseWriter, token string) *http.Cookie {

	newCookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   10800,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, newCookie)
	return newCookie

}
