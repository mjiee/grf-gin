package check

import (
	"github.com/go-redis/redis/v8"
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	conf  *conf.Config
	log   *zap.Logger
	db    *gorm.DB
	redis *redis.Client
}

func NewApp(conf *conf.Config, log *zap.Logger, db *gorm.DB, redis *redis.Client) App {
	return App{conf, log, db, redis}
}
