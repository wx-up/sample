package auth

import (
	"github.com/gin-gonic/gin"
	"sample/app/models/user"
)

// CurrentUid 当前用户ID
func CurrentUid(ctx *gin.Context) int64 {
	uid, ok := ctx.Get("current_user_id")
	if !ok {
		return 0
	}
	return uid.(int64)
}

// CurrentUser 当前用户
func CurrentUser(ctx *gin.Context) user.User {
	obj, ok := ctx.Get("current_user")
	if !ok {
		return user.User{}
	}
	return obj.(user.User)
}
