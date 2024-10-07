package restHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
