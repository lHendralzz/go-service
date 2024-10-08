package service

import (
	"go-service/repository"
	"go-service/service/order"
	"go-service/service/product"
	"go-service/service/user"
	"go-service/stdlib/redis"

	log "github.com/sirupsen/logrus"
)

// Usecase usecase struct
type Service struct {
	User    user.UserService
	Product product.ProductService
	Order   order.OrderService
}

type Options struct {
	User    user.Option
	Product product.Option
	Order   order.Option
}

func Init(repo *repository.Repository, logger *log.Logger, option Options, redis *redis.RedisLock) *Service {
	return &Service{
		User:    user.NewUserService(repo.User, logger, option.User),
		Product: product.NewProductService(repo.Product, logger, option.Product),
		Order:   order.NewOrderService(repo.Order, logger, option.Order, redis),
	}
}
