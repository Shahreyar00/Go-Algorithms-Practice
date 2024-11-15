package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Printf("Incoming request: %s %s\n", c.Request.Method, c.Request.URL.Path)

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		fmt.Printf("Request processed: %s %s - Status: %d - Duration: %v\n", c.Request.Method, c.Request.URL.Path, statusCode, duration)
	}
}
