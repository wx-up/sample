package user

import (
	"sample/pkg/database"
)

func Get(id int64) (user User) {
	database.DB.Where("id = ?", id).First(&user)
	return
}

func GetBy(field, value string) (user User) {
	database.DB.Where("? = ?", field, value).First(&user)
	return
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(User{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}
