// Package auth
package auth

import (
	"go-adv/http/configs"
	req "go-adv/http/pkg/request"
	"go-adv/http/pkg/response"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		data := LoginResponse{
			Token: "123",
		}
		response.EncodeResponse(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		response.EncodeResponse(w, body, http.StatusCreated)

	}
}
