package fill

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/expressions"
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

func (a FillConstructibleBson) ToBsonDocument() bsonx.BsonDocument {
	baseDoc := a.base.ToBsonDocument()
	if baseDoc.IsEmpty() && a.appended.IsEmpty() {
		return bsonx.NewEmptyDoc()
	}

	if a.appended.IsEmpty() {
		return baseDoc
	}

	return bsonx.NewMerged(baseDoc, a.appended.ToBsonDocument())
}

func (a FillConstructibleBson) Document() bsonx.Document {
	return a.ToBsonDocument().Document()
}

func (a FillConstructibleBson) newAppended(name string, value any) FillConstructibleBson {
	return a.newMutated(bsonx.NewDocument(name, value))
}

func (a FillConstructibleBson) newMutated(d bsonx.Document) FillConstructibleBson {
	newAppended := bsonx.NewDoc()
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
	doc := a.base.ToBsonDocument()
	doc.Remove(key)
	appended := a.appended
	appended.Remove(key)
	return FillConstructibleBson{base: doc, appended: a.appended}
}

func (a FillConstructibleBson) PartitionBy(expression expressions.TExpression) FillOptions {
	doc := bsonx.NewDoc()
	a.remove("partitionByFields")
	doc.Append("partitionBy", expression)
	return a.newMutated(doc)
}
func (a FillConstructibleBson) PartitionByFields(fields ...string) FillOptions {
	a.remove("partitionBy")
	doc := bsonx.NewDoc()
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
