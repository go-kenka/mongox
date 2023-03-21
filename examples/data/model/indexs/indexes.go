package indexs

import "github.com/go-kenka/mongox/examples/data/bsonx"

type indexes struct {
}

func Indexes() indexes {
	return indexes{}
}

func (i indexes) Ascending(fieldNames ...string) bsonx.Bson {
	return i.CompoundIndexWithField(fieldNames, bsonx.NewBsonInt32(1))
}

func (i indexes) Descending(fieldNames ...string) bsonx.Bson {
	return i.CompoundIndexWithField(fieldNames, bsonx.NewBsonInt32(-1))
}
func (i indexes) Geo2DSphere(fieldNames ...string) bsonx.Bson {
	return i.CompoundIndexWithField(fieldNames, bsonx.NewBsonString("2dsphere"))
}
func (i indexes) Geo2D(fieldName string) bsonx.Bson {
	return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonString("2d"))
}
func (i indexes) GeoHaystack(fieldName string, additional bsonx.Bson) bsonx.Bson {
	return i.CompoundIndex(bsonx.NewBsonDocument(fieldName, bsonx.NewBsonString("geoHaystack")), additional)
}
func (i indexes) Text(fieldName string) bsonx.Bson {
	return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonString("text"))
}
func (i indexes) TextEveyFiled() bsonx.Bson {
	return bsonx.NewBsonDocument("$**", bsonx.NewBsonString("text"))
}
func (i indexes) Hashed(fieldName string) bsonx.Bson {
	return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonString("hashed"))
}
func (i indexes) CompoundIndex(indexes ...bsonx.Bson) bsonx.Bson {
	return CompoundIndex(indexes...)
}
func (i indexes) CompoundIndexWithField(fieldNames []string, value bsonx.IBsonValue) bsonx.Bson {
	doc := bsonx.NewEmptyDoc()
	for _, name := range fieldNames {
		doc.Append(name, value)
	}
	return doc
}

type compoundIndex struct {
	indexes []bsonx.Bson
}

func CompoundIndex(indexes ...bsonx.Bson) bsonx.Bson {
	return compoundIndex{indexes: indexes}
}

func (i compoundIndex) ToBsonDocument() bsonx.BsonDocument {
	c := bsonx.NewEmptyDoc()
	for _, index := range i.indexes {
		indexDocument := index.ToBsonDocument()
		for _, key := range indexDocument.Keys() {
			c.Append(key, indexDocument.GetValue(key))
		}
	}
	return c
}

func (i compoundIndex) Document() bsonx.Document {
	return i.ToBsonDocument().Document()
}
