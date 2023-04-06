// Package array Array Expression Operators
package array

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type arrayOperator struct {
	doc bsonx.Bson
}

func (o arrayOperator) Exp() bsonx.IBsonValue {
	return o.doc.Pro()
}

// ArrayElemAt Returns the element at the specified array index.
// $arrayElemAt
// has the following syntax:
// { $arrayElemAt: [ <array>, <idx> ] }
// The <array> expression can be any valid expression that resolves to an array.
// The <idx> expression can be any valid expression that resolves to an integer.
func ArrayElemAt[A expression.ArrayExpression, I expression.IntExpression](array A, idx I) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$arrayElemAt",
		bsonx.Array(array, idx))}
}

// ArrayToObject Converts an array into a single document; the array must be either:
// An array of two-element arrays where the first element is the field name, and the second element is the field value:
// [ [ [ "item", "abc123" ], [ "qty", 25 ] ] ]
// - OR -
// An array of documents that contains two fields, k and v where:
// The k field contains the field name.
// The v field contains the value of the field.
// [ [ { "k": "item", "v": "abc123" }, { "k": "qty", "v": 25 } ] ]
func ArrayToObject[A expression.ArrayExpression](array A) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$arrayToObject", array)}
}

// ConcatArrays Concatenates arrays to return the concatenated array.
// $concatArrays
// has the following syntax:
// { $concatArrays: [ <array1>, <array2>, ... ] }
// The <array> expressions can be any valid expression as long as they resolve to an array. For more information on expressions, see Expressions.
// If any argument resolves to a value of null or refers to a field that is missing,
// $concatArrays
// returns null.
func ConcatArrays[A expression.ArrayExpression](arrays ...A) arrayOperator {
	data := bsonx.Array()
	for _, array := range arrays {
		data.Append(array)
	}
	return arrayOperator{doc: bsonx.BsonDoc("$concatArrays", data)}
}

// Filter Selects a subset of an array to return based on the specified condition.
// Returns an array with only those elements that match the condition.
// The returned elements are in the original order.
func Filter[A expression.ArrayExpression, T expression.BooleanExpression, N expression.NumberExpression](array A, cond T, as string, limit N) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$filter", bsonx.BsonDoc("input", array).
		Append("cond", cond).Append("as", bsonx.String(as)).Append("limit", limit))}
}

// First New in version 4.4.
// Returns the first element in an array.
// $first
// has the following syntax:
// { $first: <expression> }
// The <expression> can be any valid expression as long as it resolves to an array, null or missing. For more information on expressions, see Expressions.
// The
// $first
// operator is an alias for the following $arrayElemAt expression:
// { $arrayElemAt: [ <array expression>, 0 ] }
func First[A expression.ArrayExpression](array A) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$first", array)}
}

// FirstN New in version 5.2.
// Returns a specified number of elements from the beginning of an array.
// $firstN
// has the following syntax:
// { $firstN: { n: <expression>, input: <expression> } }
func FirstN[A expression.ArrayExpression, N expression.IntExpression](array A, n N) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$firstN",
		bsonx.BsonDoc("input", array).Append("n", n))}
}

// In Returns a boolean indicating whether a specified value is in an array.
// $in
// has the following operator expression syntax:
// { $in: [ <expression>, <array expression> ] }
func In[A expression.ArrayExpression, T expression.AnyExpression](e T, array A) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$in", bsonx.Array(e, array))}
}

type IndexOfArrayOptions struct {
	start, end int32
}

// IndexOfArray Searches an array for an occurrence of a specified value and returns the array index (zero-based) of the first occurrence. If the value is not found, returns -1.
// $indexOfArray
// has the following operator expression syntax:
// { $indexOfArray: [ <array expression>, <search expression>, <start>, <end> ]
func IndexOfArray[A expression.ArrayExpression, T expression.AnyExpression](array A, search T, options IndexOfArrayOptions) arrayOperator {
	data := bsonx.Array()
	data.Append(array)
	data.Append(search)
	if options.start > 0 {
		data.Append(bsonx.Int32(options.start))
	}
	if options.end > 0 {
		data.Append(bsonx.Int32(options.end))
	}
	return arrayOperator{doc: bsonx.BsonDoc("$indexOfArray", data)}
}

// IsArray Determines if the operand is an array. Returns a boolean.
// $isArray
// has the following syntax:
// { $isArray: [ <expression> ] }
func IsArray[T expression.AnyExpression](a T) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$isArray", a)}
}

// Last New in version 4.4.
// Returns the last element in an array.
// The $last operator has the following syntax:
// { $last: <expression> }
// The <expression> can be any valid expression as long as it resolves to an array, null, or missing. For more information on expressions, see Expressions.
// The $last operator is an alias for the following $arrayElemAt expression:
// { $arrayElemAt: [ <array expression>, -1 ] }
func Last[A expression.ArrayExpression](array A) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$last", array)}
}

// LastN New in version 5.2.
// Returns a specified number of elements from the end of an array.
// $lastN has the following syntax:
// { $lastN: { n: <expression>, input: <expression> } }
func LastN[A expression.ArrayExpression, N expression.IntExpression](array A, n N) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$lastN", bsonx.BsonDoc("input", array).
		Append("n", n))}
}

