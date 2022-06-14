package cmd

import (
	"fmt"
	"os"

	"github.com/mjiee/scaffold-gin/cmd/check"
	"github.com/mjiee/scaffold-gin/cmd/migrate"
	"github.com/mjiee/scaffold-gin/cmd/run"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scaffold-gin",
	Short: "scaffold-gin",
	Long:  "scaffold-gin",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("欢迎使用scaffold-gin, 使用-h查看帮助")
	},
}

func init() {
	rootCmd.AddCommand(check.CheckCmd)
	rootCmd.AddCommand(migrate.InitCmd)
	rootCmd.AddCommand(run.RunCmd)
}

// 执行命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
