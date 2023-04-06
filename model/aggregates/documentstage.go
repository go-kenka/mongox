package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type DocumentsStage Stage

// Documents Changed in version 5.1. Returns literal documents from input values. The
// $documents stage has the following form:
// { $documents: <expression> }
func Documents(documents []bsonx.Bson) DocumentsStage {
	return NewDocumentsStage(documents)
}

type documentsStage struct {
	documents []bsonx.Bson
}

func (f documentsStage) Bson() bsonx.Bson {
	return f.Pro()
}

func NewDocumentsStage(documents []bsonx.Bson) documentsStage {
	return documentsStage{
		documents: documents,
	}
}

func (f documentsStage) Pro() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	documents := bsonx.Array()
	for _, value := range f.documents {
		documents.Append(value.Pro())
	}

	doc.Append("$documents", documents)
	return doc
}

func (f documentsStage) Document() bson.D {
	return f.Pro().Document()
}
