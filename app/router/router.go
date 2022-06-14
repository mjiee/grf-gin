package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/mjiee/scaffold-gin/app/api"
	"github.com/mjiee/scaffold-gin/app/lib"
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
	"github.com/mjiee/scaffold-gin/app/pkg/middleware"
	"go.uber.org/zap"
)

// RouterSet 路由providers
var RouterSet = wire.NewSet(NewAuthApi, NewNoAuthApi, NewRouter)

// NoAuthApi 不需要jwt认证的api
type NoAuthApi struct {
	authH *api.AuthHandler
}

// NoAuthApi 需要jwt认证的api
type AuthApi struct {
	appName string
	jwtSrv  *lib.JwtService
	userSrv *lib.UserService
	userH   *api.UserHandler
}

// NewNoAuthApi 创建NoAuthApi
func NewNoAuthApi(authH *api.AuthHandler) *NoAuthApi {
	return &NoAuthApi{authH}
}

// NewAuthApi 创建AuthApi
func NewAuthApi(
	cfg *conf.Config, jwtSrv *lib.JwtService,
	userSrv *lib.UserService, userH *api.UserHandler,
) *AuthApi {
	return &AuthApi{cfg.App.Name, jwtSrv, userSrv, userH}
}

// NewRouter 创建gin路由
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

// 路由组
var (
	noAuthRouters = make([]func(*gin.RouterGroup, *NoAuthApi), 0)
	authRouters   = make([]func(*gin.RouterGroup, *AuthApi), 0)
)

// 注册不需要jwt认证路由
func noAuthRouter(r *gin.Engine, noAuth *NoAuthApi) {
	v := r.Group("/api/v1")
	for _, fn := range noAuthRouters {
		fn(v, noAuth)
	}
}

// 注册需要jwt认证路由
func AuthRoter(r *gin.Engine, auth *AuthApi) {
	v := r.Group("/api/v1")
	for _, fn := range authRouters {
		fn(v, auth)
	}
}
