package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/grf-gin/app/lib"
	"github.com/mjiee/grf-gin/app/pkg/apperr"
	"github.com/mjiee/grf-gin/app/pkg/response"
)

func JwtAuth(iss string, jwtSrv *lib.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerAuth := c.Request.Header.Get("Authorization")
		claims, _, err := jwtSrv.RequestAuth(iss, headerAuth)

		if err != nil {
			response.Failure(c, apperr.TokenErr, err.Error())
			c.Abort()
			return
		}

		// if claims.IsAdmin {
		// 	c.Set("isAdmin", claims.IsAdmin)
		// }

		c.Set("id", claims.ID)
		c.Next()
	}
}
