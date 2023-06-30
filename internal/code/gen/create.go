package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-kenka/mongox/internal/code/load"
)

func genCreate(base string, schema *load.Schema) error {
	dir := filepath.Join(base, strings.ToLower(schema.Name))
	genFile := filepath.Join(fmt.Sprintf("%s/%s_create.go", dir, strings.ToLower(schema.Name)))

	// 生成之前，先删除文件
	_ = os.Remove(genFile)

	tmp := template.New("create.tmpl")
	tmp.Funcs(template.FuncMap{
		"camelCase": CamelCase,
		"lower":     Lower,
	})
	tmp, err := tmp.ParseFS(tmpl, "template/create.tmpl")
	if err != nil {
		return err
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	fs, err := os.OpenFile(genFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fs.Close()

	return tmp.Execute(fs, schema)
}
