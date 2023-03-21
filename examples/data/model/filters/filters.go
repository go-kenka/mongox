package filters

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/expressions"
	"github.com/go-kenka/mongox/examples/data/model/geojson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

func Eq(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewSimpleFilter(fieldName, value)
}
func Ne(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewOperatorFilter("$ne", fieldName, value)
}
func Gt(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewOperatorFilter("$gt", fieldName, value)
}
func Lt(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewOperatorFilter("$lt", fieldName, value)
}
func Gte(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewOperatorFilter("$gte", fieldName, value)
}
func Lte(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewOperatorFilter("$lte", fieldName, value)
}
func In(fieldName string, values ...expressions.TItem) bsonx.Bson {
	return NewIterableOperatorFilter(fieldName, "$in", values)
}
func Nin(fieldName string, values ...expressions.TItem) bsonx.Bson {
	return NewIterableOperatorFilter(fieldName, "$nin", values)
}
func And(filters ...bsonx.Bson) bsonx.Bson {
	return NewAddFilter(filters)
}
func Or(filters ...bsonx.Bson) bsonx.Bson {
	return NewOrNorFilter(OR, filters)
}
func Nor(filters ...bsonx.Bson) bsonx.Bson {
	return NewOrNorFilter(NOR, filters)
}
func Not(filter bsonx.Bson) bsonx.Bson {
	return NewNotFilter(filter)
}
func Exists(fieldName string, value bool) bsonx.Bson {
	return NewOperatorFilter("$exists", fieldName, bsonx.NewBsonBoolean(value))
}
func Type(fieldName string, value bsonx.BsonType) bsonx.Bson {
	return NewOperatorFilter("$type", fieldName, bsonx.NewBsonInt32(value.Value()))
}
func Mod(fieldName string, divisor, remainder int64) bsonx.Bson {
	return NewOperatorFilter("$mod", fieldName, bsonx.NewBsonArray(bsonx.NewBsonInt64(divisor), bsonx.NewBsonInt64(remainder)))
}
func Regex(fieldName string, pattern, options string) bsonx.Bson {
	return NewSimpleFilter(fieldName, bsonx.NewBsonRegularExpression(primitive.Regex{
		Pattern: pattern,
		Options: options,
	}))
}
func Text(search string, textSearchOptions TextSearchOptions) bsonx.Bson {
	return NewTextFilter(search, textSearchOptions)
}
func Where(javaScriptExpression string) bsonx.Bson {
	return bsonx.NewBsonDocument("$where", bsonx.NewBsonString(javaScriptExpression))
}
func Expr(expression expressions.TExpression) bsonx.Bson {
	return NewSimpleEncodingFilter("$expr", expression)
}
func All(fieldName string, values ...expressions.TItem) bsonx.Bson {
	return NewIterableOperatorFilter(fieldName, "$all", values)
}
func ElemMatch(fieldName string, filter bsonx.Bson) bsonx.Bson {
	return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonDocument("$elemMatch", filter.ToBsonDocument()))
}
func Size(fieldName string, size int32) bsonx.Bson {
	return NewOperatorFilter("$size", fieldName, bsonx.NewBsonInt32(size))
}
func BitsAllClear(fieldName string, bitmask int64) bsonx.Bson {
	return NewOperatorFilter("$bitsAllClear", fieldName, bsonx.NewBsonInt64(bitmask))
}
func BitsAllSet(fieldName string, bitmask int64) bsonx.Bson {
	return NewOperatorFilter("$bitsAllSet", fieldName, bsonx.NewBsonInt64(bitmask))
}
func BitsAnyClear(fieldName string, bitmask int64) bsonx.Bson {
	return NewOperatorFilter("$bitsAnyClear", fieldName, bsonx.NewBsonInt64(bitmask))
}
func BitsAnySet(fieldName string, bitmask int64) bsonx.Bson {
	return NewOperatorFilter("$bitsAnySet", fieldName, bsonx.NewBsonInt64(bitmask))
}
func GeoWithin(fieldName string, geometry geojson.IGeometry) bsonx.Bson {
	return NewGeometryOperatorFilter("$geoWithin", fieldName, geometry.Encode(), 0, 0)
}
func GeoWithinBson(fieldName string, geometry bsonx.Bson) bsonx.Bson {
	return NewGeometryOperatorFilter("$geoWithin", fieldName, geometry.ToBsonDocument(), 0, 0)
}
func GeoWithinBox(fieldName string, lowerLeftX, lowerLeftY, upperRightX, upperRightY float64) bsonx.Bson {
	box := bsonx.NewBsonDocument("$box", bsonx.NewBsonArray(
		bsonx.NewBsonArray(
			bsonx.NewBsonDouble(lowerLeftX),
			bsonx.NewBsonDouble(lowerLeftY),
		),
		bsonx.NewBsonArray(
			bsonx.NewBsonDouble(upperRightX),
			bsonx.NewBsonDouble(upperRightY),
		)))

	return NewGeometryOperatorFilter("$geoWithin", fieldName, box, 0, 0)
}
func GeoWithinPolygon(fieldName string, points [][]float64) bsonx.Bson {
	pointsArray := bsonx.NewBsonArray()
	for _, point := range points {
		pointsArray.Append(bsonx.NewBsonArray(bsonx.NewBsonDouble(point[0]), bsonx.NewBsonDouble(point[1])))
	}
	polygon := bsonx.NewBsonDocument("$polygon", pointsArray)
	return NewOperatorFilter("$geoWithin", fieldName, polygon)
}

