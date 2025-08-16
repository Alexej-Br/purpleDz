package main

import (
	"fmt"
	"net/http"
)

type RandomSixHandler struct{}

func NewRandomSixHandler(router *http.ServeMux) {
	handler := &RandomSixHandler{}
	router.HandleFunc("/hello", handler.RandomSix())
}

func (handler *RandomSixHandler) RandomSix() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GHello")
		w.Write([]byte(Random()))
	}
}
