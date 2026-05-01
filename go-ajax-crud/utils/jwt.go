package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("secret123")

func GenerateToken(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": user,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(secret)
}

func ValidateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}