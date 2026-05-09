package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(hashed), err
}

func IsPasswordMatch(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err != nil
}
