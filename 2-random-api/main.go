package main

import "net/http"

func main() {
	router := http.NewServeMux()
	NewRandomSixHandler(router)

	server := http.Server{
		Addr:    ":8083",
		Handler: router,
	}
	server.ListenAndServe()
}
