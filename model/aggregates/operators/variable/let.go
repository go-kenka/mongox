package variable

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type variableOperator struct {
	doc bsonx.Bson
}

func (o variableOperator) Exp() bsonx.IBsonValue {
	return o.doc.Pro()
}

func Let[T expression.AnyExpression](vars bsonx.Bson, in T) variableOperator {
	return variableOperator{doc: bsonx.BsonDoc("$let",
		bsonx.BsonDoc("vars", vars.Pro()).Append("in", in))}
}
