package cookies

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/al-kirpichenko/shortlinks/internal/services/jwtstringbuilder"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

type ContextKey string

const ContextUserKey ContextKey = "token"

func createToken() (string, error) {

	userID := uuid.New().String()
	cookieString, err := jwtstringbuilder.BuildJWTSting(userID)
	if err != nil {
		log.Println("Don't create cookie string")
		return "", err
	}
	return cookieString, nil

}

func Cookies(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var token string

		userCookie, err := r.Cookie("token")

		if err != nil {

			token, err = createToken()

			if err != nil {
				logger.ZapLogger.Error("Don't create cookie string", zap.Error(err))
			}
			userCookie = setCookie(w, token)

		}

		if _, err = userid.GetUserID(userCookie.Value); err != nil {

			http.Error(w, "user id not found in cookie", http.StatusUnauthorized)
			return

		}

		if !userid.ValidationToken(userCookie.Value) {

			logger.ZapLogger.Error("userCookie is not valid")
			token, err = createToken()

			if err != nil {
				logger.ZapLogger.Error("Don't create cookie string")
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
