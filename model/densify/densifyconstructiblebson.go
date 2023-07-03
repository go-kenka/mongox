package densify

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type DensifyConstructibleBson struct {
	base     bsonx.Bson
	appended *bsonx.Document
}

func NewDensifyConstructibleBson(base bsonx.Bson, appended *bsonx.Document) DensifyConstructibleBson {
	a := DensifyConstructibleBson{
		base:     base,
		appended: EmptyDoc,
	}
	if !appended.IsEmpty() {
		a.appended = appended
	}
	return a
}

func (a DensifyConstructibleBson) BsonDocument() *bsonx.BsonDocument {
	baseDoc := a.base.BsonDocument()
	if baseDoc.IsEmpty() && a.appended.IsEmpty() {
		return bsonx.BsonEmpty()
	}

	if a.appended.IsEmpty() {
		return baseDoc
	}

	return bsonx.NewMerged(baseDoc, a.appended.BsonDocument())
}

func (a DensifyConstructibleBson) Document() bson.D {
	return a.BsonDocument().Document()
}

func (a DensifyConstructibleBson) newAppended(name string, value any) DensifyConstructibleBson {
	return a.newMutated(bsonx.Doc(name, value))
}

func (a DensifyConstructibleBson) newMutated(d *bsonx.Document) DensifyConstructibleBson {
	newAppended := bsonx.Empty()
	for _, v := range a.appended.Document() {
		newAppended.Append(v.Key, v.Value)
	}
	for _, v := range d.Document() {
		newAppended.Append(v.Key, v.Value)
	}

	return DensifyConstructibleBson{
		base:     a.base,
		appended: newAppended,
	}
}
func (a DensifyConstructibleBson) remove(key string) DensifyConstructibleBson {
	doc := a.base.BsonDocument()
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
	doc := bsonx.Empty()
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
