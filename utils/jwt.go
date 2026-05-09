package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key string = "super-secret"

func GenerateJwtToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(key))
}

func VerifyToke(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("wrong sign method")
		}
		return []byte(key), nil
	})

	if err != nil || !parsedToken.Valid {
		return 0, errors.New("failed to parse token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("failed to parse claims")
	}

	userId, ok := (claims["userId"]).(float64)
	if !ok {
		panic("couldnt parse")
	}

	return int64(userId), nil

}
