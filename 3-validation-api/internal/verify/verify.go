// Package verify
package verify

import (
	"encoding/json"
	"go-adv/http/configs"
	"go-adv/http/internal/auth"
	"go-adv/http/pkg"
	"go-adv/http/pkg/request"
	"go-adv/http/pkg/response"
	"net/http"
	"os"
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
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
	router.HandleFunc("POST /send", handler.Send())
}

func (handler *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		var vf []pkg.VerifyFile
		file, err := os.ReadFile("verify.json")
		if err != nil {
			response.EncodeResponse(w, err, http.StatusInternalServerError)
		}
		err = json.Unmarshal(file, &vf)
		if err != nil {
			response.EncodeResponse(w, err, http.StatusInternalServerError)
		}
		for i, v := range vf {
			if v.Hash == hash {
				vf = append(vf[:i], vf[i+1:]...)
				content, _ := json.Marshal(vf)
				err = os.WriteFile("verify.json", content, 0644)
				if err != nil {
					response.EncodeResponse(w, err, http.StatusInternalServerError)
				}
				response.EncodeResponse(w, "True", http.StatusOK)
				return
			}
		}

		err = response.EncodeResponse(w, "False", http.StatusNotFound)
		if err != nil {
			response.EncodeResponse(w, err, http.StatusInternalServerError)
		}
	}
}

func (handler *EmailHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := request.HandleBody[auth.VerifyResponse](&w, r)
		if err != nil {
			return
		}
		hash, err := pkg.MakeHash(res.Email)
		if err != nil {
			response.EncodeResponse(w, err, http.StatusInternalServerError)
		}
		urlString, err := pkg.MakeFile(res.Email, hash)
		if err != nil {
			response.EncodeResponse(w, err, http.StatusInternalServerError)
		}
		err = pkg.SendEmail(configs.LoadConfig(), res.Email, "UUUU", urlString)
		if err != nil {
			response.EncodeResponse(w, err, http.StatusInternalServerError)
		}
	}
}
