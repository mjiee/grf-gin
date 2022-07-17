package zlog

import (
	"github.com/mjiee/grf-gin/app/pkg/conf"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg *conf.Config) *zap.Logger {
	writeSyncer := getWriteSyncer(cfg.Log.Filename, cfg.Log.Compress, cfg.Log.MaxSize,
		cfg.Log.MaxBackups, cfg.Log.MaxAge)
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.Level(cfg.Log.Level))

	return zap.New(core, zap.AddCaller())
}

// 文件writeSyncer
func getWriteSyncer(filename string, compress bool, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	})
}

// 日志编码格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
