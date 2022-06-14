package migrate

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mjiee/scaffold-gin/app/model"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	confFile string

	InitCmd = &cobra.Command{
		Use:     "init",
		Short:   "scaffold-gin init",
		Long:    "初始化项目",
		Example: "scaffold-gin init -f ./config/default.yaml",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
			// return nil
		},
	}
)

func init() {
	InitCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "../conf/default.yaml", "提供项目配置文件")
}

// 运行检测
func run() error {
	data, cleanup, err := initData(confFile)
	if err != nil {
		return err
	}

	defer cleanup()

	if err := dataMigrate(data.db); err != nil {
		fmt.Println(err.Error())
		panic("初始化数据库表失败...")
	}

	return nil
}

// 迁移schema
func dataMigrate(db *gorm.DB) error {
	fmt.Println("迁移user表...")
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}); err != nil {
		return err
	}

	return nil
}

// redis数据导入
func importRedis(redis *redis.Client) error {
	return nil
}
