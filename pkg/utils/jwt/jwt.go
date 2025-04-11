package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateJwtToken(userId, secretKey string) (JwtTokens, error) {

	claims := jwt.MapClaims{
		"sub":    userId,
		"iss":    "go_blogs",
		"userId": userId,
		"exp":    time.Now().Add(time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"sub":    userId,
		"iss":    "go_blogs",
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"iat":    time.Now().Unix(),
	}

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	tokenString, err := access_token.SignedString([]byte(secretKey))
	refreshTokenString, refreshErr := refresh_token.SignedString([]byte(secretKey))
	if err != nil || refreshErr != nil {
		return JwtTokens{}, err
	}

	return JwtTokens{AccessToken: tokenString, RefreshToken: refreshTokenString}, nil
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
