package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"sample/pkg/console"
)

func init() {
	makeCommand.AddCommand(requestCommand)
}

const (
	requestFileDir = "app/http/requests"
	requestTplPath = "template/request.tpl"
)

var requestCommand = &cobra.Command{
	Use:   "request",
	Short: "自动创建请求参数文件",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.MkdirAll(requestFileDir, os.ModePerm); err != nil {
			console.Exit("创建目录失败：" + err.Error())
			return
		}

		name := args[0]
		model := generateModel(name)
		filePath := requestFileDir + "/" + model.PackageName + "_request.go"
		makeFile(filePath, requestTplPath, func() map[string]string {
			return model.ToMap()
		})
	},
}
