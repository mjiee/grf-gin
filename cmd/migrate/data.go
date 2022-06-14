package migrate

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type data struct {
	db    *gorm.DB
	redis *redis.Client
}

func newData(db *gorm.DB, redis *redis.Client) data {
	return data{db, redis}
}
