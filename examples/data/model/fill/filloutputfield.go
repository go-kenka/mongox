package fill

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/expressions"
)

type FillOutputField interface {
	bsonx.Bson
}

func Value(field string, expression expressions.TExpression) ValueFillOutputField {
	return NewFillConstructibleBsonElement(field, bsonx.NewDocument("value", expression), FillConstructibleBson{})
}

func Locf(field string) LocfFillOutputField {
	return NewFillConstructibleBsonElement(field, bsonx.NewDocument("method", "locf"), FillConstructibleBson{})
}

func Linear(field string) LinearFillOutputField {
	return NewFillConstructibleBsonElement(field, bsonx.NewDocument("method", "linear"), FillConstructibleBson{})
}

func Of(fill bsonx.Bson) FillOutputField {
	return NewFillConstructibleBsonElement("fill", bsonx.NewEmptyDoc(), FillConstructibleBson{})
}
