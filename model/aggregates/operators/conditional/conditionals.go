// Package conditional Conditional Expression Operators
package conditional

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/bsonx/expression"
	"github.com/go-kenka/mongox/model/aggregates"
)

type conditionalOperator struct {
	doc bsonx.Bson
}

func (o conditionalOperator) Exp() bsonx.IBsonValue {
	return o.doc.BsonDocument()
}

// Cond Evaluates a boolean expression to return one of the two specified return expressions.
// The $cond expression has one of two syntaxes:
// { $cond: { if: <boolean-expression>, then: <true-case>, else: <false-case> } }
// Or:
// { $cond: [ <boolean-expression>, <true-case>, <false-case> ] }
// $cond requires all three arguments (if-then-else) for either syntax.
// If the <boolean-expression> evaluates to true, then
// $cond evaluates and returns the value of the <true-case> expression. Otherwise,
// $cond evaluates and returns the value of the <false-case> expression.
// The arguments can be any valid expression.
func Cond[T expression.AnyExpression](fieldName string, ifExpr T, thenExpr, elseExpr T) conditionalOperator {
	return conditionalOperator{doc: aggregates.NewSimpleFilter(fieldName, bsonx.BsonDoc("$cond", bsonx.BsonDoc("if", ifExpr).
		Append("then", thenExpr).
		Append("else", elseExpr)))}
}

// IfNull Changed in version 5.0.
// The $ifNull expression evaluates input expressions for null values and returns:
// The first non-null input expression value found.
// A replacement expression value if all input expressions evaluate to null.
// $ifNull treats undefined values and missing fields as null.
// Syntax:
//
//	{
//	  $ifNull: [
//	     <input-expression-1>,
//	     ...
//	     <input-expression-n>,
//	     <replacement-expression-if-null>
//	  ]
//	}
//
// In MongoDB 4.4 and earlier versions,
// $ifNull only accepts a single input expression:
//
//	{
//	  $ifNull: [
//	     <input-expression>,
//	     <replacement-expression-if-null>
//	  ]
//	}
func IfNull[T expression.AnyExpression](fieldName string, es []T, replace T) conditionalOperator {
	data := bsonx.Array()
	for _, e := range es {
		data.Append(e)
	}
	data.Append(replace)
	return conditionalOperator{doc: aggregates.NewSimpleFilter(fieldName, bsonx.BsonDoc("$ifNull", data))}
}

type Branch[T expression.AnyExpression] struct {
	Case T
	Then T
}

// Switch Evaluates a series of case expressions. When it finds an expression which evaluates to true, $switch executes a specified expression and breaks out of the control flow.
// $switch has the following syntax:
//
//	$switch: {
//	  branches: [
//	     { case: <expression>, then: <expression> },
//	     { case: <expression>, then: <expression> },
//	     ...
//	  ],
//	  default: <expression>
//	}
//
// The objects in the branches array must contain only a case field and a then field.
func Switch[T expression.AnyExpression](fieldName string, branches []Branch[T], defaultExpr T) conditionalOperator {
	bs := bsonx.Array()
	for _, branch := range branches {
		bs.Append(bsonx.BsonDoc("case", branch.Case).Append("then", branch.Then))
	}
	data := bsonx.BsonDoc("branches", bs)
	data.Append("default", defaultExpr)
	return conditionalOperator{doc: aggregates.NewSimpleFilter(fieldName, bsonx.BsonDoc("$switch", data))}
}
