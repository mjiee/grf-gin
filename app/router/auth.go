package router

import "github.com/gin-gonic/gin"

func init() {
	noAuthRouters = append(noAuthRouters, registerAndLogin)
}

func registerAndLogin(v *gin.RouterGroup, noAuth *NoAuthApi) {
	r := v.Group("/auth")
	{
		r.POST("/register", noAuth.authH.Register)
		r.POST("/login", noAuth.authH.Login)
	}
}