func GeoWithinCenter(fieldName string, x, y, radius float64) bsonx.Bson {
	center := bsonx.NewBsonDocument("$center", bsonx.NewBsonArray(
		bsonx.NewBsonArray(bsonx.NewBsonDouble(x), bsonx.NewBsonDouble(y)),
		bsonx.NewBsonDouble(radius),
	))
	return NewOperatorFilter("$geoWithin", fieldName, center)
}

func GeoWithinCenterSphere(fieldName string, x, y, radius float64) bsonx.Bson {
	center := bsonx.NewBsonDocument("$centerSphere", bsonx.NewBsonArray(
		bsonx.NewBsonArray(bsonx.NewBsonDouble(x), bsonx.NewBsonDouble(y)),
		bsonx.NewBsonDouble(radius),
	))
	return NewOperatorFilter("$geoWithin", fieldName, center)
}
func GeoIntersects(fieldName string, geometry geojson.IGeometry) bsonx.Bson {
	return NewGeometryOperatorFilter("$geoIntersects", fieldName, geometry.Encode(), 0, 0)
}
func GeoIntersectsBson(fieldName string, geometry bsonx.Bson) bsonx.Bson {
	return NewGeometryOperatorFilter("$geoIntersects", fieldName, geometry.ToBsonDocument(), 0, 0)
}
func Near(fieldName string, geometry geojson.Point, maxDistance, minDistance float64) bsonx.Bson {
	return NewGeometryOperatorFilter("$near", fieldName, geometry.Encode(), maxDistance, minDistance)
}
func NearBson(fieldName string, geometry bsonx.Bson, maxDistance, minDistance float64) bsonx.Bson {
	return NewGeometryOperatorFilter("$near", fieldName, geometry.ToBsonDocument(), maxDistance, minDistance)
}
func NearCoordinate(fieldName string, x, y float64, maxDistance, minDistance float64) bsonx.Bson {
	return createNearFilterDocument(fieldName, x, y, maxDistance, minDistance, "$near")
}
func NearSphere(fieldName string, geometry geojson.Point, maxDistance, minDistance float64) bsonx.Bson {
	return NewGeometryOperatorFilter("$nearSphere", fieldName, geometry.Encode(), maxDistance, minDistance)
}
func NearSphereBson(fieldName string, geometry bsonx.Bson, maxDistance, minDistance float64) bsonx.Bson {
	return NewGeometryOperatorFilter("$nearSphere", fieldName, geometry.ToBsonDocument(), maxDistance, minDistance)
}
func NearSphereCoordinate(fieldName string, x, y float64, maxDistance, minDistance float64) bsonx.Bson {
	return createNearFilterDocument(fieldName, x, y, maxDistance, minDistance, "$nearSphere")
}
func JsonSchema(schema bsonx.Bson) bsonx.Bson {
	return NewSimpleEncodingFilter("$jsonSchema", schema.ToBsonDocument())
}
func Empty() bsonx.Bson {
	return bsonx.NewEmptyDoc()
}
func createNearFilterDocument(fieldName string, x, y float64, maxDistance, minDistance float64, operator string) bsonx.Bson {
	nearFilter := bsonx.NewBsonDocument(operator, bsonx.NewBsonArray(bsonx.NewBsonDouble(x), bsonx.NewBsonDouble(y)))
	if maxDistance > 0 {
		nearFilter.Append("$maxDistance", bsonx.NewBsonDouble(maxDistance))
	}
	if minDistance > 0 {
		nearFilter.Append("$minDistance", bsonx.NewBsonDouble(minDistance))
	}
	return bsonx.NewBsonDocument(fieldName, nearFilter)
}

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

