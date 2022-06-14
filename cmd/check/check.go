package check

import (
	"github.com/go-redis/redis/v8"
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type checker struct {
	conf  *conf.Config
	log   *zap.Logger
	db    *gorm.DB
	redis *redis.Client
}

// NewChecker 初始化检查器
func newChecker(conf *conf.Config, log *zap.Logger, db *gorm.DB, redis *redis.Client) checker {
	return checker{conf, log, db, redis}
}
