package usecase

import (
	"go-service/repository"
	"go-service/stdlib/redis"
	"go-service/usecase/order"

	log "github.com/sirupsen/logrus"
)

// Usecase usecase struct
type Usecase struct {
	Order order.OrderUsecase
}

type Options struct {
	Order order.Option
}

func Init(repo *repository.Repository, logger *log.Logger, option Options, redis *redis.RedisLock) *Usecase {
	return &Usecase{
		Order: order.NewOrderUsecase(repo.Order, logger, option.Order, redis),
	}
}
