package cookie

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"github.com/al-kirpichenko/shortlinks/internal/services/JWTStringBuilder"
)

func GetUserID(tokenString string) (string, error) {
	// создаём экземпляр структуры с утверждениями
	claims := &JWTStringBuilder.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(JWTStringBuilder.SECRET_KEY), nil
		})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("Token is not valid")
	}

	return claims.UserID, nil
}
