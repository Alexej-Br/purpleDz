// Package product
package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Manufacturer string         `json:"manufacturer"`
	Price        float32        `json:"price"`
	Images       pq.StringArray `gorm:"type:text[]" json:"image"`
}

func NewProduct(name, description string, price float32, images []string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Images:      images,
	}
}
