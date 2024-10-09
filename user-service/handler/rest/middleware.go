package restHandler

import (
	"go-service/stdlib/auth"
	"net/http"
	"strings"

	x "go-service/stdlib/error"

	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
)

// Custom middleware to log request processing time
func (r *rest) AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Check if the token format is correct (e.g., "Bearer token_value")
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || (tokenParts[0] != "Bearer" && tokenParts[0] != "bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is invalid"})
			c.Abort()
			return
		}

		// Get the actual token from the header
		tokenString := tokenParts[1]

		claims, err := auth.ValidateToken(tokenString, r.opt.JWTKey)
		if err != nil {
			r.logger.Error(err)
			r.HttpRespError(c, stacktrace.PropagateWithCode(err, x.ErrorUnauthorized, "invalid Token"))
			c.Abort()
			return
		}

		r.logger.Debug("welcome ", claims.Username)
		c.Set("userID", claims.UserID)
		// Process request
		c.Next()
	}
}
