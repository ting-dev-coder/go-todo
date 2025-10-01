package middleware

import (
	"gin-todo/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(strings.ToLower(auth), "bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized attempt"})
			return
		}

		token := strings.TrimSpace(auth[len("Bearer "):])

		_, err := utils.ParseToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}

		c.Next()
	}
}
