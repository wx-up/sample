package cmd

import (
	"github.com/spf13/cobra"
	"sample/pkg/console"
	"sample/pkg/helpers"
)

func init() {
	Register(keyCommand)
}

var keyCommand = &cobra.Command{
	Use:   "key",
	Short: "创建 App Key",
	Run: func(cmd *cobra.Command, args []string) {
		console.Success("App Key：")
		console.Success(helpers.RandomString(32))
		console.Warning("请将 App Key 复制到 .env 文件中")
	},
	Args: cobra.NoArgs,
}
