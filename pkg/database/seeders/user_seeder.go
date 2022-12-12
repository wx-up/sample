package seeders

//import (
//	"fmt"
//
//	"sample/pkg/seeder"
//
//	"gorm.io/gorm"
//	"sample/pkg/console"
//	"sample/pkg/database/factories"
//)
//
//func init() {
//	// 添加 Seeder
//	seeder.Add("SeedUsersTable", func(db *gorm.DB) error {
//		// 创建 10 个用户对象
//		users := factories.MakeUsers(10)
//
//		// 批量创建用户（注意批量创建不会调用模型钩子）
//		result := db.Table("users").Create(&users)
//
//		// 记录错误
//		if err := result.Error; err != nil {
//			return err
//		}
//
//		// 打印运行情况
//		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
//
//		return nil
//	})
//}
