package factories

import (
	"github.com/bxcodec/faker/v3"
	"sample/app/models/user"
)

func MakeUsers(times int) []user.User {
	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		obj := user.User{
			Username: faker.Username(),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, obj)
	}

	return objs
}
