package main

import (
	"fmt"
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/internal/product"
	"go-adv/4-order-api/pkg/db"

	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDB(config)
	router := http.NewServeMux()
	productRepository := product.NewProductRepository(db)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr:    ":8083",
		Handler: router,
	}
	fmt.Printf("Server startted on %v\n", server.Addr)
	server.ListenAndServe()
}
