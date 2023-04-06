// Package set Set expressions performs set operation on arrays, treating arrays as sets. Set expressions ignores the duplicate entries in each input array and the order of the elements.
// If the set operation returns a set, the operation filters out duplicates in the result to output an array that contains only unique entries. The order of the elements in the output array is unspecified.
// If a set contains a nested array element, the set expression does not descend into the nested array but evaluates the array at top-level.
package set

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type setOperator struct {
	doc bsonx.Bson
}

func (o setOperator) Exp() bsonx.IBsonValue {
	return o.doc.Pro()
}

// AllElementsTrue Evaluates an array as a set and returns true if no element in the array is false. Otherwise, returns false. An empty array returns true.
// $allElementsTrue has the following syntax:
// { $allElementsTrue: [ <expression> ] }
// The <expression> itself must resolve to an array, separate from the outer array that denotes the argument list
func AllElementsTrue[T expression.ArrayExpression](n T) setOperator {
	return setOperator{doc: bsonx.BsonDoc("$allElementsTrue", bsonx.Array(n))}
}

// AnyElementTrue Evaluates an array as a set and returns true if any of the elements are true and false otherwise. An empty array returns false.
// $anyElementTrue has the following syntax:
// { $anyElementTrue: [ <expression> ] }
// The <expression> itself must resolve to an array, separate from the outer array that denotes the argument list.
func AnyElementTrue[T expression.ArrayExpression](n T) setOperator {
	return setOperator{doc: bsonx.BsonDoc("$anyElementTrue", bsonx.Array(n))}
}

// SetDifference Takes two sets and returns an array containing the elements that only exist in the first set; i.e. performs a
// relative complement
// of the second set relative to the first.
// $setDifference has the following syntax:
// { $setDifference: [ <expression1>, <expression2> ] }
// The arguments can be any valid expression as long as they each resolve to an array
func SetDifference[T expression.ArrayExpression](e1, e2 T) setOperator {
	return setOperator{doc: bsonx.BsonDoc("$setDifference", bsonx.Array(e1, e2))}
}

// SetEquals Compares two or more arrays and returns true if they have the same distinct elements and false otherwise.
// $setEquals has the following syntax:
// { $setEquals: [ <expression1>, <expression2>, ... ] }
// The arguments can be any valid expression as long as they each resolve to an array.
func SetEquals[T expression.ArrayExpression](es ...T) setOperator {
	data := bsonx.Array()
	for _, t := range es {
		data.Append(t)
	}
	return setOperator{doc: bsonx.BsonDoc("$setEquals", data)}
}

// SetIntersection Takes two or more arrays and returns an array that contains the elements that appear in every input array.
// $setIntersection has the following syntax:
// { $setIntersection: [ <array1>, <array2>, ... ] }
// The arguments can be any valid expression as long as they each resolve to an array
func SetIntersection[T expression.ArrayExpression](as ...T) setOperator {
	data := bsonx.Array()
	for _, t := range as {
		data.Append(t)
	}
	return setOperator{doc: bsonx.BsonDoc("$setIntersection", data)}
}

// SetIsSubset Takes two arrays and returns true when the first array is a subset of the second, including when the first array equals the second array, and false otherwise.
// $setIsSubset has the following syntax:
// { $setIsSubset: [ <expression1>, <expression2> ] }
// The arguments can be any valid expression as long as they each resolve to an array
func SetIsSubset[T expression.ArrayExpression](e1, e2 T) setOperator {
	return setOperator{doc: bsonx.BsonDoc("$setIsSubset", bsonx.Array(e1, e2))}
}

// SetUnion Takes two or more arrays and returns an array containing the elements that appear in any input array.
// $setUnion has the following syntax:
// { $setUnion: [ <expression1>, <expression2>, ... ] }
// The arguments can be any valid expression as long as they each resolve to an array.
func SetUnion[T expression.ArrayExpression](as ...T) setOperator {
	data := bsonx.Array()
	for _, t := range as {
		data.Append(t)
	}
	return setOperator{doc: bsonx.BsonDoc("$setUnion", data)}
}
