// Package product
package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string         `json:"name" validate:"required"`
	Description  string         `json:"description"`
	Manufacturer string         `json:"manufacturer"`
	Price        float32        `json:"price"`
	Images       pq.StringArray `gorm:"type:text[]" json:"images"`
}

func NewProduct(name, description, manufacturer string, price float32, images []string) *Product {
	return &Product{
		Name:         name,
		Description:  description,
		Manufacturer: manufacturer,
		Price:        price,
		Images:       images,
	}
}
