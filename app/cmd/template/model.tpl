
//Package {{.PackageName}} 模型
package {{.PackageName}}

import (

    "sample/app/models"
    "sample/pkg/database"
)

type {{.StructName}} struct {
    models.BaseModel

    // Put fields in here
    // TODO()

    models.CommonTimestampsField
}

func ({{.VariableName}} *{{.StructName}}) Create() {
    database.DB.Create({{.VariableName}})
}

func ({{.VariableName}} *{{.StructName}}) Save() (rowsAffected int64) {
    result := database.DB.Save({{.VariableName}})
    return result.RowsAffected
}
