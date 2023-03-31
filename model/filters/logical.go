package filters

import (
	"strings"

	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"go.mongodb.org/mongo-driver/bson"
)

type logicalFilter struct {
	filter bsonx.Bson
}

func (f logicalFilter) Exp() bsonx.IBsonValue {
	return f.filter.ToBsonDocument()
}

// And Joins query clauses with a logical AND returns all documents that match the
// conditions of both clauses.
func And[T FilterExpression](filters ...T) logicalFilter {
	return logicalFilter{filter: NewAddFilter(filters)}
}

// Or Inverts the effect of a query expression and returns documents that do not
// match the query expression.
func Or[T FilterExpression](filters ...T) logicalFilter {
	return logicalFilter{filter: NewOrNorFilter(OR, filters)}
}

// Nor Joins query clauses with a logical NOR returns all documents that fail to
// match both clauses.
func Nor[T FilterExpression](filters ...T) logicalFilter {
	return logicalFilter{filter: NewOrNorFilter(NOR, filters)}
}

// Not Joins query clauses with a logical OR returns all documents that match the
// conditions of either clause.
func Not[T FilterExpression](f T) logicalFilter {
	return logicalFilter{filter: NewNotFilter(f)}
}

type addFilter[T expression.Expression] struct {
	filters []T
}

func NewAddFilter[T expression.Expression](filters []T) addFilter[T] {
	return addFilter[T]{
		filters: filters,
	}
}

func (s addFilter[T]) ToBsonDocument() *bsonx.BsonDocument {
	clauses := bsonx.Array()
	for _, filter := range s.filters {
		clauses.Append(filter.Exp())
	}
	return bsonx.BsonDoc("$and", clauses)
}

func (s addFilter[T]) Document() bson.D {
	return s.ToBsonDocument().Document()
}

type orNorFilter[T expression.Expression] struct {
	operator Operator
	filters  []T
}

func NewOrNorFilter[T expression.Expression](operator Operator, filters []T) orNorFilter[T] {
	return orNorFilter[T]{
		operator: operator,
		filters:  filters,
	}
}

func (s orNorFilter[T]) ToBsonDocument() *bsonx.BsonDocument {
	filtersArray := bsonx.Array()
	for _, filter := range s.filters {
		filtersArray.Append(filter.Exp())
	}
	return bsonx.BsonDoc(s.operator.name, filtersArray)
}

func (s orNorFilter[T]) Document() bson.D {
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

type notFilter[T expression.Expression] struct {
	filter T
}

func NewNotFilter[T expression.Expression](value T) notFilter[T] {
	return notFilter[T]{
		filter: value,
	}
}

func (f notFilter[T]) ToBsonDocument() *bsonx.BsonDocument {
	filter := f.filter.Exp()
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

	return f.createFilter("$and", f.filter.Exp())
}

func (f notFilter[T]) Document() bson.D {
	return f.ToBsonDocument().Document()
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
