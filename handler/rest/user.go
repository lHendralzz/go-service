package restHandler

import (
	"go-service/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func (r *rest) Testing(ctx *gin.Context) {
	idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    ctx.JSON(http.StatusOK, id)
}

// Login godoc
// @Summary Login
// @Description Endpoint For Login user
// @Tags user
// @Accept json
// @Produce json
// @Param data body model.LoginRequest true "Login Request"
// @Success 200 {object} model.LoginResponse
// @Failure 500 {object} string
// @Router /login [post]
func (r *rest) Login(ctx *gin.Context) {
    username,password := "user", "password"

    token, err := r.svc.User.Login(username, password)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, "cannot login")
        return
    }

    ctx.JSON(http.StatusOK, model.LoginResponse{Token: token})
}