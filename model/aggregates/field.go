package aggregates

import (
	"github.com/go-kenka/mongox/bsonx/expression"
)

type Field[T expression.AnyExpression] struct {
	name  string
	value expression.AnyExpression
}

func NewField[T expression.AnyExpression](name string, value T) Field[T] {
	return Field[T]{
		name:  name,
		value: value,
	}
}
