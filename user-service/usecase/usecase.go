package usecase

import (
	"go-service/repository"
	"go-service/stdlib/redis"
	"go-service/usecase/order"
	"go-service/usecase/product"
	"go-service/usecase/user"

	log "github.com/sirupsen/logrus"
)

// Usecase usecase struct
type Usecase struct {
	User    user.UserUsecase
	Product product.ProductUsecase
	Order   order.OrderUsecase
}

type Options struct {
	User    user.Option
	Product product.Option
	Order   order.Option
}

func Init(repo *repository.Repository, logger *log.Logger, option Options, redis *redis.RedisLock) *Usecase {
	return &Usecase{
		User:    user.NewUserUsecase(repo.User, logger, option.User),
		Product: product.NewProductUsecase(repo.Product, logger, option.Product),
		Order:   order.NewOrderUsecase(repo.Order, logger, option.Order, redis),
	}
}
