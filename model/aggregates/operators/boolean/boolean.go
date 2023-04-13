// Package boolean Boolean expressions evaluate their argument expressions as booleans and return a boolean as the result.
// In addition to the false boolean value, Boolean expression evaluates as false the following: null, 0, and undefined values.
// The Boolean expression evaluates all other values as true, including non-zero numeric values and arrays.
package boolean

import (
	"strings"

	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"go.mongodb.org/mongo-driver/bson"
)

type booleanOperator struct {
	doc bsonx.Bson
}

func (o booleanOperator) Exp() bsonx.IBsonValue {
	return o.doc.BsonDocument()
}

// And Evaluates one or more expressions and returns true if all of the
// expressions are true or if run with no argument expressions. Otherwise, $and
// returns false.
func And[T expression.AnyExpression](filters ...T) booleanOperator {
	return booleanOperator{doc: NewAddFilter(filters)}
}

// Or Evaluates one or more expressions and returns true if any of the
// expressions are true. Otherwise, $or returns false. $or has the following
// syntax: { $or: [ <expression1>, <expression2>, ... ] }
func Or[T expression.AnyExpression](filters ...T) booleanOperator {
	return booleanOperator{doc: NewOrNorFilter(OR, filters)}
}

// Nor Evaluates one or more expressions and returns true if any of the
// expressions are true. Otherwise, $or returns false. $or has the following
// syntax: { $or: [ <expression1>, <expression2>, ... ] }
func Nor[T expression.AnyExpression](filters ...T) booleanOperator {
	return booleanOperator{doc: NewOrNorFilter(NOR, filters)}
}

// Not Evaluates a boolean and returns the opposite boolean value; i.e. when
// passed an expression that evaluates to true, $not returns false; when passed
// an expression that evaluates to false, $not returns true. $not has the
// following syntax: { $not: [ <expression> ] }
func Not[T expression.AnyExpression](f T) booleanOperator {
	return booleanOperator{doc: NewNotFilter(f)}
}

type addFilter[T expression.AnyExpression] struct {
	filters []T
}

func NewAddFilter[T expression.AnyExpression](filters []T) addFilter[T] {
	return addFilter[T]{
		filters: filters,
	}
}

func (s addFilter[T]) BsonDocument() *bsonx.BsonDocument {
	clauses := bsonx.Array()
	for _, filter := range s.filters {
		clauses.Append(filter)
	}
	return bsonx.BsonDoc("$and", clauses)
}

func (s addFilter[T]) Document() bson.D {
	return s.BsonDocument().Document()
}

type orNorFilter[T expression.AnyExpression] struct {
	operator Operator
	filters  []T
}

func NewOrNorFilter[T expression.AnyExpression](operator Operator, filters []T) orNorFilter[T] {
	return orNorFilter[T]{
		operator: operator,
		filters:  filters,
	}
}

func (s orNorFilter[T]) BsonDocument() *bsonx.BsonDocument {
	filtersArray := bsonx.Array()
	for _, filter := range s.filters {
		filtersArray.Append(filter)
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

type notFilter[T expression.AnyExpression] struct {
	filter T
}

func NewNotFilter[T expression.AnyExpression](value T) notFilter[T] {
	return notFilter[T]{
		filter: value,
	}
}

func (f notFilter[T]) BsonDocument() *bsonx.BsonDocument {
	if f.filter.IsDocument() {
		filterDocument := f.filter.AsDocument()

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

	return f.createFilter("$and", f.filter)
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
