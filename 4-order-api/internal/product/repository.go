// Package product
package product

import (
	"go-adv/4-order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	DB *db.DB
}

func NewProductRepository(database *db.DB) *ProductRepository {
	return &ProductRepository{
		DB: database,
	}
}

func (repo *ProductRepository) GetByID(id uint) (*Product, error) {
	var product Product
	result := repo.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) DeleteByID(id uint) error {
	result := repo.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return nil
	}
	return nil
}

func (repo *ProductRepository) UpdateByID(product *Product) (*Product, error) {
	result := repo.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) GetAll() (*[]Product, error) {
	var products []Product
	result := repo.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}
