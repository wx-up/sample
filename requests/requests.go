package requests

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type Validator interface {
	Validate() map[string][]string
}

type ValidateFunc func(ctx *gin.Context, data any) map[string][]string

func Validate(c *gin.Context, obj Validator) bool {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err.Error(),
		})
		return false
	}

	// 2. 表单验证
	errs := obj.Validate()

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不通过，具体请查看 errors",
			"errors":  errs,
		})
		return false
	}

	return true
}

func reflectCallValidateFunc(obj any, name string) map[string][]string {
	refV := reflect.ValueOf(obj)
	if !(refV.Kind() == reflect.Ptr && refV.Elem().Kind() == reflect.Struct) {
		return nil
	}
	refT := refV.Type()
	function, ok := refT.MethodByName(name)
	if !ok {
		return nil
	}
	res := function.Func.Call([]reflect.Value{refV})
	if len(res) <= 0 {
		return nil
	}
	return res[0].Interface().(map[string][]string)
}

func validate(data any, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}
