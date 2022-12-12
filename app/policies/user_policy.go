package policies

import (
	"sample/app/models/user"
	"sample/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyUser(c *gin.Context, obj user.User) bool {
	return auth.CurrentUid(c) == obj.ID
}

// func CanViewUser(c *gin.Context, obj user.User) bool {}
// func CanCreateUser(c *gin.Context, obj user.User) bool {}
// func CanUpdateUser(c *gin.Context, obj user.User) bool {}
// func CanDeleteUser(c *gin.Context, obj user.User) bool {}
