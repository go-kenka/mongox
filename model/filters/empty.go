package filters

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type emptyFilter struct {
	filter bsonx.Bson
}

func (f emptyFilter) Value() bsonx.IBsonValue {
	return f.filter.BsonDocument()
}

func (f emptyFilter) Document() bson.D {
	return f.filter.Document()
}

func Empty() emptyFilter {
	return emptyFilter{filter: bsonx.BsonEmpty()}
}
