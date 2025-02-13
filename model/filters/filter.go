package filters

import (
	"strings"

	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/internal/options"
	"go.mongodb.org/mongo-driver/bson"
)

type Filter interface {
	Value() bsonx.IBsonValue
	Document() bson.D
}

type MatchFilter interface {
	Filter
	logicalFilter | comparisonFilter | arrayFilter | bitwiseFilter | elementFilter | evaluationFilter | emptyFilter
}

type addFilter[T Filter] struct {
	filters []T
}

func newAddFilter[T Filter](filters []T) addFilter[T] {
	return addFilter[T]{
		filters: filters,
	}
}

func (s addFilter[T]) BsonDocument() *bsonx.BsonDocument {
	clauses := bsonx.Array()
	for _, filter := range s.filters {
		clauses.Append(filter.Value())
	}
	return bsonx.BsonDoc("$and", clauses)
}

func (s addFilter[T]) Document() bson.D {
	return s.BsonDocument().Document()
}

type orNorFilter[T Filter] struct {
	operator Operator
	filters  []T
}

func newOrNorFilter[T Filter](operator Operator, filters []T) orNorFilter[T] {
	return orNorFilter[T]{
		operator: operator,
		filters:  filters,
	}
}

func (s orNorFilter[T]) BsonDocument() *bsonx.BsonDocument {
	filtersArray := bsonx.Array()
	for _, filter := range s.filters {
		filtersArray.Append(filter.Value())
	}
	return bsonx.BsonDoc(s.operator.name, filtersArray)
}

func (s orNorFilter[T]) Document() bson.D {
	return s.BsonDocument().Document()
}

type Operator struct {
	name         string
	toStringName string
}

func newOperator(name string, toStringName string) Operator {
	return Operator{
		name:         name,
		toStringName: toStringName,
	}
}

var (
	OR  = newOperator("$or", "Or")
	NOR = newOperator("$nor", "Nor")
)

var (
	DBREFKeys = []string{
		"$ref",
		"$id",
	}
	DBREFKeysWithDb = []string{
		"$ref",
		"$id",
		"$db",
	}
)

type notFilter[T Filter] struct {
	filter T
}

func newNotFilter[T Filter](value T) notFilter[T] {
	return notFilter[T]{
		filter: value,
	}
}

func (f notFilter[T]) BsonDocument() *bsonx.BsonDocument {
	filter := f.filter.Value()
	if filter.IsDocument() {
		filterDocument := filter.AsDocument()

		if filterDocument.Size() == 1 {
			v := filterDocument.Data()
			key := v[0].Key
			return f.createFilter(key, filterDocument.GetValue(key))
		}

		values := bsonx.Array()
		for _, v := range filterDocument.Data() {
			values.Append(bsonx.BsonDoc(v.Key, v.Value))
		}
		return f.createFilter("$and", values)
	}

	return f.createFilter("$and", f.filter.Value())
}

func (f notFilter[T]) Document() bson.D {
	return f.BsonDocument().Document()
}

func (f notFilter[T]) createFilter(fieldName string, value bsonx.IBsonValue) *bsonx.BsonDocument {
	if strings.HasPrefix(fieldName, "$") {
		return bsonx.BsonDoc("$not", bsonx.BsonDoc(fieldName, value))
	}
	if (value.IsDocument() && f.containsOperator(value.AsDocument())) || value.IsRegularExpression() {
		return bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$not", value))
	}
	return bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$not", bsonx.BsonDoc("$eq", value)))
}

func (f notFilter[T]) containsOperator(value *bsonx.BsonDocument) bool {
	keys := value.Keys()
	if equals(keys, DBREFKeys) || equals(keys, DBREFKeysWithDb) {
		return false
	}
	for _, key := range keys {
		if strings.HasPrefix(key, "$") {
			return true
		}
	}
	return false
}

func equals(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type simpleFilter struct {
	fieldName string
	value     bsonx.IBsonValue
}

func newSimpleFilter(fieldName string, value bsonx.IBsonValue) simpleFilter {
	return simpleFilter{
		fieldName: fieldName,
		value:     value,
	}
}

func (s simpleFilter) BsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(s.fieldName, s.value)
}

func (s simpleFilter) Document() bson.D {
	return s.BsonDocument().Document()
}

type operatorFilter[T expression.AnyExpression] struct {
	operatorName string
	fieldName    string
	value        T
}

func newOperatorFilter[T expression.AnyExpression](operatorName string, fieldName string, value T) operatorFilter[T] {
	return operatorFilter[T]{
		operatorName: operatorName,
		fieldName:    fieldName,
		value:        value,
	}
}

func (s operatorFilter[T]) BsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	operator := bsonx.BsonDoc(s.operatorName, s.value)
	doc.Append(s.fieldName, operator)
	return doc
}
func (s operatorFilter[T]) Document() bson.D {
	return s.BsonDocument().Document()
}

type iterableOperatorFilter[T expression.AnyExpression] struct {
	operatorName string
	fieldName    string
	values       []T
}

func newIterableOperatorFilter[T expression.AnyExpression](fieldName string, operatorName string, values []T) iterableOperatorFilter[T] {
	return iterableOperatorFilter[T]{
		operatorName: operatorName,
		fieldName:    fieldName,
		values:       values,
	}
}

func (s iterableOperatorFilter[T]) BsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	values := bsonx.Array()
	for _, value := range s.values {
		values.Append(value)
	}
	operator := bsonx.BsonDoc(s.operatorName, values)
	doc.Append(s.fieldName, operator)
	return doc
}

func (s iterableOperatorFilter[T]) Document() bson.D {
	return s.BsonDocument().Document()
}

type simpleEncodingFilter[T expression.AnyExpression] struct {
	fieldName string
	value     T
}

func newSimpleEncodingFilter[T expression.AnyExpression](fieldName string, value T) simpleEncodingFilter[T] {
	return simpleEncodingFilter[T]{
		fieldName: fieldName,
		value:     value,
	}
}

func (s simpleEncodingFilter[T]) BsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(s.fieldName, s.value)
}

func (s simpleEncodingFilter[T]) Document() bson.D {
	return s.BsonDocument().Document()
}

type geometryOperatorFilter[T expression.AnyExpression] struct {
	fieldName    string
	operatorName string
	geometry     T
	maxDistance  float64
	minDistance  float64
}

func newGeometryOperatorFilter[T expression.AnyExpression](
	fieldName string,
	operatorName string,
	geometry T,
	maxDistance float64,
	minDistance float64,
) geometryOperatorFilter[T] {
	return geometryOperatorFilter[T]{
		fieldName:    fieldName,
		operatorName: operatorName,
		geometry:     geometry,
		maxDistance:  maxDistance,
		minDistance:  minDistance,
	}
}

func (s geometryOperatorFilter[T]) BsonDocument() *bsonx.BsonDocument {
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

func (s geometryOperatorFilter[T]) Document() bson.D {
	return s.BsonDocument().Document()
}

type textFilter struct {
	search            string
	textSearchOptions options.TextSearchOptions
}

func newTextFilter(
	search string,
	textSearchOptions options.TextSearchOptions,
) textFilter {
	return textFilter{
		search:            search,
		textSearchOptions: textSearchOptions,
	}
}

func (s textFilter) BsonDocument() *bsonx.BsonDocument {
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
func (s textFilter) Document() bson.D {
	return s.BsonDocument().Document()
}
