package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterAPIRouters(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	_ = v1
}
