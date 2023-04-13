package indexs

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

func Ascending(fieldNames ...string) bsonx.Bson {
	return CompoundIndexWithField(fieldNames, bsonx.Int32(1))
}

func Descending(fieldNames ...string) bsonx.Bson {
	return CompoundIndexWithField(fieldNames, bsonx.Int32(-1))
}
func Geo2DSphere(fieldNames ...string) bsonx.Bson {
	return CompoundIndexWithField(fieldNames, bsonx.String("2dsphere"))
}
func Geo2D(fieldName string) bsonx.Bson {
	return bsonx.BsonDoc(fieldName, bsonx.String("2d"))
}
func GeoHaystack(fieldName string, additional bsonx.Bson) bsonx.Bson {
	return CompoundIndex(bsonx.BsonDoc(fieldName, bsonx.String("geoHaystack")), additional)
}
func Text(fieldName string) bsonx.Bson {
	return bsonx.BsonDoc(fieldName, bsonx.String("text"))
}
func TextEveyFiled() bsonx.Bson {
	return bsonx.BsonDoc("$**", bsonx.String("text"))
}
func Hashed(fieldName string) bsonx.Bson {
	return bsonx.BsonDoc(fieldName, bsonx.String("hashed"))
}
func CompoundIndex(indexes ...bsonx.Bson) bsonx.Bson {
	return newCompoundIndex(indexes...)
}
func CompoundIndexWithField(fieldNames []string, value bsonx.IBsonValue) bsonx.Bson {
	doc := bsonx.BsonEmpty()
	for _, name := range fieldNames {
		doc.Append(name, value)
	}
	return doc
}

type compoundIndex struct {
	indexes []bsonx.Bson
}

func newCompoundIndex(indexes ...bsonx.Bson) bsonx.Bson {
	return compoundIndex{indexes: indexes}
}

func (i compoundIndex) BsonDocument() *bsonx.BsonDocument {
	c := bsonx.BsonEmpty()
	for _, index := range i.indexes {
		indexDocument := index.BsonDocument()
		for _, key := range indexDocument.Keys() {
			c.Append(key, indexDocument.GetValue(key))
		}
	}
	return c
}

func (i compoundIndex) Document() bson.D {
	return i.BsonDocument().Document()
}
