package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassWord(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashPassword), err
}

func CheckPasswordHashed(password, hashedPasswod string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswod), []byte(password))
	return err == nil
}
