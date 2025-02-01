package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(secretKey string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
