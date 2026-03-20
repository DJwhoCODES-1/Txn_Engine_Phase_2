package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(secret, id string) (string, string, error) {
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	at, err := access.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	rt, err := refresh.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}
