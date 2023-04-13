package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type UnionWithStage struct {
	stage bsonx.Bson
}

func (s UnionWithStage) Bson() bsonx.Bson {
	return s.stage
}

func (s UnionWithStage) Document() bson.D {
	return s.stage.Document()
}

// UnionWith NewDefaultStage in version 4.4.
// Performs a union of two collections.
// $unionWith
// combines pipeline results from two collections into a single result set. The DefaultStage outputs the combined result set (including duplicates) to the next DefaultStage.
// The order in which the combined result set documents are output is unspecified.
func UnionWith(collection string, pipeline ...bsonx.Bson) UnionWithStage {
	return UnionWithStage{stage: NewUnionWithStage(collection, pipeline...)}
}

type unionWithStage struct {
	collection string
	pipeline   []bsonx.Bson
}

func NewUnionWithStage(collection string, pipeline ...bsonx.Bson) unionWithStage {
	return unionWithStage{
		collection: collection,
		pipeline:   pipeline,
	}
}

func (f unionWithStage) BsonDocument() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonEmpty()
	data.Append("coll", bsonx.String(f.collection))

	pipeline := bsonx.Array()
	for _, s := range f.pipeline {
		pipeline.Append(s.BsonDocument())
	}
	data.Append("pipeline", pipeline)

	b.Append("$unionWith", data)
	return b
}

func (f unionWithStage) Document() bson.D {
	return f.BsonDocument().Document()
}
