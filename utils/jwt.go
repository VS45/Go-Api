package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Not Verified")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("Wrong Token")
	}
	tokenIsValid := parseToken.Valid
	if !tokenIsValid {
		return 0, errors.New("Invalid Token")
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token Claims")
	}
	//email:=claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
