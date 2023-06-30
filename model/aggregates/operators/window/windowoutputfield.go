package window

import (
	"github.com/go-kenka/mongox/bsonx/expression"
)

type WindowOutputField interface {
	expression.Expression
	outputOperator
}
