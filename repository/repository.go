package repository

import (
	"go-service/repository/product"
	"go-service/repository/user"

	"gorm.io/gorm"
)

type Repository struct {
	User    user.UserRepository
	Product product.ProductRepository
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		User:    user.NewUserRepository(db),
		Product: product.NewProductRepository(db),
	}
}
