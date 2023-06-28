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

func createCookieString() (string, error) {
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

		var cookieString string

		token, err := r.Cookie("token")

		if err != nil {

			cookieString, err = createCookieString()

			if err != nil {
				logger.ZapLogger.Error("Don't create cookie string", zap.Error(err))
			}
			token = setCookie(w, cookieString)

		}

		if _, err = userid.GetUserID(token.Value); err != nil {

			http.Error(w, "user id not found in cookie", http.StatusUnauthorized)
			return

		}

		if !userid.ValidationToken(token.Value) {

			logger.ZapLogger.Error("token is not valid")
			cookieString, err = createCookieString()

			if err != nil {
				logger.ZapLogger.Error("Don't create cookie string")
			}
			setCookie(w, cookieString)
			token = setCookie(w, cookieString)

		}

		ctx := context.WithValue(r.Context(), ContextUserKey, token.Value)
		h.ServeHTTP(w, r.WithContext(ctx))

	})
}

func setCookie(w http.ResponseWriter, cookieString string) *http.Cookie {

	newCookie := &http.Cookie{
		Name:     "token",
		Value:    cookieString,
		MaxAge:   10800,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, newCookie)
	return newCookie
}
