package bootstrap

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"sample/routes"
)

func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleware(router)

	// 注册 API 路由
	routes.RegisterAPIRouters(router)

	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		contentType := ctx.ContentType()
		if strings.Contains(contentType, "application/json") {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确",
			})
		} else {
			ctx.String(http.StatusNotFound, "页面返回 404")
		}
	})
}
