//go:build wireinject
// +build wireinject

package check

import (
	"github.com/google/wire"
	"github.com/mjiee/grf-gin/app/pkg/conf"
	"github.com/mjiee/grf-gin/app/pkg/db"
	"github.com/mjiee/grf-gin/app/pkg/zlog"
)

func initChecker(confFile string) (checker, func(), error) {
	wire.Build(newChecker, conf.NewConfig, zlog.NewLogger, db.DbSet)

	return checker{}, nil, nil
}
