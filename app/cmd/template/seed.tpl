package seeders

import (
	"fmt"

	"sample/pkg/seeder"

	"gorm.io/gorm"
	"sample/pkg/console"
	"sample/pkg/database/factories"
)

func init() {
	// 添加 Seeder
	seeder.Add("Seed{{.StructNamePlural}}Table", func(db *gorm.DB) error {
		// 创建 10 个对象
		objs := factories.Make{{.StructNamePlural}}(10)

		// 批量创建用户（注意批量创建不会调用模型钩子）
		result := db.Table("{{.TableName}}").Create(&objs)

		// 记录错误
		if err := result.Error; err != nil {
			return err
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))

		return nil
	})
}
