// Package user 模型
package user

import (
	"sample/app/models"
	"sample/pkg/database"
)

type User struct {
	models.BaseModel

	Password string
	Username string

	models.CommonTimestampsField
}

func (user *User) Create() {
	database.DB.Create(user)
}

func (user *User) Save() (rowsAffected int64) {
	result := database.DB.Save(user)
	return result.RowsAffected
}
