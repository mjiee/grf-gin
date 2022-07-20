package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/grf-gin/app/pkg/middleware"
)

func init() {
	authRouters = append(authRouters, ossRouter)
}

// oss操作
func ossRouter(v *gin.RouterGroup, auth *AuthApi) {
	r := v.Group("/oss").Use(middleware.JwtAuth(auth.appName, auth.jwtSrv))
	{
		r.GET("/getStsToken", auth.ossH.GetStsToken)
	}
}
