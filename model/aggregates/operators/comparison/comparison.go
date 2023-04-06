// Package comparison Comparison expressions return a boolean except for $cmp which returns a number.
// The comparison expressions take two argument expressions and compare both value and type,
// using the specified BSON comparison order for values of different types.
package comparison

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/model/aggregates"
)

type comparisonOperator struct {
	doc bsonx.Bson
}

func (o comparisonOperator) Exp() bsonx.IBsonValue {
	return o.doc.ToBsonDocument()
}

// Cmp Compares two values and returns:
// -1 if the first value is less than the second.
// 1 if the first value is greater than the second.
// 0 if the two values are equivalent.
// The $cmp
// compares both value and type, using the specified BSON comparison order for values of different types.
// $cmp has the following syntax:
// { $cmp: [ <expression1>, <expression2> ] }
func Cmp[T expression.AnyExpression](fieldName string, e1, e2 T) comparisonOperator {
	return comparisonOperator{doc: aggregates.NewSimpleFilter(fieldName, bsonx.BsonDoc("$cmp", bsonx.Array(e1, e2)))}
}

// Eq Compares two values and returns:
// true when the values are equivalent.
// false when the values are not equivalent.
// The $eq compares both value and type, using the specified BSON comparison order for values of different types.
// $eq has the following syntax:
// { $eq: [ <expression1>, <expression2> ] }
func Eq[T expression.AnyExpression](fieldName string, value T) comparisonOperator {
	return comparisonOperator{doc: aggregates.NewSimpleFilter(fieldName, value)}
}

// Ne Compares two values and returns:
// true when the values are not equivalent.
// false when the values are equivalent.
// The $ne compares both value and type, using the specified BSON comparison order for values of different types.
// $ne has the following syntax:
// { $ne: [ <expression1>, <expression2> ] }
func Ne[T expression.AnyExpression](fieldName string, value T) comparisonOperator {
	return comparisonOperator{doc: aggregates.NewOperatorFilter("$ne", fieldName, value)}
}

// Gt Compares two values and returns:
// true when the first value is greater than the second value.
// false when the first value is less than or equivalent to the second value.
// The $gt compares both value and type, using the specified BSON comparison order for values of different types.
// $gt has the following syntax:
// { $gt: [ <expression1>, <expression2> ] }
func Gt[T expression.AnyExpression](fieldName string, value T) comparisonOperator {
	return comparisonOperator{doc: aggregates.NewOperatorFilter("$gt", fieldName, value)}
}

// Lt Compares two values and returns:
// true when the first value is less than the second value.
// false when the first value is greater than or equivalent to the second value.
// The $lt compares both value and type, using the specified BSON comparison order for values of different types.
// $lt has the following syntax:
// { $lt: [ <expression1>, <expression2> ] }
func Lt[T expression.AnyExpression](fieldName string, value T) comparisonOperator {
	return comparisonOperator{doc: aggregates.NewOperatorFilter("$lt", fieldName, value)}
}

// Gte Compares two values and returns:
// true when the first value is greater than or equivalent to the second value.
// false when the first value is less than the second value.
// The $gte compares both value and type, using the specified BSON comparison order for values of different types.
// $gte has the following syntax:
// { $gte: [ <expression1>, <expression2> ] }
func Gte[T expression.AnyExpression](fieldName string, value T) comparisonOperator {
	return comparisonOperator{doc: aggregates.NewOperatorFilter("$gte", fieldName, value)}
}

// Lte Compares two values and returns:
// true when the first value is less than or equivalent to the second value.
// false when the first value is greater than the second value.
// The $lte compares both value and type, using the specified BSON comparison order for values of different types.
// $lte has the following syntax:
// { $lte: [ <expression1>, <expression2> ] }
func Lte[T expression.AnyExpression](fieldName string, value T) comparisonOperator {
	return comparisonOperator{doc: aggregates.NewOperatorFilter("$lte", fieldName, value)}
}
