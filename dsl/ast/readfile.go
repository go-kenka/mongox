package ast

import (
	"fmt"
	"github.com/go-kenka/mongox/gen"
	"go/ast"
	"go/parser"
	"go/token"
)

func astReadFile(name, source string) *gen.Collection {
	fest := token.NewFileSet()
	f, err := parser.ParseFile(fest, name, source, parser.ParseComments)
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	t := &gen.Collection{}

	ast.Inspect(f, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.GenDecl:

			if n.Tok == token.VAR {

				if spec, ok := n.Specs[0].(*ast.ValueSpec); ok {
					if table, ok := spec.Values[0].(*ast.CallExpr); ok {
						//等于定义
						if spec.Names[0].Name == "_" {
							readTable(t, table)
						}
					}
				}
			}
		}
		return true
	})

	return t
}
