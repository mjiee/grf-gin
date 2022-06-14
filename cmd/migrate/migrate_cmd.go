package migrate

import "github.com/spf13/cobra"

var (
	confFile string

	InitCmd = &cobra.Command{
		Use:     "init",
		Short:   "scaffold-gin init",
		Long:    "初始化项目",
		Example: "scaffold-gin init -f ./config/default.yaml",
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
	InitCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "../conf/default.yaml", "提供项目配置文件")
}

// 初始化配置
func setup() {

}

// 运行检测
func run() error {
	return nil
}
