package cmd

import (
	"github.com/spf13/cobra"
	"sample/bootstrap"
	"sample/pkg/config"
)

func init() {
	rootCmd.PersistentFlags().StringP("env", "e", "", "当前环境")
}

var rootCmd = &cobra.Command{
	Use:   "sample",
	Short: "通用web框架",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 配置初始化，依赖命令行 --env 参数
		env, _ := cmd.Flags().GetString("env")
		config.InitConfig(env)

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

func Execute() error {
	return rootCmd.Execute()
}
