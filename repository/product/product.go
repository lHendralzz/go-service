package product

import (
	"go-service/model"

	"gorm.io/gorm"
)

// ProductRepository is an interface that defines the data operations.
type ProductRepository interface {
	GetProduct() ([]model.Product, error)
	AddStockProduct(model.AddStockProductRequest) error
}

// productRepository is a concrete implementation of the ProductRepository interface.
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository returns an instance of the concrete implementation of ProductRepository.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
