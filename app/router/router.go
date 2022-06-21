package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/scaffold-gin/app/api"
	"github.com/mjiee/scaffold-gin/app/lib"
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
	"github.com/mjiee/scaffold-gin/app/pkg/middleware"
	"go.uber.org/zap"
)

var (
	noAuthRouters = make([]func(*gin.RouterGroup, *NoAuthApi), 0)
	authRouters   = make([]func(*gin.RouterGroup, *AuthApi), 0)
)

type NoAuthApi struct {
	authH *api.AuthHandler
}

type AuthApi struct {
	appName string
	jwtSrv  *lib.JwtService
	userSrv *lib.UserService
	userH   *api.UserHandler
}

func NewNoAuthApi(authH *api.AuthHandler) *NoAuthApi {
	return &NoAuthApi{authH}
}

func NewAuthApi(
	cfg *conf.Config, jwtSrv *lib.JwtService,
	userSrv *lib.UserService, userH *api.UserHandler,
) *AuthApi {
	return &AuthApi{cfg.App.Name, jwtSrv, userSrv, userH}
}

func NewRouter(
	cfg *conf.Config, logger *zap.Logger,
	noAuth *NoAuthApi, auth *AuthApi,
) *gin.Engine {
	router := gin.New()

	router.Use(
		middleware.RecoveryWithZap(logger, true),
		middleware.ZapLogger(logger, cfg.Log.SkipPaths),
		middleware.Cors(),
	)

	noAuthRouter(router, noAuth)
	AuthRoter(router, auth)

	return router
}

func noAuthRouter(r *gin.Engine, noAuth *NoAuthApi) {
	v := r.Group("/api/v1")
	for _, fn := range noAuthRouters {
		fn(v, noAuth)
	}
}

func AuthRoter(r *gin.Engine, auth *AuthApi) {
	v := r.Group("/api/v1")
	for _, fn := range authRouters {
		fn(v, auth)
	}
}
