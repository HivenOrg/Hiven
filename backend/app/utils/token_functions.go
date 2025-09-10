package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id uint, jwt_secret_key string, expires_in_hrs int) (string, error) {

	if jwt_secret_key == "" || expires_in_hrs <= 0 {
		return "", errors.New("invalid parameter passed")
	}

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Duration(expires_in_hrs) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwt_secret_key))
}
