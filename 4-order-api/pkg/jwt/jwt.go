// Package jwt
package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

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
