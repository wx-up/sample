package {{.PackageName}}

import (
    "sample/pkg/database"
)

func Get(id int64) ({{.VariableName}} {{.StructName}}) {
    database.DB.Where("id = ?", id).First(&{{.VariableName}})
    return
}

func GetBy(field, value string) ({{.VariableName}} {{.StructName}}) {
    database.DB.Where("? = ?", field, value).First(&{{.VariableName}})
    return
}

func All() ({{.VariableNamePlural}} []{{.StructName}}) {
    database.DB.Find(&{{.VariableNamePlural}})
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model({{.StructName}}{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}