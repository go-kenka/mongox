package projections

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/model/aggregates"
	"go.mongodb.org/mongo-driver/bson"
)

type Projections interface {
	Pro() *bsonx.BsonDocument
}

type fields struct {
	doc bsonx.Bson
}

func (p fields) Pro() *bsonx.BsonDocument {
	return p.doc.Pro()
}

func Computed[E expression.AnyExpression](fieldName string, expression E) fields {
	return fields{doc: aggregates.NewSimpleFilter(fieldName, expression)}
}

func ComputedSearchMeta(fieldName string) fields {
	return Computed(fieldName, bsonx.String("$$SEARCH_META"))
}

func Include(fieldNames ...string) fields {
	return combine(fieldNames, bsonx.Int32(1))
}

func ExcludeId() fields {
	return fields{doc: bsonx.BsonDoc("_id", bsonx.Int32(0))}
}

func ElemMatch(fieldName string) fields {
	return fields{doc: bsonx.BsonDoc(fieldName+".$", bsonx.Int32(1))}
}

func ElemMatchWithFilter(fieldName string, filter bsonx.Bson) fields {
	return fields{doc: newElemMatchFilterProjection(fieldName, filter)}
}

func Meta(fieldName string, metaFieldName string) fields {
	return fields{doc: bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$meta", bsonx.String(metaFieldName)))}
}

func MetaTextScore(fieldName string) fields {
	return Meta(fieldName, "textScore")
}

func MetaSearchScore(fieldName string) fields {
	return Meta(fieldName, "searchScore")
}

func MetaSearchHighlights(fieldName string) fields {
	return Meta(fieldName, "searchHighlights")
}
func Slice(fieldName string, limit int32) fields {
	return fields{doc: bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$slice", bsonx.Int32(limit)))}
}
func SliceWithSkip(fieldName string, limit, skip int32) fields {
	return fields{doc: bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$slice", bsonx.Array(
		bsonx.Int32(skip), bsonx.Int32(limit))))}
}

func Fields(projections ...fields) Projections {
	return newFieldsProjection(projections)
}

func combine(fieldNames []string, value bsonx.IBsonValue) fields {
	doc := bsonx.BsonEmpty()
	for _, name := range fieldNames {
		doc.Remove(name)
		doc.Append(name, value)
	}
	return fields{doc: doc}
}

type fieldsProjection struct {
	projections []fields
}

func newFieldsProjection(projections []fields) fieldsProjection {
	return fieldsProjection{
		projections: projections,
	}
}

func (p fieldsProjection) Pro() *bsonx.BsonDocument {
	combinedDocument := bsonx.BsonEmpty()
	for _, sort := range p.projections {
		sortDocument := sort.Pro()
		for _, key := range sortDocument.Keys() {
			combinedDocument.Remove(key)
			combinedDocument.Append(key, sortDocument.GetValue(key))
		}
	}
	return combinedDocument
}

type elemMatchFilterProjection struct {
	fieldName string
	filter    bsonx.Bson
}

func newElemMatchFilterProjection(fieldName string, filter bsonx.Bson) elemMatchFilterProjection {
	return elemMatchFilterProjection{
		filter:    filter,
		fieldName: fieldName,
	}
}

func (p elemMatchFilterProjection) Pro() *bsonx.BsonDocument {
	return bsonx.BsonDoc(p.fieldName, bsonx.BsonDoc("$elemMatch", p.filter.Pro()))
}

func (p elemMatchFilterProjection) Document() bson.D {
	return p.Pro().Document()
}
