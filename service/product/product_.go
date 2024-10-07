package product

import (
	"go-service/model"

	"github.com/palantir/stacktrace"
)

func (s *productService) GetProduct() ([]model.Product, error) {
	products, err := s.productRepository.GetProduct()
	if err != nil {
		return []model.Product{}, stacktrace.Propagate(err, "Failed Get product")
	}
	return products, nil
}
