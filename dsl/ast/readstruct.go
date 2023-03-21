package ast

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-kenka/mongox/gen"
	"go/ast"
	"go/format"
	"go/token"
	"log"
	"strconv"
)

func readTable(table *gen.Collection, call *ast.CallExpr) {
	fun, ok := call.Fun.(*ast.Ident)
	//等于定义
	if ok && fun.Name == "Collection" {
		for _, arg := range call.Args {
			switch a := arg.(type) {
			case *ast.BasicLit:
				table.Name = getStringValue(a)
			case *ast.CallExpr:
				readTableFn(table, a)
			}
		}
	}
}

func readTableFn(table *gen.Collection, call *ast.CallExpr) {
	if fun, ok := call.Fun.(*ast.Ident); ok {
		switch fun.Name {
		case "Database":
			readTableDatabase(table, call.Args[0])
		case "Desc":
			readTableDesc(table, call.Args[0])
		case "Fields":
			readTableFields(table, call.Args)
		case "Indexes":
			readTableIndexes(table, call.Args)
		}
	}
}

func readTableDatabase(table *gen.Collection, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		table.Database = getStringValue(d)
	}
}

func readTableDesc(table *gen.Collection, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		table.Desc = getStringValue(d)
	}
}

func readTableFields(table *gen.Collection, args []ast.Expr) {

	for _, arg := range args {

		fs := &gen.Field{}

		if call, ok := arg.(*ast.CallExpr); ok {
			//等于定义
			if fun, ok := call.Fun.(*ast.Ident); ok && fun.Name == "Field" {
				for _, fg := range call.Args {
					switch a := fg.(type) {
					case *ast.BasicLit:
						fs.Name = getStringValue(a)
					case *ast.CallExpr:
						readFieldFn(fs, a)
					}
				}
			}
		}

		table.Fields = append(table.Fields, fs)
	}

}

func readFieldFn(fs *gen.Field, call *ast.CallExpr) {
	if fun, ok := call.Fun.(*ast.Ident); ok {
		switch fun.Name {
		case "Type":
			readFieldType(fs, call.Args[0])
		case "ObjectAttr":
			readFieldObjectAttr(fs, call.Args)
		case "ArrayType":
			readFieldArrayType(fs, call.Args[0])
		case "Comment":
			readFieldComment(fs, call.Args[0])
		}
	}

}

func readFieldType(fs *gen.Field, arg ast.Expr) {
	if d, ok := arg.(*ast.Ident); ok {
		fs.Type = gen.TypeNameMap[d.Name]
	}
}

func readFieldArrayType(fs *gen.Field, arg ast.Expr) {
	if d, ok := arg.(*ast.Ident); ok {
		fs.ArrayType = gen.TypeNameMap[d.Name]
	}
}

func readFieldObjectAttr(fs *gen.Field, args []ast.Expr) {
	for _, arg := range args {

		f := &gen.Field{}

		if call, ok := arg.(*ast.CallExpr); ok {
			//等于定义
			if fun, ok := call.Fun.(*ast.Ident); ok && fun.Name == "Field" {
				for _, fg := range call.Args {
					switch a := fg.(type) {
					case *ast.BasicLit:
						f.Name = getStringValue(a)
					case *ast.CallExpr:
						readFieldFn(f, a)
					}
				}
			}
		}

		fs.ObjectAttr = append(fs.ObjectAttr, f)
	}
}

func readFieldComment(fs *gen.Field, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		fs.Comment = getStringValue(d)
	}
}

func readTableIndexes(table *gen.Collection, args []ast.Expr) {

	for _, arg := range args {

		is := &gen.Index{}

		if call, ok := arg.(*ast.CallExpr); ok {
			//等于定义
			if fun, ok := call.Fun.(*ast.Ident); ok && fun.Name == "Index" {
				for _, fg := range call.Args {
					switch a := fg.(type) {
					case *ast.BasicLit:
						is.Name = getStringValue(a)
					case *ast.CallExpr:
						readIndexFn(is, a)
					}
				}
			}
		}

		table.Indexes = append(table.Indexes, is)
	}

}

func readIndexFn(edge *gen.Index, call *ast.CallExpr) {
	if fun, ok := call.Fun.(*ast.Ident); ok {
		switch fun.Name {
		case "Sparse":
			readIndexSparse(edge, call.Args[0])
		case "Background":
			readIndexBackground(edge, call.Args[0])
		case "Expire":
			readIndexExpire(edge, call.Args[0])
		case "Unique":
			readIndexUnique(edge, call.Args[0])
		case "SortKey":
			readIndexSortKey(edge, call.Args)
		}
	}

}

func readIndexExpire(is *gen.Index, arg ast.Expr) {
	set := token.NewFileSet()
	var output []byte
	buffer := bytes.NewBuffer(output)
	err := format.Node(buffer, set, arg)
	if err != nil {
		log.Fatal(err)
	}
	is.Expire = buffer.String()
}
func readIndexBackground(is *gen.Index, arg ast.Expr) {
	if d, ok := arg.(*ast.Ident); ok {
		is.Background = getBoolValue(d.Name)
	}
}
func readIndexUnique(is *gen.Index, arg ast.Expr) {
	if d, ok := arg.(*ast.Ident); ok {
		is.Unique = getBoolValue(d.Name)
	}
}
func readIndexSparse(is *gen.Index, arg ast.Expr) {
	if d, ok := arg.(*ast.Ident); ok {
		is.Sparse = getBoolValue(d.Name)
	}
}

func readIndexSortKey(is *gen.Index, args []ast.Expr) {

	k := &gen.Key{}

	key := args[0].(*ast.BasicLit)

	k.Key = getStringValue(key)

	val := args[1]

	switch a := val.(type) {
	case *ast.BasicLit:
		k.Value = getBasicValue(a)
	default:
		set := token.NewFileSet()
		var output []byte
		buffer := bytes.NewBuffer(output)
		err := format.Node(buffer, set, val)
		if err != nil {
			log.Fatal(err)
		}
		k.Value = buffer.String()
	}

	is.Keys = append(is.Keys, k)
}

func getBasicValue(basicLit *ast.BasicLit) interface{} {
	switch basicLit.Kind {
	case token.INT:
		value, err := strconv.Atoi(basicLit.Value)
		if err != nil {
			return err
		}
		return value
	case token.STRING:
		value, err := strconv.Unquote(basicLit.Value)
		if err != nil {
			return err
		}
		return value
	}
	return errors.New(fmt.Sprintf("%s is not support type", basicLit.Kind))
}

func getStringValue(basicLit *ast.BasicLit) string {
	if basicLit.Kind == token.STRING {
		value, _ := strconv.Unquote(basicLit.Value)
		return value
	}
	return ""
}

func getIntValue(basicLit *ast.BasicLit) int {
	if basicLit.Kind == token.INT {
		value, _ := strconv.Atoi(basicLit.Value)
		return value
	}
	return 0
}

func getBoolValue(name string) bool {
	v, _ := strconv.ParseBool(name)
	return v
}
