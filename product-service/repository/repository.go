package repository

import (
	"go-service/repository/product"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	Product product.ProductRepository
}

func Init(db *gorm.DB, logger *log.Logger) *Repository {
	return &Repository{
		Product: product.NewProductRepository(db, logger),
	}
}
