package options

import (
	"github.com/go-kenka/mongox/internal/expression"
)

type Variable[T expression.AnyExpression] struct {
	name  string
	value T
}

func NewVariable[T expression.AnyExpression](name string, value T) Variable[T] {
	return Variable[T]{
		name:  name,
		value: value,
	}
}

func (v Variable[T]) GetName() string {
	return v.name
}

func (v Variable[T]) GetValue() T {
	return v.value
}
