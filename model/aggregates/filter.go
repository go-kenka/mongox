package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type simpleFilter struct {
	fieldName string
	value     bsonx.IBsonValue
}

func NewSimpleFilter(fieldName string, value bsonx.IBsonValue) simpleFilter {
	return simpleFilter{
		fieldName: fieldName,
		value:     value,
	}
}

func (s simpleFilter) BsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(s.fieldName, s.value)
}

func (s simpleFilter) Document() bson.D {
	return s.BsonDocument().Document()
}

type operatorFilter[T bsonx.Expression] struct {
	operatorName string
	fieldName    string
	value        T
}

func NewOperatorFilter[T bsonx.Expression](operatorName string, fieldName string, value T) operatorFilter[T] {
	return operatorFilter[T]{
		operatorName: operatorName,
		fieldName:    fieldName,
		value:        value,
	}
}

func (s operatorFilter[T]) BsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	operator := bsonx.BsonDoc(s.operatorName, s.value)
	doc.Append(s.fieldName, operator)
	return doc
}
func (s operatorFilter[T]) Document() bson.D {
	return s.BsonDocument().Document()
}
