// Package product
package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Manufacturer string         `json:"manufacturer"`
	Price        float32        `json:"price"`
	Images       pq.StringArray `gorm:"type:text[]" json:"images"`
}
