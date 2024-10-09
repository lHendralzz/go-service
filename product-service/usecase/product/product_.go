package product

import (
	"go-service/model"

	"github.com/palantir/stacktrace"
)

func (s *productUsecase) GetProduct() ([]model.Product, error) {
	products, err := s.productRepository.GetProduct()
	if err != nil {
		return []model.Product{}, stacktrace.Propagate(err, "Failed Get product")
	}
	return products, nil
}

func (s *productUsecase) AddStockProduct(param model.AddStockProductRequest) error {
	// TODO : validate if param.Quantity minus then cannot less than 0
	err := s.productRepository.AddStockProduct(param)
	if err != nil {
		return stacktrace.Propagate(err, "Failed AddStock Product")
	}
	return nil
}
