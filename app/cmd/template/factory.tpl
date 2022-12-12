package factories

import (
	"github.com/bxcodec/faker/v3"
	"sample/app/models/{{.PackageName}}"
)

func Make{{.StructNamePlural}}(times int) []{{.PackageName}}.{{.StructName}} {
	var objs []{{.PackageName}}.{{.StructName}}

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		obj := {{.PackageName}}.{{.StructName}}{
			// TODO()
		}
		objs = append(objs, obj)
	}

	return objs
}
