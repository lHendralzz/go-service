package restHandler

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Custom middleware to log request processing time
func (r *rest)LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		r.logger.Infof("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)
        start := time.Now()
		rw := &responseWriter{
            ResponseWriter: c.Writer,
            body:          bytes.NewBuffer([]byte{}),
        }
        c.Writer = rw
        // Process request
        c.Next()

        // Calculate latency
        latency := time.Since(start)
        r.logger.Debug(fmt.Sprintf("Request processed in %s", latency))

        // Access the status code
        status := c.Writer.Status()
        r.logger.Info(fmt.Sprintf("Status: %d", status))
    }
}


// Custom response writer to capture the response body
type responseWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

// Write method to capture the response body
func (rw *responseWriter) Write(b []byte) (int, error) {
    rw.body.Write(b) // Capture the response body
    return rw.ResponseWriter.Write(b) // Write the response
}