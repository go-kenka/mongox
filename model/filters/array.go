package filters

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"go.mongodb.org/mongo-driver/bson"
)

type arrayFilter struct {
	filter bsonx.Bson
}

func (f arrayFilter) Value() bsonx.IBsonValue {
	return f.filter.ToBsonDocument()
}

func (f arrayFilter) Document() bson.D {
	return f.filter.Document()
}

// All The $all operator selects the documents where the value of a field is an
// array that contains all the specified elements. To specify an $all
// expression, use the following prototype: { <field>: { $all: [ <value1> ,
// <value2> ... ] } }
func All[I expression.AnyExpression](fieldName string, values ...I) arrayFilter {
	return arrayFilter{filter: newIterableOperatorFilter(fieldName, "$all", values)}
}

// ElemMatch The $elemMatch operator matches documents that contain an array field with at
// least one element that matches all the specified query criteria. { <field>: {
// $elemMatch: { <query1>, <query2>, ... } } } If you specify only a single
// <query> condition in the $elemMatch expression, and are not using the $not or
// $ne operators inside of $elemMatch, $elemMatch can be omitted.
func ElemMatch(fieldName string, filter bsonx.Bson) arrayFilter {
	return arrayFilter{filter: bsonx.BsonDoc(fieldName, bsonx.BsonDoc("$elemMatch", filter.ToBsonDocument()))}
}

// Size The $size operator matches any array with the number of elements
// specified by the argument. For example: db.collection.find( { field: { $size:
// 2 } } ); returns all documents in collection where field is an array with 2
// elements. For instance, the above expression will return { field: [ red,
// green ] } and { field: [ apple, lime ] } but not { field: fruit } or { field:
// [ orange, lemon, grapefruit ] }. To match fields with only one element within
// an array use $size with a value of 1, as follows: db.collection.find( {
// field: { $size: 1 } } ); $size does not accept ranges of values. To select
// documents based on fields with different numbers of elements, create a
// counter field that you increment when you add elements to a field. Queries
// cannot use indexes for the $size portion of a query, although the other
// portions of a query can use indexes if applicable.
func Size(fieldName string, size int32) arrayFilter {
	return arrayFilter{filter: newOperatorFilter("$size", fieldName, bsonx.Int32(size))}
}
