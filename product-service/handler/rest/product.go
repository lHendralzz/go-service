package restHandler

import (
	"go-service/model"
	x "go-service/stdlib/error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
)

// GetProduct godoc
// @Summary Get All Product
// @Description Endpoint For Authenticate User view all of products along with their stock
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} model.GetProductResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /product [get]
// @Security BearerAuth
func (r *rest) GetProduct(ctx *gin.Context) {
	resp, err := r.svc.Product.GetProduct()
	if err != nil {
		r.logger.Error(err)
		r.HttpRespError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// AddStockProduct godoc
// @Summary Add Stock of a Product
// @Description Endpoint to add stock number of a Product
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path int true "Product ID"
// @Param data body model.AddStockProductRequest true "Add Product"
// @Success 200 {object} model.AddStockProductResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /product/{product_id}/add-stock [post]
// @Security BearerAuth
func (r *rest) AddStockProduct(ctx *gin.Context) {

	productId, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		r.HttpRespError(ctx, stacktrace.PropagateWithCode(err, x.ErrorInvalidRequest, "failed conver lender_id"))
		return
	}

	var req model.AddStockProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// If there's an error, return a 400 response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ProductID = productId

	err = r.svc.Product.AddStockProduct(req)
	if err != nil {
		r.HttpRespError(ctx, err)
		return
	}

	resp := model.AddStockProductResponse{
		Message: "Success Add Stock Product",
	}

	ctx.JSON(http.StatusOK, resp)
}
