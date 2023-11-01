// Package jwtstringbuilder отвечает за создание токена авторизации
package jwtstringbuilder

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims -
type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

// TokenExp - время жизни токена
const TokenExp = time.Hour * 3

// SecretKey - секретный ключ
const SecretKey = "bvaEFBtr5e"

// BuildJWTSting - создание токена
func BuildJWTSting(uuid string) (string, error) {

	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		// собственное утверждение
		UserID: uuid,
	})

	// создаём строку токена
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	// возвращаем строку токена
	return tokenString, nil
}
