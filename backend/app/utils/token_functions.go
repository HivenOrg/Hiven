package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, jwtSecretKey string, expiresInHrs int) (string, error) {

	if jwtSecretKey == "" || expiresInHrs <= 0 {
		return "", errors.New("invalid parameter passed")
	}

	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Duration(expiresInHrs) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecretKey))
}

func ValidateJWT(tokenString string, jwtSecretKey string) (uint, error) {

	if tokenString == "" || jwtSecretKey == "" {
		return 0, errors.New("invalid parameter passed")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethod(jwt.SigningMethodHS256).Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("invalid token")
	}

	return uint(userId), nil
}
