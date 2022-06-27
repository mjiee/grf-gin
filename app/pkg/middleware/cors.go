package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, UPDATE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization,Accept,Content-Type,Content-Length,Origin,X-CSRF-Token")
		c.Header("Access-Control-Expose-Headers", "Content-Type,Content-Length,Content-Language,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Cache-Control,New-Token,New-Expires-In")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
