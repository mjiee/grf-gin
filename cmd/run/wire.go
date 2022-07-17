//go:build wireinject
// +build wireinject

package run

import (
	"github.com/google/wire"
	"github.com/mjiee/grf-gin/app/api"
	"github.com/mjiee/grf-gin/app/lib"
	"github.com/mjiee/grf-gin/app/pkg/conf"
	"github.com/mjiee/grf-gin/app/pkg/db"
	"github.com/mjiee/grf-gin/app/pkg/zlog"
	"github.com/mjiee/grf-gin/app/router"
)

func initApp(confFile string) (*app, func(), error) {
	wire.Build(newApp, conf.NewConfig, zlog.NewLogger, db.DbSet, router.RouterSet,
		api.ApiSet, lib.LibSet)

	return &app{}, nil, nil
}
