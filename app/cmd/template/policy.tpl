package policies

import (
    "sample/app/models/{{.PackageName}}"
    "sample/pkg/auth"

    "github.com/gin-gonic/gin"
)

func CanModify{{.StructName}}(c *gin.Context, obj {{.PackageName}}.{{.StructName}}) bool {
    return auth.CurrentUid(c) == obj.UserID
}

// func CanView{{.StructName}}(c *gin.Context, obj {{.PackageName}}.{{.StructName}}) bool {}
// func CanCreate{{.StructName}}(c *gin.Context, obj {{.PackageName}}.{{.StructName}}) bool {}
// func CanUpdate{{.StructName}}(c *gin.Context, obj {{.PackageName}}.{{.StructName}}) bool {}
// func CanDelete{{.StructName}}(c *gin.Context, obj {{.PackageName}}.{{.StructName}}) bool {}
