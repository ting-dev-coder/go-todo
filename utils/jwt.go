package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	userId int
	jwt.RegisteredClaims
}

var jwtKey = []byte("your_secret_key")

func ParseToken(token string) (string, error) {
	return "123", nil
}

func GenerateToken(userId int) (string, error) {
	expiredTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		userId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtKey)
}
