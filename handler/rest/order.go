package restHandler

import (
	"go-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckoutOrder godoc
// @Summary checkout an Order
// @Description Endpoint to Insert an Order with status checkout will need shop_id and list of product_id and their quantity
// @Tags order
// @Accept json
// @Produce json
// @Param data body model.CheckoutOrderRequest true "Add Product"
// @Success 200 {object} model.CheckoutOrderResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /order/check-out [post]
// @Security BearerAuth
func (r *rest) CheckoutOrder(ctx *gin.Context) {
	var req model.CheckoutOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// If there's an error, return a 400 response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserID = ctx.GetInt("userID")

	r.logger.Info(req.UserID)

	// TODO checkout order
	err := r.svc.Order.CheckoutOrder(ctx.Request.Context(), req)
	if err != nil {
		r.logger.Error(err)
		r.HttpRespError(ctx, err)
		return
	}

	resp := model.CheckoutOrderResponse{
		Message: "Success Checkout An Order",
	}

	ctx.JSON(http.StatusOK, resp)
}
