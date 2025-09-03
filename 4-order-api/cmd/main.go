package main

import (
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/internal/auth"
	"go-adv/4-order-api/internal/product"
	"go-adv/4-order-api/pkg/db"
	"go-adv/4-order-api/pkg/jwt"
	"go-adv/4-order-api/pkg/middleware"

	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	config := configs.LoadConfig()
	middleware.SetupLogger()
	db := db.NewDB(config)
	router := http.NewServeMux()
	jwt := jwt.NewJWT(config.JWT.Secret)

	productRepository := product.NewProductRepository(db)
	userRepository := auth.NewUserRepository(db)

	auth.NewUserHandler(router, auth.UserHandlerDeps{
		UserRepository: userRepository,
		JWT:            jwt,
	})
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr:    ":8083",
		Handler: middleware.Log(router),
	}
	logrus.Infof("Server started on %v", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}
