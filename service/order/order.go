package order

import (
	"context"
	"go-service/model"
	orderRepository "go-service/repository/order"
	"go-service/stdlib/redis"
	"time"

	log "github.com/sirupsen/logrus"
)

// OrderService defines methods to interact with the business logic.
type OrderService interface {
	CheckoutOrder(context.Context, model.CheckoutOrderRequest) (int, error)
	ReleaseOrderFromCheckoutStatus(x time.Duration) error
}

// Option define configuration in orderService
type Option struct {
	LockDuration int // in Second
}

// orderService is a concrete implementation of OrderService.
type orderService struct {
	orderRepository orderRepository.OrderRepository
	opt             Option
	logger          *log.Logger
	redis           *redis.RedisLock
}

// NewOrderService returns a new OrderService with a repository injected.
func NewOrderService(orderRepo orderRepository.OrderRepository, logger *log.Logger, opt Option, redis *redis.RedisLock) OrderService {
	return &orderService{
		orderRepository: orderRepo,
		opt:             opt,
		logger:          logger,
		redis:           redis,
	}
}
