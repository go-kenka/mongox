package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type DocumentsStage struct {
	stage bsonx.Bson
}

func (s DocumentsStage) Bson() bsonx.Bson {
	return s.stage
}

func (s DocumentsStage) Document() bson.D {
	return s.stage.Document()
}
func (s DocumentsStage) Database() {
}

// Documents Changed in version 5.1. Returns literal documents from input values. The
// $documents DefaultStage has the following form:
// { $documents: <expression> }
func Documents(documents []bsonx.Bson) DocumentsStage {
	return DocumentsStage{stage: NewDocumentsStage(documents)}
}

type documentsStage struct {
	documents []bsonx.Bson
}

func NewDocumentsStage(documents []bsonx.Bson) documentsStage {
	return documentsStage{
		documents: documents,
	}
}

func (f documentsStage) BsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	documents := bsonx.Array()
	for _, value := range f.documents {
		documents.Append(value.BsonDocument())
	}

	doc.Append("$documents", documents)
	return doc
}

func (f documentsStage) Document() bson.D {
	return f.BsonDocument().Document()
}
