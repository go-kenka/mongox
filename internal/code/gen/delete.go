package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-kenka/mongox/internal/code/load"
)

func genDelete(base string, schema *load.Schema) error {
	dir := filepath.Join(base, strings.ToLower(schema.Name))
	genFile := filepath.Join(fmt.Sprintf("%s/%s_delete.go", dir, strings.ToLower(schema.Name)))

	// 生成之前，先删除文件
	_ = os.Remove(genFile)

	tmp := template.New("delete.tmpl")
	tmp.Funcs(template.FuncMap{
		"camelCase": CamelCase,
		"lower":     Lower,
	})
	tmp, err := tmp.ParseFS(tmpl, "template/delete.tmpl")
	if err != nil {
		return err
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	fs, err := os.OpenFile(genFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer fs.Close()

	return tmp.Execute(fs, schema)
}
