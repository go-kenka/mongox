package aggregates

import (
	"github.com/go-kenka/mongox/examples/data/model/expressions"
)

type Field[T expressions.TExpression] struct {
	name  string
	value expressions.TExpression
}
