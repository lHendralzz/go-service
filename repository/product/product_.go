package product

import (
	"go-service/model"
	x "go-service/stdlib/error"

	"github.com/palantir/stacktrace"
)

func (r *productRepository) GetProduct() ([]model.Product, error) {
	var products []model.Product

	err := r.db.Find(&products).Error
	if err != nil {
		return products, stacktrace.PropagateWithCode(err, x.ErrorQuery, "Failed Select product")
	}

	return products, nil
}
