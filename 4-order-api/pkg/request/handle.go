// Package request
package request

import (
	"go-adv/4-order-api/pkg/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	res, err := Decode[T](r.Body)
	if err != nil {
		response.EncodeResponse(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	err = IsValid(res)
	if err != nil {
		response.EncodeResponse(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &res, nil
}
