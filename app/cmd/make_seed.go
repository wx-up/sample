package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"sample/pkg/console"
)

func init() {
	makeCommand.AddCommand(makeSeedCommand)
}

const (
	seedDir     = "pkg/database/seeders"
	seedTplPath = "template/seed.tpl"
)

var makeSeedCommand = &cobra.Command{
	Use:   "seed",
	Short: "创建 数据填充 模板",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.MkdirAll(seedDir, os.ModePerm); err != nil {
			console.Exit("创建目录失败：" + err.Error())
			return
		}

		model := generateModel(args[0])
		makeFile(seedDir+"/"+model.PackageName+"_seeder.go", seedTplPath, func() map[string]string {
			return model.ToMap()
		})
	},
}
