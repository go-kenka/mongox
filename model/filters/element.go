package filters

import (
	"github.com/go-kenka/mongox/bsonx"
)

type elementFilter struct {
	filter bsonx.Bson
}

func (f elementFilter) Value() bsonx.IBsonValue {
	return f.filter.ToBsonDocument()
}

// Exists When <boolean> is true, $exists matches the documents that contain the
// field, including documents where the field value is null. If <boolean> is
// false, the query returns only the documents that do not contain the field.
// [1] MongoDB $exists does not correspond to SQL operator exists. For SQL
// exists, refer to the $in operator.
func Exists(fieldName string, value bool) elementFilter {
	return elementFilter{filter: newOperatorFilter("$exists", fieldName, bsonx.Boolean(value))}
}

// Type selects documents where the value of the field is an instance of the
// specified BSON type(s). Querying by data type is useful when dealing with
// highly unstructured data where data types are not predictable. A $type
// expression for a single BSON type has the following syntax: { field: { $type:
// <BSON type> } } You can specify either the number or alias for the BSON type
// The $type expression can also accept an array of BSON types and has the
// following syntax: { field: { $type: [ <BSON type1> , <BSON type2>, ... ] } }
// The above query will match documents where the field value is any of the
// listed types. The types specified in the array can be either numeric or
// string aliases.
func Type(fieldName string, value bsonx.BsonType) elementFilter {
	return elementFilter{filter: newOperatorFilter("$type", fieldName, bsonx.Int32(value.Value()))}
}
