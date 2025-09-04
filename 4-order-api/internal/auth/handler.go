package auth

import (
	"errors"
	"fmt"
	"go-adv/4-order-api/pkg/jwt"
	"go-adv/4-order-api/pkg/request"
	"go-adv/4-order-api/pkg/response"
	"net/http"
)

type UserHandlerDeps struct {
	UserRepository *UserRepository
	JWT            *jwt.JWTStruct
}

type UserHandler struct {
	UserRepository *UserRepository
	JWT            *jwt.JWTStruct
}

func NewUserHandler(router *http.ServeMux, deps UserHandlerDeps) {
	handler := &UserHandler{
		UserRepository: deps.UserRepository,
		JWT:            deps.JWT,
	}
	router.HandleFunc("POST /auth", handler.Create())
	router.HandleFunc("POST /auth/verify", handler.Verify())

}

func (handler *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[AuthRequest](&w, r)
		if err != nil {
			return
		}
		resp := NewAuthResponse(*NewAuthPreResponse(CreateSession()), CreateCode())
		_, err = handler.UserRepository.FindByPhone(body.Phone)
		if err == nil {
			handler.UserRepository.Update(NewUser(body.Phone, resp.SessionID, resp.Code))
			response.EncodeResponse(w, resp.AuthPreResponse, http.StatusAccepted)
			return
		}
		user := NewUser(body.Phone, resp.SessionID, resp.Code)
		err = handler.UserRepository.NewUser(user)
		if err != nil {
			response.EncodeResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.EncodeResponse(w, resp.AuthPreResponse, http.StatusAccepted)
	}
}

func (handler *UserHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[AuthResponse](&w, r)
		if err != nil {
			return
		}
		isNorm, _ := handler.UserRepository.FindBySession(body.SessionID)
		if isNorm == nil {
			http.Error(w, errors.New("not find by session").Error(), http.StatusBadGateway)
			return
		}
		if body.Code != isNorm.Code {
			http.Error(w, errors.New("incorrect code").Error(), http.StatusBadGateway)
			return
		}
		token, err := handler.JWT.CreateJWT(isNorm.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		data := make(map[string]string, 1)
		data["token"] = token
		response.EncodeResponse(w, data, http.StatusAccepted)
	}
}
