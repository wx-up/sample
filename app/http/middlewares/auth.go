// Package middlewares Gin 中间件
package middlewares

import (
	"sample/app/models/user"
	"sample/pkg/jwt"
	"sample/pkg/response"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.New().ParseToken(c)
		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, "token 未传递或者格式错误")
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := user.Get(claims.Uid)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.ID)
		c.Set("current_user_name", userModel.Username)
		c.Set("current_user", userModel)
		c.Next()
	}
}
