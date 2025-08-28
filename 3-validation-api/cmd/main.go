package main

import (
	"fmt"
	"go-adv/http/configs"
	"go-adv/http/internal/auth"
	"go-adv/http/internal/link"
	"go-adv/http/internal/verify"
	"go-adv/http/pkg/db"
	"go-adv/http/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDB(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	verify.NewEmailHandler(router, verify.EmailHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8082",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8082")
	server.ListenAndServe()
}
