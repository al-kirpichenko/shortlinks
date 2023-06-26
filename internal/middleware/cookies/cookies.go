package cookies

import (
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/services/jwtstringbuilder"
	"github.com/al-kirpichenko/shortlinks/internal/services/userid"
)

func Cookies(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")

		tokenStr := token.String()
		if err != nil {
			zap.L().Error("cookies not found")
			setCookie(w)
		}
		if _, err := userid.GetUserID(tokenStr); err != nil {
			zap.L().Error("cookie is not valid")
			setCookie(w)
		}

		ctx := context.WithValue(r.Context(), app.Token, tokenStr)
		h.ServeHTTP(w, r.WithContext(ctx))

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
	http.SetCookie(w, newCookie)
}
