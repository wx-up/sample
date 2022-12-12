package requests

import (
	"errors"
	"reflect"

	"sample/pkg/helpers"

	"sample/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type Validator interface {
	Validate() map[string][]string
}

type ValidateFunc func(ctx *gin.Context, data any) map[string][]string

func validateStruct(opts govalidator.Options) map[string][]string {
	opts.TagIdentifier = "valid"
	return govalidator.New(opts).ValidateStruct()
}

func Validate(c *gin.Context, obj Validator) {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c)
		return
	}

	// 2. 表单验证
	errs := obj.Validate()

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		response.BadRequest(c, helpers.ParseRequestErrs(errs))
		return
	}
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

func Copy(src any, dst any) error {
	srcV := reflect.ValueOf(src)
	dstV := reflect.ValueOf(dst)
	if !(srcV.Kind() == reflect.Ptr && srcV.Elem().Kind() == reflect.Struct) {
		return errors.New("src 非指向指针的结构体")
	}
	if !(dstV.Kind() == reflect.Ptr && dstV.Elem().Kind() == reflect.Struct) {
		return errors.New("dst 非指向指针的结构体")
	}
	srcV = srcV.Elem()
	dstV = dstV.Elem()

	dstT := dstV.Type()
	srvT := srcV.Type()
	num := srvT.NumField()
	for i := 0; i < num; i++ {
		name := srvT.Field(i).Name
		value := srcV.Field(i)
		_, ok := dstT.FieldByName(name)
		if !ok {
			continue
		}
		found := dstV.FieldByName(name)

		// 相同类型才可以设置值
		if found.CanSet() && found.Kind() == value.Kind() {
			found.Set(value)
		}
	}
	return nil
}

func validate(data any, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}
