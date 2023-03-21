package main

import (
	"fmt"
	"github.com/go-kenka/mongox/dsl/ast"
)

func main() {
	cs := ast.ReadDir("examples/data/schema")

	fmt.Println(cs)

}
