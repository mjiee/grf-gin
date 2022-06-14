//go:build wireinject
// +build wireinject

package check

import (
	"github.com/google/wire"
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
	"github.com/mjiee/scaffold-gin/app/pkg/db"
	"github.com/mjiee/scaffold-gin/app/pkg/zlog"
)

func NewChecker(confFile string) (App, func(), error) {
	wire.Build(NewApp, conf.NewConfig, zlog.NewLogger, db.DbSet)

	return App{}, nil, nil
}
