package densify

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
)

type DensifyConstructibleBson struct {
	base     bsonx.Bson
	appended bsonx.Document
}

func NewDensifyConstructibleBson(base bsonx.Bson, appended bsonx.Document) DensifyConstructibleBson {
	a := DensifyConstructibleBson{
		base:     base,
		appended: EmptyDoc,
	}
	if !appended.IsEmpty() {
		a.appended = appended
	}
	return a
}

func (a DensifyConstructibleBson) ToBsonDocument() bsonx.BsonDocument {
	baseDoc := a.base.ToBsonDocument()
	if baseDoc.IsEmpty() && a.appended.IsEmpty() {
		return bsonx.NewEmptyDoc()
	}

	if a.appended.IsEmpty() {
		return baseDoc
	}

	return bsonx.NewMerged(baseDoc, a.appended.ToBsonDocument())
}

func (a DensifyConstructibleBson) Document() bsonx.Document {
	return a.ToBsonDocument().Document()
}

func (a DensifyConstructibleBson) newAppended(name string, value any) DensifyConstructibleBson {
	return a.newMutated(bsonx.NewDocument(name, value))
}

func (a DensifyConstructibleBson) newMutated(d bsonx.Document) DensifyConstructibleBson {
	newAppended := bsonx.NewDoc()
	for k, v := range a.appended {
		newAppended.Append(k, v)
	}
	for k, v := range d {
		newAppended.Append(k, v)
	}

	return DensifyConstructibleBson{
		base:     a.base,
		appended: newAppended,
	}
}
func (a DensifyConstructibleBson) remove(key string) DensifyConstructibleBson {
	doc := a.base.ToBsonDocument()
	doc.Remove(key)
	appended := a.appended
	appended.Remove(key)
	return DensifyConstructibleBson{base: doc, appended: appended}
}
func (a DensifyConstructibleBson) of(doc bsonx.Bson) DensifyConstructibleBson {
	d, ok := doc.(DensifyConstructibleBson)
	if ok {
		return d
	}
	return NewDensifyConstructibleBson(doc, nil)
}
func (a DensifyConstructibleBson) PartitionByFields(fields ...string) DensifyOptions {
	doc := bsonx.NewDoc()
	if len(fields) > 0 {
		doc.Append("partitionByFields", fields)
	} else {
		a.remove("partitionByFields")
	}

	return a.newMutated(doc)
}
func (a DensifyConstructibleBson) Option(name string, value any) DensifyOptions {
	return a.newAppended(name, value)
}
