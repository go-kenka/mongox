package projections

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/expressions"
)

func Computed(fieldName string, expression expressions.TExpression) bsonx.Bson {
	return expressions.NewSimpleExpression(fieldName, expression)
}

func ComputedSearchMeta(fieldName string) bsonx.Bson {
	return Computed(fieldName, bsonx.NewBsonString("$$SEARCH_META"))
}

func Include(fieldNames ...string) bsonx.Bson {
	return Combine(fieldNames, bsonx.NewBsonInt32(0))
}

func ExcludeId() bsonx.Bson {
	return bsonx.NewBsonDocument("_id", bsonx.NewBsonInt32(0))
}

func ElemMatch(fieldName string) bsonx.Bson {
	return bsonx.NewBsonDocument(fieldName+".$", bsonx.NewBsonInt32(1))
}

func ElemMatchWithFilter(fieldName string, filter bsonx.Bson) bsonx.Bson {
	return NewElemMatchFilterProjection(fieldName, filter)
}

func Meta(fieldName string, metaFieldName string) bsonx.Bson {
	return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonDocument("$meta", bsonx.NewBsonString(metaFieldName)))
}

func MetaTextScore(fieldName string) bsonx.Bson {
	return Meta(fieldName, "textScore")
}

func MetaSearchScore(fieldName string) bsonx.Bson {
	return Meta(fieldName, "searchScore")
}

func MetaSearchHighlights(fieldName string) bsonx.Bson {
	return Meta(fieldName, "searchHighlights")
}
func Slice(fieldName string, limit int32) bsonx.Bson {
	return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonDocument("$slice", bsonx.NewBsonInt32(limit)))
}
func SliceWithSkip(fieldName string, limit, skip int32) bsonx.Bson {
	return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonDocument("$slice", bsonx.NewBsonArray(
		bsonx.NewBsonInt32(skip), bsonx.NewBsonInt32(limit))))
}

func Fields(projections []bsonx.Bson) bsonx.Bson {
	return NewFieldsProjection(projections)
}

func Combine(fieldNames []string, value bsonx.IBsonValue) bsonx.Bson {
	doc := bsonx.NewEmptyDoc()
	for _, name := range fieldNames {
		doc.Append(name, value)
	}
	return doc
}

type FieldsProjection struct {
	projections []bsonx.Bson
}

func NewFieldsProjection(projections []bsonx.Bson) FieldsProjection {
	return FieldsProjection{
		projections: projections,
	}
}

func (p FieldsProjection) ToBsonDocument() bsonx.BsonDocument {
	combinedDocument := bsonx.NewEmptyDoc()
	for _, sort := range p.projections {
		sortDocument := sort.ToBsonDocument()
		for _, key := range sortDocument.Keys() {
			combinedDocument.Append(key, sortDocument.GetValue(key))
		}
	}
	return combinedDocument
}

func (p FieldsProjection) Document() bsonx.Document {
	return p.ToBsonDocument().Document()
}

type ElemMatchFilterProjection struct {
	fieldName string
	filter    bsonx.Bson
}

func NewElemMatchFilterProjection(fieldName string, filter bsonx.Bson) ElemMatchFilterProjection {
	return ElemMatchFilterProjection{
		filter:    filter,
		fieldName: fieldName,
	}
}

func (p ElemMatchFilterProjection) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewBsonDocument(p.fieldName, bsonx.NewBsonDocument("$elemMatch", p.filter.ToBsonDocument()))
}

func (p ElemMatchFilterProjection) Document() bsonx.Document {
	return p.ToBsonDocument().Document()
}
