package model

import (
	"time"
)

type Order struct {
	ID         int
	UserID     int
	ShopID     int
	Status     int
	Created_at time.Time
}

func (Order) TableName() string {
	return "order"
}

const (
	OrderStatusCheckout = 1
)
