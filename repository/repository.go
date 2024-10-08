package repository

import (
	"go-service/repository/order"
	"go-service/repository/product"
	"go-service/repository/user"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	User    user.UserRepository
	Product product.ProductRepository
	Order   order.OrderRepository
}

func Init(db *gorm.DB, logger *log.Logger) *Repository {
	return &Repository{
		User:    user.NewUserRepository(db, logger),
		Product: product.NewProductRepository(db, logger),
		Order:   order.NewOrderRepository(db, logger),
	}
}
