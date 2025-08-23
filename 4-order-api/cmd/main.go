package main

import (
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/pkg/db"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDB(config)
	router := http.NewServeMux()
	_ = db
	_ = router

	server := http.Server{
		Addr:    ":8083",
		Handler: router,
	}

	server.ListenAndServe()
}
