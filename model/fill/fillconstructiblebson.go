package fill

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"go.mongodb.org/mongo-driver/bson"
)

type FillConstructibleBson struct {
	base     bsonx.Bson
	appended *bsonx.Document
}

func NewFillConstructibleBson(base bsonx.Bson, appended *bsonx.Document) FillConstructibleBson {
	a := FillConstructibleBson{
		base:     base,
		appended: EmptyDoc,
	}
	if !appended.IsEmpty() {
		a.appended = appended
	}
	return a
}

func (a FillConstructibleBson) BsonDocument() *bsonx.BsonDocument {
	baseDoc := a.base.BsonDocument()
	if baseDoc.IsEmpty() && a.appended.IsEmpty() {
		return bsonx.BsonEmpty()
	}

	if a.appended.IsEmpty() {
		return baseDoc
	}

	return bsonx.NewMerged(baseDoc, a.appended.BsonDocument())
}

func (a FillConstructibleBson) Document() bson.D {
	return a.BsonDocument().Document()
}

func (a FillConstructibleBson) newAppended(name string, value any) FillConstructibleBson {
	return a.newMutated(bsonx.Doc(name, value))
}

func (a FillConstructibleBson) newMutated(d *bsonx.Document) FillConstructibleBson {
	newAppended := bsonx.Empty()
	for _, v := range a.appended.Document() {
		newAppended.Append(v.Key, v.Value)
	}
	for _, v := range d.Document() {
		newAppended.Append(v.Key, v.Value)
	}

	return FillConstructibleBson{
		base:     a.base,
		appended: newAppended,
	}
}
func (a FillConstructibleBson) of(doc bsonx.Bson) FillConstructibleBson {
	d, ok := doc.(FillConstructibleBson)
	if ok {
		return d
	}
	return NewFillConstructibleBson(doc, nil)
}
func (a FillConstructibleBson) remove(key string) FillConstructibleBson {
	doc := a.base.BsonDocument()
	doc.Remove(key)
	appended := a.appended
	appended.Remove(key)
	return FillConstructibleBson{base: doc, appended: a.appended}
}

func (a FillConstructibleBson) PartitionBy(expression expression.AnyExpression) FillOptions {
	doc := bsonx.Empty()
	a.remove("partitionByFields")
	doc.Append("partitionBy", expression)
	return a.newMutated(doc)
}
func (a FillConstructibleBson) PartitionByFields(fields ...string) FillOptions {
	a.remove("partitionBy")
	doc := bsonx.Empty()
	if len(fields) > 0 {
		doc.Append("partitionByFields", fields)
	} else {
		a.remove("partitionByFields")
	}

	return a.newMutated(doc)
}
func (a FillConstructibleBson) SortBy(sortBy bsonx.Bson) FillOptions {
	return a.newAppended("sortBy", sortBy)
}

func (a FillConstructibleBson) Option(name string, value any) FillOptions {
	return a.newAppended(name, value)
}
