package seeders

import "sample/pkg/seeder"

func Initialize() {
	// 触发加载本目录下其他文件中的 init 方法

	// 指定优先于同目录下的其他文件运行
	seeder.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
