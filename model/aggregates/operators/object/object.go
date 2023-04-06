// Package object Object Expression Operators
package object

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type objectOperator struct {
	doc bsonx.Bson
}

func (o objectOperator) Exp() bsonx.IBsonValue {
	return o.doc.Pro()
}

// MergeObjects Combines multiple documents into a single document.
// $mergeObjects is available in these stages:
// $bucket
// $bucketAuto
// $group
// When used as a $bucket, $bucketAuto, or $group stage accumulator,
// $mergeObjects has this syntax:
// { $mergeObjects: <document> }
// When used in other expressions (including in $bucket, $bucketAuto, and $group stages) but not as an accumulator,
// $mergeObjects has this syntax:
// { $mergeObjects: [ <document1>, <document2>, ... ] }
// The <document> can be any valid expression that resolves to a document.
func MergeObjects[T expression.DocumentExpression](e []T) objectOperator {
	if len(e) == 1 {
		return objectOperator{doc: bsonx.BsonDoc("$mergeObjects", e[0])}
	}
	doc := bsonx.Array()
	for _, t := range e {
		doc.Append(t)
	}
	return objectOperator{doc: bsonx.BsonDoc("$mergeObjects", doc)}
}

// ObjectToArray Converts a document to an array. The return array contains an element for each field/value pair in the original document. Each element in the return array is a document that contains two fields k and v:
// The k field contains the field name in the original document.
// The v field contains the value of the field in the original document.
// $objectToArray has the following syntax:
// { $objectToArray: <object> }
// The <object> expression can be any valid expression as long as it resolves to a document object.
// $objectToArray applies to the top-level fields of its argument. If the argument is a document that itself contains embedded document fields, the
// $objectToArray does not recursively apply to the embedded document fields.
func ObjectToArray[T expression.DocumentExpression](e T) objectOperator {
	return objectOperator{doc: bsonx.BsonDoc("$objectToArray", e)}
}

// SetField New in version 5.0.
// Adds, updates, or removes a specified field in a document.
// You can use $setField to add, update, or remove fields with names that contain periods (.) or start with dollar signs ($).
// $setField has the following syntax:
//
//	{
//	 $setField: {
//	   field: <String>,
//	   input: <Object>,
//	   value: <Expression>
//	 }
//	}
func SetField[T expression.AnyExpression, O expression.ObjectExpression](field string, input O, value T) objectOperator {
	doc := bsonx.BsonDoc("field", bsonx.String(field))
	doc.Append("input", input)
	doc.Append("value", value)
	return objectOperator{doc: bsonx.BsonDoc("$setField", doc)}
}
