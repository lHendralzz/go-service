package model

type GetProductResponse struct {
	Products []Product `json:"products"`
}

type AddStockProductRequest struct {
	ProductID int `json:"-" example:"1"`
	Quantity  int `json:"quantity" example:"1"`
}
type AddStockProductResponse struct {
	Message string `json:"message"`
}
