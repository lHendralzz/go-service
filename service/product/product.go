package product

import (
	"go-service/model"
	productRepository "go-service/repository/product"

	log "github.com/sirupsen/logrus"
)

// ProductService defines methods to interact with the business logic.
type ProductService interface {
	GetProduct() ([]model.Product, error)
}

// Option define configuration in userService
type Option struct {
}

// userService is a concrete implementation of UserService.
type productService struct {
	productRepository productRepository.ProductRepository
	opt               Option
	logger            *log.Logger
}

// NewProductService returns a new ProductService
func NewProductService(productRepository productRepository.ProductRepository, logger *log.Logger, opt Option) ProductService {
	return &productService{
		productRepository: productRepository,
		opt:               opt,
		logger:            logger,
	}
}
