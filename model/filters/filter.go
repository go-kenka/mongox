package filters

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/filter"
	"go.mongodb.org/mongo-driver/bson"
)

type Filter bsonx.Bson

type SimpleFilter struct {
	fieldName string
	value     bsonx.IBsonValue
}

func NewSimpleFilter(fieldName string, value bsonx.IBsonValue) SimpleFilter {
	return SimpleFilter{
		fieldName: fieldName,
		value:     value,
	}
}

func (s SimpleFilter) ToBsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(s.fieldName, s.value)
}

func (s SimpleFilter) Document() bson.D {
	return s.ToBsonDocument().Document()
}

type OperatorFilter[T bsonx.Expression] struct {
	operatorName string
	fieldName    string
	value        T
}

func NewOperatorFilter[T bsonx.Expression](operatorName string, fieldName string, value T) OperatorFilter[T] {
	return OperatorFilter[T]{
		operatorName: operatorName,
		fieldName:    fieldName,
		value:        value,
	}
}

func (s OperatorFilter[T]) ToBsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	operator := bsonx.BsonDoc(s.operatorName, s.value)
	doc.Append(s.fieldName, operator)
	return doc
}
func (s OperatorFilter[T]) Document() bson.D {
	return s.ToBsonDocument().Document()
}

type IterableOperatorFilter[T bsonx.Expression] struct {
	operatorName string
	fieldName    string
	values       []T
}

func NewIterableOperatorFilter[T bsonx.Expression](fieldName string, operatorName string, values []T) IterableOperatorFilter[T] {
	return IterableOperatorFilter[T]{
		operatorName: operatorName,
		fieldName:    fieldName,
		values:       values,
	}
}

func (s IterableOperatorFilter[T]) ToBsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	values := bsonx.Array()
	for _, value := range s.values {
		values.Append(value)
	}
	operator := bsonx.BsonDoc(s.operatorName, values)
	doc.Append(s.fieldName, operator)
	return doc
}

func (s IterableOperatorFilter[T]) Document() bson.D {
	return s.ToBsonDocument().Document()
}

type SimpleEncodingFilter[T bsonx.Expression] struct {
	fieldName string
	value     T
}

func NewSimpleEncodingFilter[T bsonx.Expression](fieldName string, value T) SimpleEncodingFilter[T] {
	return SimpleEncodingFilter[T]{
		fieldName: fieldName,
		value:     value,
	}
}

func (s SimpleEncodingFilter[T]) ToBsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(s.fieldName, s.value)
}

func (s SimpleEncodingFilter[T]) Document() bson.D {
	return s.ToBsonDocument().Document()
}

type GeometryOperatorFilter[T bsonx.Expression] struct {
	fieldName    string
	operatorName string
	geometry     T
	maxDistance  float64
	minDistance  float64
}

func NewGeometryOperatorFilter[T bsonx.Expression](
	fieldName string,
	operatorName string,
	geometry T,
	maxDistance float64,
	minDistance float64,
) GeometryOperatorFilter[T] {
	return GeometryOperatorFilter[T]{
		fieldName:    fieldName,
		operatorName: operatorName,
		geometry:     geometry,
		maxDistance:  maxDistance,
		minDistance:  minDistance,
	}
}

func (s GeometryOperatorFilter[T]) ToBsonDocument() *bsonx.BsonDocument {
	operator := bsonx.BsonEmpty()
	geometry := bsonx.BsonEmpty()
	geometry.Append("$geometry", s.geometry)
	if s.maxDistance > 0 {
		geometry.Append("$maxDistance", bsonx.Double(s.maxDistance))
	}
	if s.minDistance > 0 {
		geometry.Append("$minDistance", bsonx.Double(s.minDistance))
	}
	operator.Append(s.operatorName, geometry)
	return bsonx.BsonDoc(s.fieldName, operator)
}

func (s GeometryOperatorFilter[T]) Document() bson.D {
	return s.ToBsonDocument().Document()
}

type TextFilter struct {
	search            string
	textSearchOptions filter.TextSearchOptions
}

func NewTextFilter(
	search string,
	textSearchOptions filter.TextSearchOptions,
) TextFilter {
	return TextFilter{
		search:            search,
		textSearchOptions: textSearchOptions,
	}
}

func (s TextFilter) ToBsonDocument() *bsonx.BsonDocument {
	searchDocument := bsonx.BsonDoc("$search", bsonx.String(s.search))
	if s.textSearchOptions.HasLanguage() {
		searchDocument.Append("$language", bsonx.String(s.textSearchOptions.GetLanguage()))
	}
	if s.textSearchOptions.HasCaseSensitive() {
		searchDocument.Append("$caseSensitive", bsonx.Boolean(s.textSearchOptions.GetCaseSensitive()))
	}
	if s.textSearchOptions.HasDiacriticSensitive() {
		searchDocument.Append("$diacriticSensitive", bsonx.Boolean(s.textSearchOptions.GetDiacriticSensitive()))
	}
	return bsonx.BsonDoc("$text", searchDocument)
}
func (s TextFilter) Document() bson.D {
	return s.ToBsonDocument().Document()
}
