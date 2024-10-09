package product

import (
	"fmt"
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

func (r *productRepository) AddStockProduct(param model.AddStockProductRequest) error {

	err := r.db.Exec(QueryAddStockProduct, param.Quantity, param.ProductID).Error
	if err != nil {
		fmt.Println(err.Error())
		return stacktrace.PropagateWithCode(err, x.ErrorQuery, "Failed Update Product")
	}
	return nil
}
