package restHandler

import (
	"github.com/gin-gonic/gin"
)

// Custom middleware to log request processing time
func (r *rest)AuthChecker() gin.HandlerFunc {
    return func(c *gin.Context) {
        r.logger.Info("this is checking the auth")
        // Process request
        c.Next()
    }
}