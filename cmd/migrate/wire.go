//go:build wireinject
// +build wireinject

package migrate

import (
	"github.com/google/wire"
	"github.com/mjiee/grf-gin/app/pkg/conf"
	"github.com/mjiee/grf-gin/app/pkg/db"
)

func initData(confFile string) (data, func(), error) {
	wire.Build(newData, conf.NewConfig, db.DbSet)

	return data{}, nil, nil
}
