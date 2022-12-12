package cmd

import (
	"github.com/spf13/cobra"
	"sample/pkg/console"
	"sample/pkg/database/seeders"
	"sample/pkg/seeder"
)

func init() {
	rootCmd.AddCommand(seederCommand)
}

var seederCommand = &cobra.Command{
	Use:   "seed",
	Short: "执行数据填充",
	Args:  cobra.MaximumNArgs(1), // 最多允许一个参数
	Run: func(cmd *cobra.Command, args []string) {
		seeders.Initialize()

		if len(args) <= 0 {
			seeder.RunAll()
			console.Success("所有的 seeders 执行完毕")
			return
		}

		seederName := args[0]

		_ = seeder.RunSeeder(seederName)
		console.Success(seederName + " seeder 执行完毕")
	},
}
