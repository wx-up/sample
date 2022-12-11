package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"sample/app/cmd"
	"sample/bootstrap"
	"sample/config"
	configPkg "sample/pkg/config"
	"sample/pkg/console"
)

func init() {
	config.Initialize()
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "sample",
		Short: "通用web框架",
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			env, _ := cmd.Flags().GetString("env")
			configPkg.InitConfig(env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			// bootstrap.SetupDB()

			// 初始化 Redis
			// bootstrap.SetupRedis()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd.Help()
			return nil
		},
	}
	rootCmd.AddCommand(cmd.ServerCommand)
	rootCmd.PersistentFlags().StringP("env", "e", "", "当前环境")

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("启动失败：%s", err.Error()))
	}
}