func (s SimpleFilter) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewBsonDocument(s.fieldName, s.value)
}

func (s SimpleFilter) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

type OperatorFilter[T expressions.TItem] struct {
	operatorName string
	fieldName    string
	value        T
}

func NewOperatorFilter[T expressions.TItem](operatorName string, fieldName string, value T) OperatorFilter[T] {
	return OperatorFilter[T]{
		operatorName: operatorName,
		fieldName:    fieldName,
		value:        value,
	}
}

func (s OperatorFilter[T]) ToBsonDocument() bsonx.BsonDocument {
	doc := bsonx.NewEmptyDoc()
	operator := bsonx.NewBsonDocument(s.operatorName, s.value)
	doc.Append(s.fieldName, operator)
	return doc
}
func (s OperatorFilter[T]) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

type AddFilter struct {
	filters []bsonx.Bson
}

func NewAddFilter(filters []bsonx.Bson) AddFilter {
	return AddFilter{
		filters: filters,
	}
}

func (s AddFilter) ToBsonDocument() bsonx.BsonDocument {
	clauses := bsonx.NewBsonArray()
	for _, filter := range s.filters {
		clauses.Append(filter.ToBsonDocument())
	}
	return bsonx.NewBsonDocument("$and", clauses)
}

func (s AddFilter) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

type OrNorFilter struct {
	operator Operator
	filters  []bsonx.Bson
}

func NewOrNorFilter(operator Operator, filters []bsonx.Bson) OrNorFilter {
	return OrNorFilter{
		operator: operator,
		filters:  filters,
	}
}

func (s OrNorFilter) ToBsonDocument() bsonx.BsonDocument {
	filtersArray := bsonx.NewBsonArray()
	for _, filter := range s.filters {
		filtersArray.Append(filter.ToBsonDocument())
	}
	return bsonx.NewBsonDocument(s.operator.name, filtersArray)
}

func (s OrNorFilter) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

type Operator struct {
	name         string
	toStringName string
}

func NewOperator(name string, toStringName string) Operator {
	return Operator{
		name:         name,
		toStringName: toStringName,
	}
}

var (
	OR  = NewOperator("$or", "Or")
	NOR = NewOperator("$nor", "Nor")
)

type IterableOperatorFilter[T expressions.TItem] struct {
	operatorName string
	fieldName    string
	values       []T
}

func NewIterableOperatorFilter[T expressions.TItem](fieldName string, operatorName string, values []T) IterableOperatorFilter[T] {
	return IterableOperatorFilter[T]{
		operatorName: operatorName,
		fieldName:    fieldName,
		values:       values,
	}
}

func (s IterableOperatorFilter[T]) ToBsonDocument() bsonx.BsonDocument {
	doc := bsonx.NewEmptyDoc()
	values := bsonx.NewBsonArray()
	for _, value := range s.values {
		values.Append(value)
	}
	operator := bsonx.NewBsonDocument(s.operatorName, values)
	doc.Append(s.fieldName, operator)
	return doc
}

