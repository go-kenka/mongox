package size

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type sizeOperator struct {
	doc bsonx.Bson
}

func (o sizeOperator) Exp() bsonx.IBsonValue {
	return o.doc.Pro()
}

func BinarySize[T expression.BinaryExpression](n T) sizeOperator {
	return sizeOperator{doc: bsonx.BsonDoc("$binarySize", n)}
}

func BsonSize[T expression.ObjectExpression](n T) sizeOperator {
	return sizeOperator{doc: bsonx.BsonDoc("$bsonSize", n)}
}
