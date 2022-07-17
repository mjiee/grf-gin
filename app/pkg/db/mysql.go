package db

import (
	"fmt"
	"log"
	"time"

	"github.com/mjiee/grf-gin/app/pkg/conf"
	"github.com/natefinch/lumberjack"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMysql(cfg *conf.Config) (*gorm.DB, func(), error) {
	dsn := cfg.Db.UserName + ":" + cfg.Db.Password + "@tcp(" + cfg.Db.Addr + ")/" +
		cfg.Db.Database + "?charset=" + cfg.Db.Charset + "&parseTime=True&loc=Local&timeout=10000ms"

	Db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // 字符串类型默认长度
		DisableDatetimePrecision:  true,  // 禁用日期时间精度
		DontSupportRenameIndex:    false, // 支持重命名索引
		DontSupportRenameColumn:   false, // 支持重命名列
		SkipInitializeWithVersion: false, // 禁用基于mysql版本配置
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止自动创建外键约束
		Logger:                                   gormLogger(cfg),
	})

	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := Db.DB()
	if err != nil {
		return nil, nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Db.MaxLifeTime) * time.Hour)
	sqlDB.SetMaxIdleConns(cfg.Db.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Db.MaxOpenConns)

	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			fmt.Println("close mysql connection error")
		}
	}

	return Db, cleanup, nil
}

// 自定义orm日志
func gormLogger(cfg *conf.Config) logger.Interface {
	var logWriter logger.Writer = log.New(&lumberjack.Logger{
		Filename:   cfg.Db.LogFile,
		MaxSize:    cfg.Log.MaxSize,
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,
		Compress:   cfg.Log.Compress,
	}, "\r\n", log.LstdFlags)

	return logger.New(logWriter, logger.Config{
		SlowThreshold:             time.Duration(cfg.Db.SlowLog) * time.Second,
		LogLevel:                  logger.LogLevel(cfg.Db.LogLevel),
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	})
}
