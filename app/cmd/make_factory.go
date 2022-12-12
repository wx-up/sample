package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"sample/pkg/console"
)

func init() {
	makeCommand.AddCommand(makeFactoryCommand)
}

const (
	factoryDir     = "pkg/database/factories"
	factoryTplPath = "template/factory.tpl"
)

var makeFactoryCommand = &cobra.Command{
	Use:   "factory",
	Short: "创建数据工程",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.MkdirAll(factoryDir, os.ModePerm); err != nil {
			console.Exit("创建目录失败：" + err.Error())
			return
		}

		model := generateModel(args[0])
		makeFile(factoryDir+"/"+model.PackageName+"_factory.go", factoryTplPath, func() map[string]string {
			return model.ToMap()
		})
	},
}
