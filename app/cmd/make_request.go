package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	makeCommand.AddCommand(requestCommand)
}

const (
	requestFilePath = "app/http/requests/%s_request.go"
	requestTplPath  = "template/request.tpl"
)

var requestCommand = &cobra.Command{
	Use:   "request",
	Short: "自动创建请求参数文件",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		model := generateModel(name)
		filePath := fmt.Sprintf(requestFilePath, model.PackageName)
		makeFile(filePath, requestTplPath, func() map[string]string {
			return model.ToMap()
		})
	},
}
