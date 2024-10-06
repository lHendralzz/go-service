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

func (r *rest) Login(ctx *gin.Context) {
    username,password := "user", "password"

    token, err := r.svc.User.Login(username, password)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get user"})
        return
    }

    ctx.JSON(http.StatusOK, model.LoginResponse{Token: token})
}