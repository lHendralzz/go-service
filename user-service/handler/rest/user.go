package restHandler

import (
	"go-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login
// @Description Endpoint For Login user
// @Tags user
// @Accept json
// @Produce json
// @Param data body model.LoginRequest true "Login Request"
// @Success 200 {object} model.LoginResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /login [post]
func (r *rest) Login(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// If there's an error, return a 400 response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := r.svc.User.Login(req.Email, req.Password)
	if err != nil {
		r.logger.Error(err)
		r.HttpRespError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, model.LoginResponse{Token: token})
}
