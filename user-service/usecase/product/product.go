package product

import (
	"go-service/model"
	productRepository "go-service/repository/product"

	log "github.com/sirupsen/logrus"
)

// ProductUsecase defines methods to interact with the business logic.
type ProductUsecase interface {
	GetProduct() ([]model.Product, error)
	AddStockProduct(param model.AddStockProductRequest) error
}

// Option define configuration in userUsecase
type Option struct {
}

// userUsecase is a concrete implementation of UserUsecase.
type productUsecase struct {
	productRepository productRepository.ProductRepository
	opt               Option
	logger            *log.Logger
}

// NewProductUsecase returns a new ProductUsecase
func NewProductUsecase(productRepository productRepository.ProductRepository, logger *log.Logger, opt Option) ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		opt:               opt,
		logger:            logger,
	}
}
