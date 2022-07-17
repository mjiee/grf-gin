package run

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjiee/grf-gin/app/pkg/conf"
	"go.uber.org/zap"
)

type app struct {
	addr    string
	httpSrv *http.Server
	zlog    *zap.Logger
}

func newApp(cfg *conf.Config, zlog *zap.Logger, router *gin.Engine) *app {
	return &app{
		addr: cfg.App.Addr,
		httpSrv: &http.Server{
			Addr:    cfg.App.Addr,
			Handler: router,
		},
		zlog: zlog,
	}
}
