// Package product
package product

import (
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/pkg/jwt"
	"go-adv/4-order-api/pkg/middleware"
	"go-adv/4-order-api/pkg/request"
	"go-adv/4-order-api/pkg/response"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
	Config            *configs.Config
	JWT               *jwt.JWTStruct
}

type ProductHandler struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("GET /product/{id}", handler.GetOne())
	router.HandleFunc("GET /product", handler.GetAll())
	router.Handle("POST /product", middleware.IsAuth(handler.Create(), deps.Config, deps.JWT))
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
	router.HandleFunc("PATCH /product/{id}", handler.Update())

}

func (handler *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := handler.ProductRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		response.EncodeResponse(w, result, http.StatusOK)
	}
}

func (handler *ProductHandler) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result, err := handler.ProductRepository.GetByID(uint(idUint))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		response.EncodeResponse(w, result, http.StatusOK)
	}
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}
		product := NewProduct(body.Name, body.Description, body.Manufacturer, body.Price, body.Images)
		createdProduct, err := handler.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.EncodeResponse(w, createdProduct, http.StatusCreated)
	}
}

func (handler *ProductHandler) Delete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.ProductRepository.GetByID(uint(idUint))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.ProductRepository.DeleteByID(uint(idUint))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.EncodeResponse(w, nil, http.StatusNoContent)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		body, err := request.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}
		createdProduct, err := handler.ProductRepository.UpdateByID(&Product{
			Model:        gorm.Model{ID: uint(idUint)},
			Name:         body.Name,
			Description:  body.Description,
			Manufacturer: body.Manufacturer,
			Price:        body.Price,
			Images:       body.Images,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.EncodeResponse(w, createdProduct, http.StatusCreated)
	}
}
