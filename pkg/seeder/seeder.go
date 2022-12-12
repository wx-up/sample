package seeder

import (
	"errors"
	"sync"

	"sample/pkg/console"

	"sample/pkg/database"

	"gorm.io/gorm"
)

type Handle func(db *gorm.DB) error

type Seeder struct {
	Name   string
	Handle Handle
}

var (
	seeders = sync.Map{}

	orderedSeederNames []string
)

// SetRunOrder 支持一些必须按顺序执行的 seeder，例如 topic 创建的
// 时必须依赖于 user, 所以 TopicSeeder 应该在 UserSeeder 后执
func SetRunOrder(names []string) {
	orderedSeederNames = names
}

func Add(name string, handle Handle) {
	seeders.Store(name, Seeder{
		Name:   name,
		Handle: handle,
	})
}

var ErrSeederNotFound = errors.New("找不到对应的 seeder ")

func Get(name string) (Seeder, error) {
	v, ok := seeders.Load(name)
	if !ok {
		return Seeder{}, ErrSeederNotFound
	}
	return v.(Seeder), nil
}

func RunAll() {
	// 先执行顺序的 seeder
	runs := make(map[string]struct{})
	for _, name := range orderedSeederNames {
		err := RunSeeder(name)
		if err != nil {
			continue
		}
		runs[name] = struct{}{}
	}

	// 执行剩下的 seeder
	seeders.Range(func(key, value any) bool {
		keyString := key.(string)
		if _, ok := runs[keyString]; ok {
			return true
		}

		_ = RunSeeder(keyString)
		return true
	})
}

func RunSeeder(name string) error {
	seeder, err := Get(name)
	if err != nil {
		console.Warning("获取 " + name + "失败：" + err.Error())
		return err
	}
	err = seeder.Handle(database.DB)
	if err != nil {
		console.Warning(name + " seeder 执行失败：" + err.Error())
		return err
	}
	return nil
}