func (s IterableOperatorFilter[T]) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

type SimpleEncodingFilter[T expressions.TItem] struct {
	fieldName string
	value     T
}

func NewSimpleEncodingFilter[T expressions.TItem](fieldName string, value T) SimpleEncodingFilter[T] {
	return SimpleEncodingFilter[T]{
		fieldName: fieldName,
		value:     value,
	}
}

func (s SimpleEncodingFilter[T]) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewBsonDocument(s.fieldName, s.value)
}

func (s SimpleEncodingFilter[T]) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

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

type NotFilter struct {
	filter bsonx.Bson
}

func NewNotFilter(value bsonx.Bson) NotFilter {
	return NotFilter{
		filter: value,
	}
}

func (f NotFilter) ToBsonDocument() bsonx.BsonDocument {
	filterDocument := f.filter.ToBsonDocument()

	if filterDocument.Size() == 1 {
		keys := filterDocument.Keys()
		return f.createFilter(keys[0], filterDocument.GetValue(keys[0]))
	}

	values := bsonx.NewBsonArray()
	for k, v := range filterDocument.Data() {
		values.Append(bsonx.NewBsonDocument(k, v))
	}

	return f.createFilter("$and", values)

}

func (f NotFilter) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

func (f NotFilter) createFilter(fieldName string, value bsonx.IBsonValue) bsonx.BsonDocument {
	if strings.HasPrefix(fieldName, "$") {
		return bsonx.NewBsonDocument("$not", bsonx.NewBsonDocument(fieldName, value))
	}
	if (value.IsDocument() && f.containsOperator(value.AsDocument())) || value.IsRegularExpression() {
		return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonDocument("$not", value))
	}
	return bsonx.NewBsonDocument(fieldName, bsonx.NewBsonDocument("$not", bsonx.NewBsonDocument("$eq", value)))
}

func (f NotFilter) containsOperator(value bsonx.BsonDocument) bool {
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

type GeometryOperatorFilter[T expressions.TItem] struct {
	fieldName    string
	operatorName string
	geometry     T
	maxDistance  float64
	minDistance  float64
}

func NewGeometryOperatorFilter[T expressions.TItem](
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

func (s GeometryOperatorFilter[T]) ToBsonDocument() bsonx.BsonDocument {
	operator := bsonx.NewEmptyDoc()
	geometry := bsonx.NewEmptyDoc()
	geometry.Append("$geometry", s.geometry)
	if s.maxDistance > 0 {
		geometry.Append("$maxDistance", bsonx.NewBsonDouble(s.maxDistance))
	}
	if s.minDistance > 0 {
		geometry.Append("$minDistance", bsonx.NewBsonDouble(s.minDistance))
	}
	operator.Append(s.operatorName, geometry)
	return bsonx.NewBsonDocument(s.fieldName, operator)
}

func (s GeometryOperatorFilter[T]) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

type TextFilter struct {
	search            string
	textSearchOptions TextSearchOptions
}

func NewTextFilter(
	search string,
	textSearchOptions TextSearchOptions,
) TextFilter {
	return TextFilter{
		search:            search,
		textSearchOptions: textSearchOptions,
	}
}

func (s TextFilter) ToBsonDocument() bsonx.BsonDocument {
	searchDocument := bsonx.NewBsonDocument("$search", bsonx.NewBsonString(s.search))
	if s.textSearchOptions.HasLanguage() {
		searchDocument.Append("$language", bsonx.NewBsonString(s.textSearchOptions.GetLanguage()))
	}
	if s.textSearchOptions.HasCaseSensitive() {
		searchDocument.Append("$caseSensitive", bsonx.NewBsonBoolean(s.textSearchOptions.GetCaseSensitive()))
	}
	if s.textSearchOptions.HasDiacriticSensitive() {
		searchDocument.Append("$diacriticSensitive", bsonx.NewBsonBoolean(s.textSearchOptions.GetDiacriticSensitive()))
	}
	return bsonx.NewBsonDocument("$text", searchDocument)
}
func (s TextFilter) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}
