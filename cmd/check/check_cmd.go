package check

import "github.com/spf13/cobra"

var (
	confFile string

	CheckCmd = &cobra.Command{
		Use:     "check",
		Short:   "scaffold-gin check",
		Long:    "检测项目配置和依赖服务",
		Example: "scaffold-gin check -f ./config/default.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
			// return nil
		},
	}
)

func init() {
	CheckCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "../conf/default.yaml", "提供项目配置文件")
}

// 初始化配置
func setup() {

}

// 运行检测
func run() error {
	return nil
}
