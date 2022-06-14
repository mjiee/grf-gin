package check

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/mjiee/scaffold-gin/app/pkg/util"
	"github.com/spf13/cobra"
)

var (
	confFile string

	CheckCmd = &cobra.Command{
		Use:     "check",
		Short:   "scaffold-gin check",
		Long:    "检测项目配置和依赖服务",
		Example: "scaffold-gin check -f ./config/default.yaml",
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
	app, cleanup, err := NewChecker(confFile)
	if err != nil {
		return err
	}

	defer cleanup()

	checkConf(app)

	return nil
}

// check config
func checkConf(app App) {
	validate := validator.New()
	validate.RegisterValidation("addr", util.CheckAddr)

	if err := validate.Struct(app.conf); err != nil {
		fmt.Println(err.Error())
		panic("配置文件异常...")
	} else {
		fmt.Println("配置文件检测正常...")
	}
}
