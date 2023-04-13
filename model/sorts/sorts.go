package sorts

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

func Ascending(fieldNames ...string) bsonx.Bson {
	return OrderByFiled(fieldNames, bsonx.Int32(1))
}
func Descending(fieldNames ...string) bsonx.Bson {
	return OrderByFiled(fieldNames, bsonx.Int32(-1))
}
func MetaTextScore(fieldName string) bsonx.Bson {
	return bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$meta", bsonx.String("textScore")))
}
func OrderBy(sorts ...bsonx.Bson) bsonx.Bson {
	return NewCompoundSort(sorts...)
}
func OrderByFiled(fieldNames []string, value bsonx.IBsonValue) bsonx.Bson {
	document := bsonx.BsonEmpty()
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

func (s CompoundSort) BsonDocument() *bsonx.BsonDocument {
	combinedDocument := bsonx.BsonEmpty()
	for _, sort := range s.sorts {
		sortDocument := sort.BsonDocument()
		for _, key := range sortDocument.Keys() {
			combinedDocument.Append(key, sortDocument.GetValue(key))
		}
	}
	return combinedDocument
}

func (s CompoundSort) Document() bson.D {
	return s.BsonDocument().Document()
}
