package userid

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v4"

	"github.com/al-kirpichenko/shortlinks/internal/services/jwtstringbuilder"
)

// GetUserID - получает ID пользователя из токена
func GetUserID(tokenString string) (string, error) {
	// создаём экземпляр структуры с утверждениями

	claims := &jwtstringbuilder.Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(jwtstringbuilder.SecretKey), nil
		})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return claims.UserID, nil
}

// ValidationToken - валидация токена
func ValidationToken(tokenString string) bool {

	claims := &jwtstringbuilder.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(jwtstringbuilder.SecretKey), nil
		})
	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}
