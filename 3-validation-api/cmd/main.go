package main

import (
	"fmt"
	"go-adv/http/configs"
	"go-adv/http/internal/auth"
	"go-adv/http/internal/verify"
	"go-adv/http/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDB(conf)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	verify.NewEmailHandler(router, verify.EmailHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8082")
	server.ListenAndServe()
}
