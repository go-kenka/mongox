// Package literal Return a value without parsing. Use for values that the aggregation pipeline may interpret as an expression.
// For example, use a $literal expression to a string that starts with a dollar sign ($) to avoid parsing as a field path.
package expression

import (
	"github.com/go-kenka/mongox/bsonx"
)

type LiteralExpression struct {
	bsonx.BsonValue
	doc *bsonx.BsonDocument
}

// Literal Returns a value without parsing. Use for values that the aggregation pipeline may interpret as an expression.
// The $literal expression has the following syntax:
// { $literal: <value> }
func Literal[T AnyExpression](expr T) LiteralExpression {
	return LiteralExpression{
		doc: bsonx.BsonDoc("$literal", expr),
	}
}

func (l LiteralExpression) Exp() bsonx.IBsonValue {
	return l.doc
}
