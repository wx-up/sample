package cmd

import (
	"fmt"
	"html/template"
	"os"
	"testing"

	"sample/pkg/str"
)

func Test(t *testing.T) {
	tpl, err := template.ParseFS(templateFS, "template/model.tpl")
	fmt.Println(err)
	fmt.Println(tpl)
	obj := Model{
		TableName:          "open_users",
		StructName:         "OpenUser",
		StructNamePlural:   "OpenUsers",
		VariableName:       "openUser",
		VariableNamePlural: "openUsers",
		PackageName:        "open_user",
	}
	tpl.Execute(os.Stdout, obj)
}

func TestStr(t *testing.T) {
	fmt.Println(str.Camel("user_name"))
	fmt.Println(str.LowerCamel("user_name"))
	fmt.Println(str.Snake("UserName"))
	fmt.Println(str.Snake("userName"))
}
