package fill

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type FillConstructibleBson struct {
	base     bsonx.Bson
	appended bsonx.Document
}

func NewFillConstructibleBson(base bsonx.Bson, appended bsonx.Document) FillConstructibleBson {
	a := FillConstructibleBson{
		base:     base,
		appended: EmptyDoc,
	}
	if !appended.IsEmpty() {
		a.appended = appended
	}
	return a
}

func (a FillConstructibleBson) Pro() *bsonx.BsonDocument {
	baseDoc := a.base.Pro()
	if baseDoc.IsEmpty() && a.appended.IsEmpty() {
		return bsonx.BsonEmpty()
	}

	if a.appended.IsEmpty() {
		return baseDoc
	}

	return bsonx.NewMerged(baseDoc, a.appended.Pro())
}

func (a FillConstructibleBson) Document() bson.D {
	return a.Pro().Document()
}

func (a FillConstructibleBson) newAppended(name string, value any) FillConstructibleBson {
	return a.newMutated(bsonx.Doc(name, value))
}

func (a FillConstructibleBson) newMutated(d bsonx.Document) FillConstructibleBson {
	newAppended := bsonx.Empty()
	for k, v := range a.appended {
		newAppended.Append(k, v)
	}
	for k, v := range d {
		newAppended.Append(k, v)
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
	doc := a.base.Pro()
	doc.Remove(key)
	appended := a.appended
	appended.Remove(key)
	return FillConstructibleBson{base: doc, appended: a.appended}
}

func (a FillConstructibleBson) PartitionBy(expression bsonx.Expression) FillOptions {
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
