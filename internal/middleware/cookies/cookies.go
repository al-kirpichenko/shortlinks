package cookies

import (
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/services/JWTStringBuilder"
	"github.com/al-kirpichenko/shortlinks/internal/services/cookie"
)

func Cookies(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			zap.L().Error("cookies not found")
			setCookie(w)
		}
		if _, err := cookie.GetUserID(token.String()); err != nil {
			zap.L().Error("cookie is not valid")
			setCookie(w)
		}

		h.ServeHTTP(w, r)
	})
}

func setCookie(w http.ResponseWriter) {
	userID := uuid.New().String()
	cookieString, err := JWTStringBuilder.BuildJWTSting(userID)
	if err != nil {
		zap.L().Error("Don't create cookie string")
	}

	newCookie := &http.Cookie{
		Name:     "token",
		Value:    cookieString,
		MaxAge:   10800,
		HttpOnly: true,
	}
	http.SetCookie(w, newCookie)
}
