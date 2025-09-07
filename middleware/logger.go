package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		status := c.Writer.Status()
		fmt.Printf("[%d] %s %s (%v)\n", status, c.Request.Method, c.Request.URL.Path, latency)
	}
}
