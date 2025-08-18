// Package verify
package verify

import (
	"fmt"
	"go-adv/http/configs"
	"go-adv/http/internal/auth"
	"go-adv/http/pkg"
	"go-adv/http/pkg/res"
	"net/http"
)

type EmailHandler struct {
	*configs.Config
}

type EmailHandlerDeps struct {
	*configs.Config
}

func NewEmailHandler(router *http.ServeMux, deps EmailHandlerDeps) {
	handler := &EmailHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pkg.SendEmail(handler.Config, "somebody", "verify", "Hello")
	}
}

func (handler *EmailHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := auth.VerifyResponse{
			Hash: "hash",
		}
		res.EncodeResponse(w, data, http.StatusCreated)
		fmt.Println("send")
	}
}
