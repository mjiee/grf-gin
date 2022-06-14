package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
	"github.com/mjiee/scaffold-gin/app/pkg/middleware"
	"go.uber.org/zap"
)

func NewRouter(cfg *conf.Config, logger *zap.Logger) *gin.Engine {
	router := gin.New()

	router.Use(
		middleware.RecoveryWithZap(logger, true),
		middleware.ZapLogger(logger, cfg.Log.SkipPaths),
		middleware.Cors(),
	)

	return router
}
