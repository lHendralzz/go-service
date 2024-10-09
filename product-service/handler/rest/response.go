package restHandler

import (
	"net/http"

	"go-service/model"
	errorMessage "go-service/stdlib/error"

	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
)

func (r *rest) HttpRespError(ctx *gin.Context, err error) {
	r.logger.Debug(err.Error())

	errorCode := stacktrace.GetCode(err)
	if errorCode == stacktrace.NoCode {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if errMessage, ok := errorMessage.ErrorMessages[errorCode]; ok {
		ctx.JSON(errMessage.StatusCode, model.ErrorResponse{
			Message: errMessage.Message,
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, err.Error())
}
