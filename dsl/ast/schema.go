package ast

import (
	"github.com/go-kenka/mongox/gen"
	"os"
	"path/filepath"
	"strings"
)

func ReadDir(path string) []*gen.Collection {
	var tbs []*gen.Collection

	files, _ := os.ReadDir(path)

	for _, entry := range files {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			return nil
		}

		defBytes, err := os.ReadFile(filepath.Join(path, entry.Name()))
		if err != nil {
			return nil
		}

		tb := astReadFile(entry.Name(), string(defBytes))

		tbs = append(tbs, tb)
	}

	return tbs
}
