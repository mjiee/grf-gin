package router

import (
	"github.com/gin-gonic/gin"
)

func init() {
	noAuthRouters = append(noAuthRouters, registerAndLogin)
}

// 用户注册和登录
func registerAndLogin(v *gin.RouterGroup, noAuth *NoAuthApi) {
	r := v.Group("/auth")
	{
		r.POST("/register", noAuth.authH.Register)
		r.GET("/login", noAuth.authH.Login)
		r.GET("/renewToken", noAuth.authH.RenewToken)
	}
}
