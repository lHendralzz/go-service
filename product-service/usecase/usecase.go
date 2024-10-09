package usecase

import (
	"go-service/repository"
	"go-service/stdlib/redis"
	"go-service/usecase/product"

	log "github.com/sirupsen/logrus"
)

// Usecase usecase struct
type Usecase struct {
	Product product.ProductUsecase
}

type Options struct {
	Product product.Option
}

func Init(repo *repository.Repository, logger *log.Logger, option Options, redis *redis.RedisLock) *Usecase {
	return &Usecase{
		Product: product.NewProductUsecase(repo.Product, logger, option.Product),
	}
}
