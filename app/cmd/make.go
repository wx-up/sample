package cmd

import (
	"embed"
	"html/template"
	"os"

	"sample/pkg/console"

	"sample/pkg/helpers"

	"github.com/spf13/cobra"
)

//go:embed template
var templateFS embed.FS

func init() {
	Register(makeCommand)
}

var makeCommand = &cobra.Command{
	Use:   "make",
	Short: "自动创建一些文件和代码，比如模型、控制器等等",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) <= 0 {
			_ = cmd.Help()
			return nil
		}
		return nil
	},
	Args: cobra.NoArgs,
}

func makeFile(path string, tplPath string, dataFunc func() map[string]string) {
	if helpers.FileExists(path) {
		console.Exit(path + "已经存在")
		return
	}

	tpl, err := template.ParseFS(templateFS, tplPath)
	if err != nil {
		console.Exit("解析模板错误：" + err.Error())
		return
	}
	fd, err := os.Create(path)
	if err != nil {
		console.Exit("创建文件错误：" + err.Error())
		return
	}
	err = tpl.Execute(fd, dataFunc())
	if err != nil {
		console.Exit("写入模板错误：" + err.Error())
		return
	}
	console.Success("创建文件：" + path + " 成功")
	return
}
