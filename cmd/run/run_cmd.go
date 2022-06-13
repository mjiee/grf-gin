package run

import "github.com/spf13/cobra"

var (
	confFile string

	RunCmd = &cobra.Command{
		Use:     "run",
		Short:   "scaffold-gin run",
		Long:    "run子命令用于运行项目",
		Example: "scaffold-gin run -c ./conf/default.yaml",
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
	RunCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "./conf/default.yaml", "指定项目配置文件")
}

// 初始化项目配置,数据库...
func setup() {

}

func run() error {
	return nil
}
