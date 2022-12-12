package routes

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

type Visitor struct{}

func (v *Visitor) Visit(node ast.Node) (w ast.Visitor) {
	// TODO implement me
	switch v := node.(type) {
	case *ast.GenDecl:
		importControllerPath := "sample/app/http/controllers/api/v1"
		exist := false
		for _, spec := range v.Specs {
			importSpec, ok := spec.(*ast.ImportSpec)
			if !ok {
				continue
			}
			if strings.Contains(importSpec.Path.Value, importControllerPath) {
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
					Value: "\"sample/app/http/controllers/api/v1\"",
				},
				Comment: nil,
				EndPos:  0,
			})
		}
	case *ast.FuncDecl:
		v.Body.List = append(v.Body.List, &ast.AssignStmt{
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: "uc",
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
							Value: "controllers.UserController",
						},
					},
				},
			},
		})
		v.Body.List = append(v.Body.List, &ast.AssignStmt{
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: "ucGroup",
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
							Value: `"users"`,
						},
					},
				},
			},
		})
		block := &ast.BlockStmt{}
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "ucGroup.GET",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: `""`,
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: "uc.Index",
					},
				},
			},
		})
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "ucGroup.POST",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "\"\"",
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: "uc.Store",
					},
				},
			},
		})
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "ucGroup.PUT",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "\":id\"",
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: "uc.Update",
					},
				},
			},
		})
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "ucGroup.DELETE",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "\":id\"",
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: "uc.Delete",
					},
				},
			},
		})
		block.List = append(block.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "ucGroup.GET",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "\":id\"",
					},
					&ast.BasicLit{
						Kind:  token.TYPE,
						Value: "uc.Show",
					},
				},
			},
		})

		v.Body.List = append(v.Body.List, block)
	}
	return v
}

func TestRegisterAPIRouters(t *testing.T) {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, "api.go", nil, parser.ParseComments)
	fmt.Println(err)
	ast.Walk(&Visitor{}, f)

	var output []byte
	buffer := bytes.NewBuffer(output)
	fmt.Println(format.Node(buffer, fSet, f))
	fmt.Println(buffer.String())
}
