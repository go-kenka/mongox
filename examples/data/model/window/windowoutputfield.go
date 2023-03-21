package window

import "github.com/go-kenka/mongox/examples/data/bsonx"

type WindowOutputField interface {
	ToBsonField() bsonx.BsonField
}
