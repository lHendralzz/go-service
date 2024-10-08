package model

type CheckoutOrderRequest struct {
	ShopID   int                           `json:"shop_id" example:"1"`
	UserID   int                           `json:"-"`
	Products []ProductCheckoutOrderRequest `json:"products"`
}

type ProductCheckoutOrderRequest struct {
	ProductID int `json:"product_id" example:"1"`
	Quantity  int `json:"quanttity" example:"1"`
}
type CheckoutOrderResponse struct {
	OrderID int    `json:"order_id"`
	Message string `json:"message"`
}
