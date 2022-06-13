package cmd

import (
	"fmt"
	"os"

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

// 执行命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
