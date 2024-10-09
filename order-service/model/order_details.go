package model

type OrderDetails struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
}

func (OrderDetails) TableName() string {
	return "order_details"
}
