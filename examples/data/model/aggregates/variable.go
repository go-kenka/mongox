package aggregates

import (
	"github.com/go-kenka/mongox/examples/data/model/expressions"
)

type Variable[T expressions.TExpression] struct {
	name  string
	value T
}

func NewVariable[T expressions.TExpression](name string, value T) Variable[T] {
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
