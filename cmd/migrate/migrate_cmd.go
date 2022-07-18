package migrate

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mjiee/grf-gin/app/model"
	"github.com/mjiee/grf-gin/app/pkg/util"
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

	if err := insertData(data.db); err != nil {
		fmt.Println(err.Error())
		panic("初始化管理员失败...")
	}

	return nil
}

// 迁移schema
func dataMigrate(db *gorm.DB) error {
	fmt.Println("迁移user, manager表...")
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.User{},
		&model.Manager{},
	); err != nil {
		return err
	}

	return nil
}

// 插入初始数据
func insertData(db *gorm.DB) error {
	fmt.Println("插入初始化数据...")
	var count int64
	db.Model(&model.Manager{}).Select("id").Count(&count)

	if count == 0 {
		pwd, _ := util.BcryptPwd([]byte("admin123"))
		manager := &model.Manager{User: model.User{Name: "admin", Phone: "12345678900", Password: pwd}, Role: 3, Actived: true}
		return db.Create(manager).Error
	}
	return nil
}

// redis数据导入
func importRedis(redis *redis.Client) error {
	return nil
}
