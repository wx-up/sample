package migrate

import (
	"gorm.io/gorm"
	"sample/pkg/database"
)

type Migrator struct {
	folder   string
	db       *gorm.DB
	migrator gorm.Migrator
}

type Option func(*Migrator)

func WithFolder(folder string) Option {
	return func(migrator *Migrator) {
		migrator.folder = folder
	}
}

func WithGormDB(db *gorm.DB) Option {
	return func(migrator *Migrator) {
		migrator.db = db
	}
}

func New(opts ...Option) *Migrator {
	res := &Migrator{
		folder: "database/migrations/",
		db:     database.DB,
	}
	for _, opt := range opts {
		opt(res)
	}
	res.migrator = res.db.Migrator()

	// 创建迁移记录表： migrations
	res.createMigrationsTable()

	return res
}

type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// 创建 migrations 表
func (m *Migrator) createMigrationsTable() {
	migration := Migration{}

	// 不存在才创建
	if !m.migrator.HasTable(&migration) {
		_ = m.migrator.CreateTable(&migration)
	}
}
