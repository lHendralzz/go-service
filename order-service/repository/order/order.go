package order

import (
	"go-service/model"
	"time"

	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

// OrderRepository is an interface that defines the data operations.
type OrderRepository interface {
	BeginTransaction() (*gorm.DB, error)
	AddOrderWithTx(*gorm.DB, model.CheckoutOrderRequest) (int, error)
	GetOrderWithStatusAndBeforeTime(status int, searchTime time.Time) ([]model.Order, error)
}

// orderRepository is a concrete implementation of the orderRepository interface.
type orderRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

// NewOrderRepository returns an instance of the concrete implementation of OrderRepository.
func NewOrderRepository(db *gorm.DB, logger *log.Logger) OrderRepository {
	return &orderRepository{
		db:     db,
		logger: logger,
	}
}
