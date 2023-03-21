package sorts

import "github.com/go-kenka/mongox/examples/data/bsonx"

func Ascending(fieldNames ...string) bsonx.Bson {
	return OrderByFiled(fieldNames, bsonx.NewBsonInt32(1))
}
func Descending(fieldNames ...string) bsonx.Bson {
	return OrderByFiled(fieldNames, bsonx.NewBsonInt32(-1))
}
func MetaTextScore(fieldName string) bsonx.Bson {
	return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonDocument("$meta", bsonx.NewBsonString("textScore")))
}
func OrderBy(sorts ...bsonx.Bson) bsonx.Bson {
	return NewCompoundSort(sorts...)
}
func OrderByFiled(fieldNames []string, value bsonx.IBsonValue) bsonx.Bson {
	document := bsonx.NewEmptyDoc()
	for _, fieldName := range fieldNames {
		document.Append(fieldName, value)
	}
	return document
}

type CompoundSort struct {
	sorts []bsonx.Bson
}

func NewCompoundSort(sorts ...bsonx.Bson) CompoundSort {
	return CompoundSort{sorts: sorts}
}

func (s CompoundSort) ToBsonDocument() bsonx.BsonDocument {
	combinedDocument := bsonx.NewEmptyDoc()
	for _, sort := range s.sorts {
		sortDocument := sort.ToBsonDocument()
		for _, key := range sortDocument.Keys() {
			combinedDocument.Append(key, sortDocument.GetValue(key))
		}
	}
	return combinedDocument
}

func (s CompoundSort) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}
