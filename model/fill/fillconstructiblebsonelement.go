package fill

import (
	"fmt"

	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type FillConstructibleBsonElement struct {
	baseElement          bsonx.Bson
	appendedElementValue FillConstructibleBson
}

func NewFillConstructibleBsonElement(name string, value bsonx.Bson, appendedElementValue bsonx.Bson) FillConstructibleBsonElement {
	return FillConstructibleBsonElement{
		baseElement:          bsonx.Doc(name, value),
		appendedElementValue: FillConstructibleBson{}.of(appendedElementValue),
	}
}
func (a FillConstructibleBsonElement) newSelf(baseElement bsonx.Bson, appendedElementValue bsonx.Bson) FillConstructibleBsonElement {
	return FillConstructibleBsonElement{
		baseElement:          baseElement,
		appendedElementValue: FillConstructibleBson{}.of(appendedElementValue),
	}
}

func (a FillConstructibleBsonElement) newWithAppendedValue(name string, value any) FillConstructibleBsonElement {
	return a.newWithMutatedValue(bsonx.Doc(name, value))
}

func (a FillConstructibleBsonElement) newWithMutatedValue(d bsonx.Document) FillConstructibleBsonElement {
	return a.newSelf(a.baseElement, a.appendedElementValue.newMutated(d))
}

func (a FillConstructibleBsonElement) Pro() *bsonx.BsonDocument {
	baseDoc := a.baseElement.Pro()
	if baseDoc.Size() != 1 {
		panic(fmt.Errorf("baseElement must contain exactly one element, but contains %d", baseDoc.Size()))
	}

	baseElementValueDoc := bsonx.BsonEmpty()
	baseElementName := ""
	for k, baseElementValue := range baseDoc.Data() {
		baseElementName = k
		v, ok := baseElementValue.(bsonx.BsonValue)
		if !ok {
			panic(fmt.Errorf("baseElement value must be a document, but it is %v", v))
		}
		if !v.IsDocument() {
			panic(fmt.Errorf("baseElement value must be a document, but it is %v", v.GetBsonType()))
		}
		baseElementValueDoc = v.AsDocument()
	}
	appendedElementValueDoc := a.appendedElementValue.Pro()
	if appendedElementValueDoc.IsEmpty() {
		return baseDoc
	}
	return bsonx.BsonDoc(baseElementName, bsonx.NewMerged(baseElementValueDoc, appendedElementValueDoc))
}

func (a FillConstructibleBsonElement) Document() bson.D {
	return a.Pro().Document()
}
