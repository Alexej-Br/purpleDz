// Package jwt
package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Phone string
}
type JWTStruct struct {
	Secret string
}

func NewJWT(sec string) *JWTStruct {
	return &JWTStruct{Secret: sec}
}

func (j *JWTStruct) CreateJWT(phone string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"phone": phone})
	token, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JWTStruct) ParseJWT(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	phone := t.Claims.(jwt.MapClaims)["phone"]
	return t.Valid, &JWTData{Phone: phone.(string)}
}
