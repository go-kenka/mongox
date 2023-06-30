package gen

import (
	"strings"

	"github.com/go-kenka/mongox/types"
	"github.com/gobeam/stringy"
)

func CamelCase(str string) string {
	return stringy.New(str).CamelCase()
}
func Lower(str string) string {
	return strings.ToLower(str)
}

func GoType(t types.MongoType) string {
	return TypeGoNames[t]
}
