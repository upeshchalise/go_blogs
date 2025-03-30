package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(userId, secretKey string) (string, error) {

	claims := jwt.MapClaims{
		"sub":    userId,
		"iss":    "go_blogs",
		"userId": userId,
		"exp":    time.Now().Add(time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ValidateJwtToken(tokenString, secretKey string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	},
	)
	if err != nil {
		return "", err
	}

	// check if token is expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", jwt.ErrTokenExpired
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["userId"].(string), nil
	}
	return "", jwt.ErrSignatureInvalid
}
