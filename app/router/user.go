package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/scaffold-gin/app/pkg/middleware"
)

func init() {
	authRouters = append(authRouters, userRouter)
}

// 用户操作
func userRouter(v *gin.RouterGroup, auth *AuthApi) {
	r := v.Group("/user").Use(middleware.JwtAuth(auth.appName, auth.jwtSrv, auth.userSrv))
	{
		r.GET("/getUserInfo", auth.userH.GetUserInfo)
	}
}