// Map Applies an expression to each item in an array and returns an array with the applied results.
// The $map expression has the following syntax:
// { $map: { input: <expression>, as: <string>, in: <expression> } }
func Map[A expression.ArrayExpression, T expression.AnyExpression](input A, in T, as string) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$lastN",
		bsonx.BsonDoc("input", input).
			Append("as", bsonx.String(as)).
			Append("in", in))}
}

// MaxN New in version 5.2.
// Returns the n largest values in an array.
// $maxN has the following syntax:
// { $maxN: { n: <expression>, input: <expression> } }
func MaxN[A expression.ArrayExpression, N expression.IntExpression](array A, n N) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$maxN", bsonx.BsonDoc("input", array).Append("n", n))}
}

// MinN New in version 5.2.
// Returns the n smallest values in an array.
// $minN has the following syntax:
// { $minN: { n: <expression>, input: <expression> } }
func MinN[A expression.ArrayExpression, N expression.IntExpression](array A, n N) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$minN", bsonx.BsonDoc("input", array).Append("n", n))}
}

// Range Returns an array whose elements are a generated sequence of numbers.
// $range generates the sequence from the specified starting number by successively incrementing the starting number by the specified step value up to but not including the end point.
// $range has the following operator expression syntax:
// { $range: [ <start>, <end>, <non-zero step> ] }
func Range[N expression.IntExpression](start, end N, step int32) arrayOperator {
	if step == 0 {
		step = 1
	}
	return arrayOperator{doc: bsonx.BsonDoc("$range", bsonx.Array(start, end, bsonx.Int32(step)))}
}

// Reduce Applies an expression to each element in an array and combines them into a single value.
// $reduce has the following syntax:
//
//	{
//	   $reduce: {
//	       input: <array>,
//	       initialValue: <expression>,
//	       in: <expression>
//	   }
//	}
func Reduce[A expression.ArrayExpression, T expression.AnyExpression](input A, initialValue, in T) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$reduce", bsonx.BsonDoc("input", input).
		Append("initialValue", initialValue).Append("in", in))}
}

// ReverseArray Accepts an array expression as an argument and returns an array with the elements in reverse order.
// $reverseArray has the following operator expression syntax:
// { $reverseArray: <array expression> }
// The argument can be any valid expression as long as it resolves to an array.
func ReverseArray[A expression.ArrayExpression](input A) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$reverseArray", bsonx.BsonDoc("input", input))}
}

// Size Counts and returns the total number of items in an array.
// $size has the following syntax:
// { $size: <expression> }
// The argument for $size
// can be any expression as long as it resolves to an array. For more information on expressions, see Expressions.
func Size[A expression.ArrayExpression](input A) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$size", input)}
}

// Slice Returns a subset of an array.
// $slice has one of two syntax forms:
// The following syntax returns elements from either the start or end of the array:
// { $slice: [ <array>, <n> ] }
// The following syntax returns elements from the specified position in the array:
// { $slice: [ <array>, <position>, <n> ] }
func Slice[A expression.ArrayExpression, N expression.IntExpression](input A, position, n N) arrayOperator {
	data := bsonx.Array()
	data.Append(input)
	if position != nil {
		data.Append(position)
	}
	data.Append(n)
	return arrayOperator{doc: bsonx.BsonDoc("$slice", data)}
}

// SortArray New in version 5.2.
// Sorts an array based on its elements. The sort order is user specified.
// $sortArray  has the following syntax:
//
//	$sortArray: {
//	  input: <array>,
//	  sortBy: <sort spec>
//	}
func SortArray[A expression.ArrayExpression](input A, sortBy bsonx.Bson) arrayOperator {
	return arrayOperator{doc: bsonx.BsonDoc("$sortArray", bsonx.BsonDoc("input", input).
		Append("sortBy", sortBy.Pro()))}
}

// Zip Transposes an array of input arrays so that the first element of the output array would be an array containing, the first element of the first input array, the first element of the second input array, etc.
// For example,
// $zip would transform [ [ 1, 2, 3 ], [ "a", "b", "c" ] ] into [ [ 1, "a" ], [ 2, "b" ], [ 3, "c" ] ].
// $zip has the following syntax:
//
//	{
//	   $zip: {
//	       inputs: [ <array expression1>,  ... ],
//	       useLongestLength: <boolean>,
//	       defaults:  <array expression>
//	   }
//	}
func Zip[A expression.ArrayExpression](inputs []A, useLongestLength bool, defaults []A) arrayOperator {
	inputData := bsonx.Array()
	for _, array := range inputs {
		inputData.Append(array)
	}

	doc := bsonx.BsonDoc("inputs", inputData)
	if useLongestLength {
		doc.Append("useLongestLength", bsonx.Boolean(true))
	}
	if len(defaults) > 0 {
		defaultData := bsonx.Array()
		for _, array := range defaults {
			defaultData.Append(array)
		}
		doc.Append("defaults", defaultData)
	}

	return arrayOperator{doc: bsonx.BsonDoc("$zip", doc)}
}
