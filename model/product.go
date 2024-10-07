package model

type Product struct {
	ID     int
	Name   string
	Price  int
	Stock  int
	ShopID int
}

func (Product) TableName() string {
	return "product"
}
