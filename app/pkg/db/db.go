package db

import "github.com/google/wire"

var DbSet = wire.NewSet(NewMysql, NewRedis)
