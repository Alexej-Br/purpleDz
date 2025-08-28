package main

import (
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/internal/product"
	"go-adv/4-order-api/pkg/db"
	"go-adv/4-order-api/pkg/middleware"
	

	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	config := configs.LoadConfig()
	middleware.SetupLogger()
	db := db.NewDB(config)
	router := http.NewServeMux()

	productRepository := product.NewProductRepository(db)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr:    ":8083",
		Handler: middleware.Log(router),
	}
	logrus.Infof("Server started on %v", server.Addr)
	server.ListenAndServe()
}

