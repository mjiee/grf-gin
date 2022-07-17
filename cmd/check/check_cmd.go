package check

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mjiee/grf-gin/app/pkg/util"
	"github.com/spf13/cobra"
)

var (
	confFile string

	CheckCmd = &cobra.Command{
		Use:     "check",
		Short:   "scaffold-gin check",
		Long:    "检测项目配置和依赖服务",
		Example: "scaffold-gin check -c ./config/default.yaml",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
			// return nil
		},
	}
)

func init() {
	CheckCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "../conf/default.yaml", "提供项目配置文件")
}

// 运行检测
func run() error {
	checker, cleanup, err := initChecker(confFile)
	if err != nil {
		return err
	}

	defer cleanup()

	checkConf(checker)
	checkLogger(checker)
	checkMysql(checker)
	checkRedis(checker)

	return nil
}

// check config
func checkConf(checker checker) {
	validate := validator.New()
	validate.RegisterValidation("addr", util.CheckAddr)

	if err := validate.Struct(checker.conf); err != nil {
		fmt.Println(err.Error())
		panic("配置文件异常...")
	} else {
		fmt.Println("配置文件正常...")
	}
}

// check logger
func checkLogger(checker checker) {
	checker.log.Info("测试日志写入...")
	fmt.Println("日志写入正常...")
}

// check mysql
func checkMysql(checker checker) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sqlDb, _ := checker.db.DB()
	if err := sqlDb.PingContext(ctx); err != nil {
		panic("mysql连接测试异常...")
	} else {
		fmt.Println("mysql连接正常...")
	}
}

// check redis
func checkRedis(checker checker) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := checker.redis.Ping(ctx).Err(); err != nil {
		panic("redis连接测试异常...")
	} else {
		fmt.Println("redis连接正常...")
	}
}
