package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"sample/pkg/console"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"sample/pkg/str"
)

func init() {
	makeCommand.AddCommand(makeModelCommand)
}

const (
	modelDir = "app/models/%s"

	modelTplPath     = "template/model.tpl"
	modelHookTplPath = "template/model_hook.tpl"
	modelUtilTplPath = "template/model_util.tpl"
)

var makeModelCommand = &cobra.Command{
	Use:   "model",
	Short: "自动创建模型文件",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		model := generateModel(name)

		// 目录名和包名一致
		dir := fmt.Sprintf(modelDir, model.PackageName)

		// 父目录和子目录都会创建，第二个参数是目录权限，使用 0777
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			console.Exit("创建目录失败：" + err.Error())
			return
		}

		dataFunc := func() map[string]string {
			return model.ToMap()
		}
		makeFile(dir+"/"+model.PackageName+".go", modelTplPath, dataFunc)
		makeFile(dir+"/"+model.PackageName+"_util.go", modelUtilTplPath, dataFunc)
		makeFile(dir+"/"+model.PackageName+"_hook.go", modelHookTplPath, dataFunc)
	},
}

/*
用户输入：users、user、User、Users，得到的 Model 结构
{
	"TableName": "users",
	"StructName": "User",
	"StructNamePlural": "Users"
	"VariableName": "user",
	"VariableNamePlural": "users",
	"PackageName": "user"
}

用户输入：open_user、open_users、OpenUser、OpenUsers，得到的 Model 结构
{
	"TableName": "open_users",
	"StructName": "OpenUser",
	"StructNamePlural": "OpenUsers"
	"VariableName": "openUser",
	"VariableNamePlural": "openUsers",
	"PackageName": "open_user"
}
*/

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

func generateModel(in string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(in))
	model.StructNamePlural = str.Plural(model.StructName)
	model.VariableName = str.LowerCamel(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
	model.TableName = str.Snake(model.StructNamePlural)
	model.PackageName = str.Snake(model.StructName)
	return model
}

func (m Model) ToMap() map[string]string {
	var res map[string]string
	bs, _ := json.Marshal(m)
	_ = json.Unmarshal(bs, &res)
	return res
}
