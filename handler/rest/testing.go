package restHandler

import (
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

    user, err := r.svc.User.GetUserProfile(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get user"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}