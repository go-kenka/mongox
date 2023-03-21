package expressions

import "github.com/go-kenka/mongox/examples/data/bsonx"

type SimpleExpression struct {
	name       string
	expression TExpression
}

func NewSimpleExpression(name string, expr TExpression) SimpleExpression {
	return SimpleExpression{
		name:       name,
		expression: expr,
	}
}

func (s SimpleExpression) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewBsonDocument(s.name, s.expression)
}

func (s SimpleExpression) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}
