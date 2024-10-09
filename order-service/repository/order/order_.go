package order

import (
	"go-service/model"
	x "go-service/stdlib/error"
	"time"

	"github.com/palantir/stacktrace"
	"gorm.io/gorm"
)

func (o *orderRepository) AddOrderWithTx(tx *gorm.DB, param model.CheckoutOrderRequest) (int, error) {

	o.logger.Info(param)

	newOrder := model.Order{
		UserID:     param.UserID,
		ShopID:     param.ShopID,
		Status:     model.OrderStatusCheckout,
		Created_at: time.Now(),
	}

	result := tx.Create(&newOrder)
	if err := result.Error; err != nil {
		return 0, stacktrace.PropagateWithCode(err, x.ErrorQuery, "Failed Create Order")
	}

	for _, product := range param.Products {
		orderDetail := model.OrderDetails{
			OrderID:   newOrder.ID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}

		err := tx.Create(&orderDetail).Error
		if err != nil {
			return 0, stacktrace.PropagateWithCode(err, x.ErrorQuery, "Failed Create Order Detail")
		}
	}

	return newOrder.ID, nil
}

func (o *orderRepository) BeginTransaction() (*gorm.DB, error) {
	tx := o.db.Begin()
	err := tx.Error
	if err != nil {
		return nil, stacktrace.PropagateWithCode(err, x.ErrorBeginTransaction, "Failed Begin Transaction")
	}

	return tx, nil
}

func (o *orderRepository) GetOrderWithStatusAndBeforeTime(status int, searchTime time.Time) ([]model.Order, error) {
	var orders []model.Order
	err := o.db.Where("created_at < ? and status = ?", searchTime, status).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}
