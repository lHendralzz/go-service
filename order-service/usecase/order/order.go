package order

import (
	"context"
	"go-service/model"
	orderRepository "go-service/repository/order"
	"go-service/stdlib/redis"
	"time"

	log "github.com/sirupsen/logrus"
)

// OrderUsecase defines methods to interact with the business logic.
type OrderUsecase interface {
	CheckoutOrder(context.Context, model.CheckoutOrderRequest) (int, error)
	ReleaseOrderFromCheckoutStatus(x time.Duration) error
}

// Option define configuration in orderUsecase
type Option struct {
	LockDuration int // in Second
}

// orderUsecase is a concrete implementation of OrderUsecase.
type orderUsecase struct {
	orderRepository orderRepository.OrderRepository
	opt             Option
	logger          *log.Logger
	redis           *redis.RedisLock
}

// NewOrderUsecase returns a new OrderUsecase with a repository injected.
func NewOrderUsecase(orderRepo orderRepository.OrderRepository, logger *log.Logger, opt Option, redis *redis.RedisLock) OrderUsecase {
	return &orderUsecase{
		orderRepository: orderRepo,
		opt:             opt,
		logger:          logger,
		redis:           redis,
	}
}
