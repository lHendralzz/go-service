package repository

import (
	"go-service/repository/order"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	Order order.OrderRepository
}

func Init(db *gorm.DB, logger *log.Logger) *Repository {
	return &Repository{
		Order: order.NewOrderRepository(db, logger),
	}
}
