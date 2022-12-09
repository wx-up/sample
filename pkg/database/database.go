// Package database 数据库操作
package database

import (
	"database/sql"
	"sync"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// DB 对象
var DB *gorm.DB
var SqlDB *sql.DB

var once sync.Once

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, log gormLogger.Interface) {
	once.Do(func() {
		// 使用 gorm.Open 连接数据库
		var err error
		DB, err = gorm.Open(dbConfig, &gorm.Config{
			Logger: log,
		})
		// 处理错误
		if err != nil {
			panic(err)
		}

		// 获取底层的 sqlDB
		SqlDB, err = DB.DB()
		if err != nil {
			panic(err)
		}
	})
}
