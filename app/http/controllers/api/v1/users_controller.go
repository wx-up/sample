package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"sample/app/http/requests"
	"sample/app/models/user"
	"sample/app/policies"
	"sample/pkg/response"
)

type UsersController struct {
	BaseController
}

// Index 列表
func (c *UsersController) Index(ctx *gin.Context) {
	objs := user.All()
	response.JSON(ctx, objs)
}

// Show 详情
func (c *UsersController) Show(ctx *gin.Context) {
	obj := user.Get(cast.ToInt64(ctx.Param("id")))
	response.JSON(ctx, obj)
}

// Store 创建
func (c *UsersController) Store(ctx *gin.Context) {
	var in requests.UserRequest
	requests.Validate(ctx, &in)

	obj := user.User{
		Password: "",
		Username: "",
	}

	obj.Create()

	if obj.ID <= 0 {
		response.Abort500(ctx, "创建失败，请稍后再试")
		return
	}
	response.JSON(ctx, obj)
}

// Update 更新
func (c *UsersController) Update(ctx *gin.Context) {
	obj := user.Get(cast.ToInt64(ctx.Param("id")))
	if obj.ID <= 0 {
		response.Abort404(ctx)
		return
	}

	if policies.CanModifyUser(ctx, obj) {
		response.Abort403(ctx)
		return
	}

	var in requests.UserRequest
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

func (c *UsersController) Delete(ctx *gin.Context) {
	obj := user.Get(cast.ToInt64(ctx.Param("id")))
	if obj.ID <= 0 {
		response.Abort404(ctx)
		return
	}

	if policies.CanModifyUser(ctx, obj) {
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
