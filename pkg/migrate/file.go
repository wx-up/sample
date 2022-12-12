package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

type MigrationFunc func(gorm.Migrator, *sql.DB)

// MigrationFile 单个迁移文件
type MigrationFile struct {
	name string
	up   MigrationFunc
	down MigrationFunc
}

var files []MigrationFile

// AddMigrationFile 添加迁移文件
func AddMigrationFile(name string, up MigrationFunc, down MigrationFunc) {
	files = append(files, MigrationFile{
		name: name,
		up:   up,
		down: down,
	})
}
