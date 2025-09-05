package middleware

import (
	"context"
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	PhoneFromContext key = "PhoneFromContext"
)

func writeHeader(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuth(next http.Handler, config *configs.Config, j *jwt.JWTStruct) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeHeader(w)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		isValid, data := j.ParseJWT(token)
		if !isValid {
			writeHeader(w)
			return
		}
		ctx := context.WithValue(context.Background(), PhoneFromContext, data.Phone)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
