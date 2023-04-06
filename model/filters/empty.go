package filters

import "github.com/go-kenka/mongox/bsonx"

type emptyFilter struct {
	filter bsonx.Bson
}

func (f emptyFilter) Value() bsonx.IBsonValue {
	return f.filter.ToBsonDocument()
}

func Empty() emptyFilter {
	return emptyFilter{filter: bsonx.BsonEmpty()}
}
