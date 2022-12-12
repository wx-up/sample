// Package response 响应处理工具
package response

import (
	"net/http"

	"sample/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Data struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSON 响应 200 和 JSON 数据
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Data{
		Code:    0,
		Message: "获取成功",
		Data:    data,
	})
}

// Success 响应 200 和预设『操作成功！』的 JSON 数据
// 执行某个『没有具体返回数据』的『变更』操作成功后调用，例如删除、修改密码、修改手机号
func Success(c *gin.Context) {
	c.JSON(http.StatusOK, Data{
		Code:    0,
		Message: "操作成功",
	})
}

// Created 响应 201 和 JSON 数据
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Data{
		Code:    0,
		Message: "创建成功",
		Data:    data,
	})
}

// Abort404 响应 404，未传参 msg 时使用默认消息
func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, Data{
		Code:    10000,
		Message: defaultMessage("资源不存在", msg...),
	})
}

/*
	401 Unauthorized响应应该用来表示缺失或错误的认证
	403 Forbidden响应应该在这之后用，当用户被认证后，但用户没有被授权在特定资源上执行操作
*/

// Abort403 响应 403，未传参 msg 时使用默认消息
func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, Data{
		Code:    10000,
		Message: defaultMessage("权限不足", msg...),
	})
}

func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
	})
}

// Abort500 响应 500，未传参 msg 时使用默认消息
func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Data{
		Code:    10000,
		Message: defaultMessage("服务器内部错误", msg...),
	})
}

// BadRequest 响应 400，传参 err 对象，未传参 msg 时使用默认消息
// 在解析用户请求，请求的格式或者方法不符合预期时调用
func BadRequest(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, Data{
		Code:    10000,
		Message: defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
	})
}

// Error 响应 404 或 422，未传参 msg 时使用默认消息
// 处理请求时出现错误 err，会附带返回 error 信息，如登录错误、找不到 ID 对应的 Model
func Error(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	// error 类型为『数据库未找到内容』
	if err == gorm.ErrRecordNotFound {
		Abort404(c)
		return
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("请求处理失败，请查看 error 的值", msg...),
		"error":   err.Error(),
	})
}

// ValidationError 处理表单验证不通过的错误，返回的 JSON 示例：
//
//	{
//	    "errors": {
//	        "phone": [
//	            "手机号为必填项，参数名称 phone",
//	            "手机号长度必须为 11 位的数字"
//	        ]
//	    },
//	    "message": "请求验证不通过，具体请查看 errors"
//	}
func ValidationError(c *gin.Context, errors map[string][]string) {
}

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	message = defaultMsg

	if len(msg) > 0 {
		message = msg[0]
	}
	return
}
