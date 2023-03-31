package expression

import (
	"github.com/go-kenka/mongox/bsonx"
)

type ObjExpression struct {
	bsonx.BsonValue
	doc *bsonx.BsonDocument
}

func Obj[T AnyExpression](fieldName string, expr T) ObjExpression {
	return ObjExpression{
		doc: bsonx.BsonDoc(fieldName, expr),
	}
}

func (o ObjExpression) Exp() bsonx.IBsonValue {
	return o.doc
}
