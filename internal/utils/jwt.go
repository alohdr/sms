package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Data interface{}
	*jwt.StandardClaims
}

func GenerateToken(data interface{}) (*string, error) {
	dataToken := &Token{
		Data: data,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpiredTokenDuration).Unix(),
			Issuer:    "budiman",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dataToken)
	result, err := token.SignedString([]byte(SecretToken))
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ValidateHeader(bearerHeader string) (interface{}, error) {
	hmacSecret := []byte(SecretToken)
	token, err := jwt.Parse(bearerHeader, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, errors.New("INVALID TOKEN")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data := claims["Data"]

		return data, nil
	}
	return nil, errors.New("FAILED TO CLAIM TOKEN")
}
