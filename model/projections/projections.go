package projections

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

func Computed[E bsonx.Expression](fieldName string, expression E) bsonx.Bson {
	// return expr.NewSimpleExpression(fieldName, expression)
	return nil
}

func ComputedSearchMeta(fieldName string) bsonx.Bson {
	return Computed(fieldName, bsonx.String("$$SEARCH_META"))
}

func Include(fieldNames ...string) bsonx.Bson {
	return Combine(fieldNames, bsonx.Int32(0))
}

func ExcludeId() bsonx.Bson {
	return bsonx.BsonDoc("_id", bsonx.Int32(0))
}

func ElemMatch(fieldName string) bsonx.Bson {
	return bsonx.BsonDoc(fieldName+".$", bsonx.Int32(1))
}

func ElemMatchWithFilter(fieldName string, filter bsonx.Bson) bsonx.Bson {
	return NewElemMatchFilterProjection(fieldName, filter)
}

func Meta(fieldName string, metaFieldName string) bsonx.Bson {
	return bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$meta", bsonx.String(metaFieldName)))
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
	return bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$slice", bsonx.Int32(limit)))
}
func SliceWithSkip(fieldName string, limit, skip int32) bsonx.Bson {
	return bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$slice", bsonx.Array(
		bsonx.Int32(skip), bsonx.Int32(limit))))
}

func Fields(projections []bsonx.Bson) bsonx.Bson {
	return NewFieldsProjection(projections)
}

func Combine(fieldNames []string, value bsonx.IBsonValue) bsonx.Bson {
	doc := bsonx.BsonEmpty()
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

func (p FieldsProjection) ToBsonDocument() *bsonx.BsonDocument {
	combinedDocument := bsonx.BsonEmpty()
	for _, sort := range p.projections {
		sortDocument := sort.ToBsonDocument()
		for _, key := range sortDocument.Keys() {
			combinedDocument.Append(key, sortDocument.GetValue(key))
		}
	}
	return combinedDocument
}

func (p FieldsProjection) Document() bson.D {
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

func (p ElemMatchFilterProjection) ToBsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(p.fieldName, bsonx.BsonDoc("$elemMatch", p.filter.ToBsonDocument()))
}

func (p ElemMatchFilterProjection) Document() bson.D {
	return p.ToBsonDocument().Document()
}
