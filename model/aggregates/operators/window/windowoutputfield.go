package window

import (
	"github.com/go-kenka/mongox/internal/expression"
)

type WindowOutputField interface {
	expression.Expression
	outputOperator
}
