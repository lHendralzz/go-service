package service

import (
	"go-service/repository"
	"go-service/service/product"
	"go-service/service/user"

	log "github.com/sirupsen/logrus"
)

// Usecase usecase struct
type Service struct {
	User    user.UserService
	Product product.ProductService
}

type Options struct {
	User    user.Option
	Product product.Option
}

func Init(repo *repository.Repository, logger *log.Logger, option Options) *Service {
	return &Service{
		User:    user.NewUserService(repo.User, logger, option.User),
		Product: product.NewProductService(repo.Product, logger, option.Product),
	}
}
