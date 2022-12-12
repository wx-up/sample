package cmd

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"sample/pkg/config"

	"sample/pkg/helpers"

	"github.com/spf13/cobra"
	"sample/pkg/console"
)

func init() {
	makeCommand.AddCommand(makeControllerCommand)
}

const (
	controllerDir     = "app/http/controllers/api/%s"
	controllerTplPath = "template/controller.tpl"
)

var makeControllerCommand = &cobra.Command{
	Use:   "controller",
	Short: "创建控制器模板，参数格式：版本/控制器单数名称（ v1/user ）",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		argSplit := strings.Split(arg, "/")
		if len(argSplit) != 2 {
			console.Error("参数格式错误，请使用 版本/控制器单数名称 的格式")
			return
		}

		version, name := argSplit[0], argSplit[1]
		model := generateModel(name)

		controllerDir := fmt.Sprintf(controllerDir, version)
		if err := os.MkdirAll(controllerDir, os.ModePerm); err != nil {
			console.Exit("创建目录失败：" + err.Error())
			return
		}

		makeFile(controllerDir+"/"+model.TableName+"_controller.go", controllerTplPath, func() map[string]string {
			return model.ToMap()
		})

		generateRoute(version, model)
	},
}

func generateRoute(version string, model Model) {
	fSet, filename := token.NewFileSet(), "routes/api.go"
	f, err := parser.ParseFile(fSet, filename, nil, parser.ParseComments)
	if err != nil {
		console.Warning("自动生成路由失败：" + err.Error())
		return
	}
	ast.Walk(&RouterVisitor{
		importPath: fmt.Sprintf(config.Get("app.name")+"/"+controllerDir, version),
		model:      model,
	}, f)

	var buffer bytes.Buffer
	err = format.Node(&buffer, fSet, f)
	if err != nil {
		console.Warning("自动生成路由失败：" + err.Error())
		return
	}
	err = helpers.PutFile(buffer.Bytes(), filename)
	if err != nil {
		console.Warning("自动生成路由失败：" + err.Error())
		return
	}

	console.Success("自动生成路由成功：" + filename)
}

type RouterVisitor struct {
	importPath string
	model      Model
}

func (obj *RouterVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch v := node.(type) {
	case *ast.GenDecl:
		obj.paddingImportSpec(v)
		return obj
	case *ast.FuncDecl:
		name := string(obj.model.TableName[0]) + "c"
		nameGroup := name + "Group"
		v.Body.List = append(v.Body.List, &ast.AssignStmt{
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: name,
				},
			},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.Ident{
						Name: "new",
					},
					Args: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.TYPE,
							Value: fmt.Sprintf("controllers.%sController", obj.model.StructNamePlural),
						},
					},
				},
			},
		}, &ast.AssignStmt{
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: nameGroup,
				},
			},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.Ident{
						Name: "v1.Group",
					},
					Args: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf("\"%s\"", obj.model.TableName),
						},
					},
				},
			},
		})
		v.Body.List = append(v.Body.List)
		block := &ast.BlockStmt{}
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: fmt.Sprintf("%s.GET", nameGroup),
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "\"\"",
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: fmt.Sprintf("%s.Index", name),
					},
				},
			},
		})
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: fmt.Sprintf("%s.POST", nameGroup),
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "\"\"",
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: fmt.Sprintf("%s.Store", name),
					},
				},
			},
		})
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: fmt.Sprintf("%s.PUT", nameGroup),
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "\":id\"",
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: fmt.Sprintf("%s.Update", name),
					},
				},
			},
		})
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: fmt.Sprintf("%s.DELETE", nameGroup),
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "\":id\"",
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: fmt.Sprintf("%s.Delete", name),
					},
				},
			},
		})
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: fmt.Sprintf("%s.GET", nameGroup),
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "\":id\"",
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: fmt.Sprintf("%s.Show", name),
					},
				},
			},
		})
		v.Body.List = append(v.Body.List, block)
	}
	return obj
}

func (obj *RouterVisitor) paddingImportSpec(v *ast.GenDecl) {
	exist := false
	for _, spec := range v.Specs {
		importSpec, ok := spec.(*ast.ImportSpec)
		if !ok {
			continue
		}
		if strings.Contains(importSpec.Path.Value, obj.importPath) {
			exist = true
		}
	}
	if !exist {
		v.Specs = append(v.Specs, &ast.ImportSpec{
			Name: &ast.Ident{
				Name: "controllers",
			},
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s\"", obj.importPath),
			},
			Comment: nil,
			EndPos:  0,
		})
	}
	return
}
