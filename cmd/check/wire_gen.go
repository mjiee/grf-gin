// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package check

import (
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
	"github.com/mjiee/scaffold-gin/app/pkg/db"
	"github.com/mjiee/scaffold-gin/app/pkg/zlog"
)

// Injectors from wire.go:

func initChecker(confFile2 string) (checker, func(), error) {
	config, err := conf.NewConfig(confFile2)
	if err != nil {
		return checker{}, nil, err
	}
	logger := zlog.NewLogger(config)
	gormDB, cleanup, err := db.NewMysql(config)
	if err != nil {
		return checker{}, nil, err
	}
	client, cleanup2 := db.NewRedis(config)
	checkChecker := newChecker(config, logger, gormDB, client)
	return checkChecker, func() {
		cleanup2()
		cleanup()
	}, nil
}
