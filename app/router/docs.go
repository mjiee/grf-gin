//go:build dev

package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mjiee/scaffold-gin/app/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	noAuthRouters = append(noAuthRouters, swaggerRouter)
}

func swaggerRouter(v *gin.RouterGroup, noAuth *NoAuthApi) {
	v.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
