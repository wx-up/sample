package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"sample/app/http/requests"
	"sample/app/models/{{.PackageName}}"
	"sample/app/policies"
	"sample/pkg/response"
)

type {{.StructNamePlural}}Controller struct {
	BaseController
}

// Index 列表
func (c *{{.StructNamePlural}}Controller) Index(ctx *gin.Context) {
	objs := {{.PackageName}}.All()
	response.JSON(ctx, objs)
}

// Show 详情
func (c *{{.StructNamePlural}}Controller) Show(ctx *gin.Context) {
	obj := {{.PackageName}}.Get(cast.ToInt64(ctx.Param("id")))
	response.JSON(ctx, obj)
}

// Store 创建
func (c *{{.StructNamePlural}}Controller) Store(ctx *gin.Context) {
	var in requests.{{.StructName}}Request
	requests.Validate(ctx, &in)

    obj := &{{.PackageName}}.{{.StructName}}{}
	if err := in.Copy(&obj); err != nil {
    	response.BadRequest(ctx)
    	return
    }

	obj.Create()

	if obj.ID <= 0 {
		response.Abort500(ctx, "创建失败，请稍后再试")
		return
	}
	response.JSON(ctx, obj)
}

// Update 更新
func (c *{{.StructNamePlural}}Controller) Update(ctx *gin.Context) {
	obj := {{.PackageName}}.Get(cast.ToInt64(ctx.Param("id")))
	if obj.ID <= 0 {
		response.Abort404(ctx)
		return
	}

	if policies.CanModify{{.StructName}}(ctx, obj) {
		response.Abort403(ctx)
		return
	}

	var in requests.{{.StructName}}Request
	requests.Validate(ctx, &in)

	if err := in.Copy(&obj); err != nil {
		response.BadRequest(ctx)
		return
	}

	rowsAffected := obj.Save()
	if rowsAffected <= 0 {
		response.Abort500(ctx, "更新失败，请稍后再试")
		return
	}
	response.JSON(ctx, obj)
}

func (c *{{.StructNamePlural}}Controller) Delete(ctx *gin.Context) {
	obj := {{.PackageName}}.Get(cast.ToInt64(ctx.Param("id")))
	if obj.ID <= 0 {
		response.Abort404(ctx)
		return
	}

	if policies.CanModify{{.StructName}}(ctx, obj) {
		response.Abort403(ctx)
		return
	}

	rowsAffected := obj.Delete()
	if rowsAffected <= 0 {
		response.Abort500(ctx, "删除失败，请稍后再试")
		return
	}

	response.Success(ctx)
}
