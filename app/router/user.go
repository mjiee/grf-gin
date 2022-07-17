package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/grf-gin/app/pkg/middleware"
)

func init() {
	authRouters = append(authRouters, userRouter)
}

// 用户操作
func userRouter(v *gin.RouterGroup, auth *AuthApi) {
	r := v.Group("/user").Use(middleware.JwtAuth(auth.appName, auth.jwtSrv))
	{
		r.GET("/getUserInfo", auth.userH.GetUserInfo)
	}
}
