package fill

import (
	"github.com/go-kenka/mongox/bsonx"
)

type FillOutputField interface {
	bsonx.Bson
}

func Value(field string, expression bsonx.Expression) ValueFillOutputField {
	return NewFillConstructibleBsonElement(field, bsonx.Doc("value", expression), FillConstructibleBson{})
}

func Locf(field string) LocfFillOutputField {
	return NewFillConstructibleBsonElement(field, bsonx.Doc("method", "locf"), FillConstructibleBson{})
}

func Linear(field string) LinearFillOutputField {
	return NewFillConstructibleBsonElement(field, bsonx.Doc("method", "linear"), FillConstructibleBson{})
}

func Of(fill bsonx.Bson) FillOutputField {
	return NewFillConstructibleBsonElement("fill", bsonx.BsonEmpty(), FillConstructibleBson{})
}
