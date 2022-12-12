package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"sample/pkg/console"
)

func init() {
	makeCommand.AddCommand(makePolicyCommand)
}

const (
	policyDir     = "app/policies"
	policyTplPath = "template/policy.tpl"
)

var makePolicyCommand = &cobra.Command{
	Use:   "policy",
	Short: "创建 授权策略 模板",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.MkdirAll(policyDir, os.ModePerm); err != nil {
			console.Exit("创建目录失败：" + err.Error())
			return
		}

		model := generateModel(args[0])

		makeFile(policyDir+"/"+model.PackageName+"_policy.go", policyTplPath, func() map[string]string {
			return model.ToMap()
		})
	},
}
