package gen

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-kenka/mongox/internal/code/load"
)

//go:embed template/*
var tmpl embed.FS

func GenClient(fullPath string, spec *load.SchemaSpec) error {

	genFile := filepath.Join(spec.PkgPath, "client.go")

	// 删除旧文件
	_ = os.Remove(genFile)

	for i, schema := range spec.Schemas {
		fmt.Printf("正在生成第%d【%s】个表的数据\n", i+1, schema.Name)
		if err := genData(spec.PkgPath, schema); err != nil {
			return err
		}
		if err := genCreate(spec.PkgPath, schema); err != nil {
			return err
		}
		if err := genDelete(spec.PkgPath, schema); err != nil {
			return err
		}
		if err := genQuery(spec.PkgPath, schema); err != nil {
			return err
		}
		if err := genUpdate(spec.PkgPath, schema); err != nil {
			return err
		}
		if err := genAggregate(spec.PkgPath, schema); err != nil {
			return err
		}
		fmt.Printf("正在生成第%d个表的数据生成成功\n", i+1)
	}

	fmt.Printf("正在生成client的数据\n")
	// 恢复原路径
	spec.PkgPath = fullPath

	tmp := template.New("client.tmpl")
	tmp.Funcs(template.FuncMap{
		"camelCase": CamelCase,
		"lower":     Lower,
	})
	tmp, err := tmp.ParseFS(tmpl, "template/client.tmpl")
	if err != nil {
		return err
	}

	fs, err := os.OpenFile(genFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer fs.Close()

	err = tmp.Execute(fs, &spec)
	if err != nil {
		return err
	}

	fmt.Printf("生成client的数据成功\n")
	return nil
}
