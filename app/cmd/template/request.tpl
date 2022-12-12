package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type {{.StructName}}Request struct {
	// Phone string `json:"phone,omitempty" valid:"phone"`
}

func (r *{{.StructName}}) Validate() map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		//"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		//"phone": []string{
		//	"required:手机号为必填项，参数名称 phone",
		//	"digits:手机号长度必须为 11 位的数字",
		//},
	}

	opts := govalidator.Options{
		Data:     r,
		Rules:    rules,
		Messages: messages,
	}
	return validateStruct(opts)
}
