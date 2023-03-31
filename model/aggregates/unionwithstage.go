package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type UnionWithStage Stage

// UnionWith NewStage in version 4.4.
// Performs a union of two collections.
// $unionWith
// combines pipeline results from two collections into a single result set. The stage outputs the combined result set (including duplicates) to the next stage.
// The order in which the combined result set documents are output is unspecified.
func UnionWith(collection string, pipeline ...bsonx.Bson) UnionWithStage {
	return NewUnionWithStage(collection, pipeline...)
}

type unionWithStage struct {
	collection string
	pipeline   []bsonx.Bson
}

func (f unionWithStage) Bson() bsonx.Bson {
	return f.ToBsonDocument()
}

func NewUnionWithStage(collection string, pipeline ...bsonx.Bson) unionWithStage {
	return unionWithStage{
		collection: collection,
		pipeline:   pipeline,
	}
}

func (f unionWithStage) ToBsonDocument() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonEmpty()
	data.Append("coll", bsonx.String(f.collection))

	var pipeline bsonx.BsonArray
	for _, s := range f.pipeline {
		pipeline.Append(s.ToBsonDocument())
	}
	data.Append("pipeline", pipeline)

	b.Append("$unionWith", data)
	return b
}

func (f unionWithStage) Document() bson.D {
	return f.ToBsonDocument().Document()
}
