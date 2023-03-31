// Package types Type Expression Operators
package types

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type typeOperator struct {
	doc bsonx.Bson
}

func (o typeOperator) Exp() bsonx.IBsonValue {
	return o.doc.ToBsonDocument()
}

// Convert Converts a value to a specified type.
// $convert has the following syntax:
//
//	{
//	  $convert:
//	     {
//	        input: <expression>, to: <type expression>, onError: <expression>, //
//	        Optional. onNull: <expression> // Optional.
//	     }
//	}
func Convert[E expression.AnyExpression, T expression.IntStrExpression](input E, to T, options ConvertOptions) typeOperator {
	data := bsonx.BsonDoc("input", input)
	data.Append("to", to)
	if options.onError != nil {
		data.Append("onError", options.onError.ToBsonDocument())
	}
	if options.onNull != nil {
		data.Append("onNull", options.onNull.ToBsonDocument())
	}

	return typeOperator{doc: bsonx.BsonDoc("$convert", data)}
}

// IsNumber New in version 4.4. $isNumber checks if the specified expression resolves to
// one of the following numeric BSON types: # Integer # Decimal # Double # Long
// $isNumber returns: true if the expression resolves to a number. false if the
// expression resolves to any other BSON type, null, or a missing field.
// $isNumber has the following operator expression syntax: { $isNumber:
// <expression> }
func IsNumber[T expression.AnyExpression](e T) typeOperator {
	return typeOperator{doc: bsonx.BsonDoc("$isNumber", e)}
}

// ToBool Converts a value to a boolean.
// $toBool
// has the following syntax:
//
//	{
//	  $toBool: <expression>
//	}
//
// The $toBool takes any valid expression. The $toBool is a shorthand for the
// following $convert expression: { $convert: { input: <expression>, to: "bool" }
// }
func ToBool[T expression.AnyExpression](e T) typeOperator {
	return typeOperator{doc: bsonx.BsonDoc("$toBool", e)}
}

// ToDate Converts a value to a date. If the value cannot be converted to a date,
// $toDate errors. If the value is null or missing, $toDate returns null. $toDate
// has the following syntax:
//
//	{
//	  $toDate: <expression>
//	}
//
// The $toDate takes any valid expression. The $toDate is a shorthand for the
// following $convert expression: { $convert: { input: <expression>, to: "date" }
// }
func ToDate[T expression.AnyExpression](e T) typeOperator {
	return typeOperator{doc: bsonx.BsonDoc("$toDate", e)}
}

// ToDecimal Converts a value to a decimal. If the value cannot be converted to a
// decimal, $toDecimal errors. If the value is null or missing, $toDecimal
// returns null. $toDecimal has the following syntax:
//
//	{
//	  $toDecimal: <expression>
//	}
//
// The $toDecimal takes any valid expression. The $toDecimal is a shorthand for
// the following $convert expression: { $convert: { input: <expression>, to:
// "decimal" } }
func ToDecimal[T expression.AnyExpression](e T) typeOperator {
	return typeOperator{doc: bsonx.BsonDoc("$toDecimal", e)}
}

// ToDouble Converts a value to a double. If the value cannot be converted to an double,
// $toDouble errors. If the value is null or missing, $toDouble returns null.
// $toDouble has the following syntax:
//
//	{
//	  $toDouble: <expression>
//	}
//
// The
// $toDouble
// takes any valid expression.
// The
// $toDouble
// is a shorthand for the following $convert expression:
// { $convert: { input: <expression>, to: "double" } }
func ToDouble[T expression.AnyExpression](e T) typeOperator {
	return typeOperator{doc: bsonx.BsonDoc("$toDouble", e)}
}

// ToInt Converts a value to an integer. If the value cannot be converted to an
// integer, $toInt errors. If the value is null or missing, $toInt returns null.
// $toInt has the following syntax:
//
//	{
//	  $toInt: <expression>
//	}
//
// The $toInt takes any valid expression. The $toInt is a shorthand for the
// following $convert expression: { $convert: { input: <expression>, to: "int" }
// }
func ToInt[T expression.AnyExpression](e T) typeOperator {
	return typeOperator{doc: bsonx.BsonDoc("$toInt", e)}
}

// ToLong Converts a value to a long. If the value cannot be converted to a long,
// $toLong errors. If the value is null or missing, $toLong returns null. $toLong
// has the following syntax:
//
//	{
//	  $toLong: <expression>
//	}
//
// The $toLong takes any valid expression. The $toLong is a shorthand for the
// following $convert expression: { $convert: { input: <expression>, to: "long" }
// }
func ToLong[T expression.AnyExpression](e T) typeOperator {
	return typeOperator{doc: bsonx.BsonDoc("$toLong", e)}
}

// ToObjectId Converts a value to an ObjectId(). If the value cannot be converted to an
// ObjectId, $toObjectId errors. If the value is null or missing, $toObjectId
// returns null. $toObjectId has the following syntax:
//
//	{
//	  $toObjectId: <expression>
//	}
//
// The $toObjectId takes any valid expression. The $toObjectId is a shorthand for
// the following $convert expression: { $convert: { input: <expression>, to:
// "objectId" } }
func ToObjectId[T expression.AnyExpression](e T) typeOperator {
	return typeOperator{doc: bsonx.BsonDoc("$toObjectId", e)}
}

// Type Returns a string that specifies the BSON type of the argument. $type has
// the following operator expression syntax: { $type: <expression> } The argument
// can be any valid expression.
func Type[T expression.AnyExpression](e T) typeOperator {
	return typeOperator{doc: bsonx.BsonDoc("$type", e)}
}
