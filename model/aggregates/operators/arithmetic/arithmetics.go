// Package arithmetic Arithmetic expressions perform mathematic operations on numbers.
// Some arithmetic expressions can also support date arithmetic.
package arithmetic

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/bsonx/expression"
)

type arithmeticOperator struct {
	doc bsonx.Bson
}

func (o arithmeticOperator) Exp() bsonx.IBsonValue {
	return o.doc.BsonDocument()
}

// Abs Returns the absolute value of a number.
// $abs has the following syntax:
// { $abs: <number> }
// The <number> expression can be any valid expression as long as it resolves to a number.
func Abs[T expression.NumberExpression](number T) arithmeticOperator {
	return arithmeticOperator{
		doc: bsonx.BsonDoc("$abs", number),
	}
}

// Add Adds numbers together or adds numbers and a date. If one of the arguments is a date,
// $add treats the other arguments as milliseconds to add to the date.
// The $add expression has the following syntax:
// { $add: [ <expression1>, <expression2>, ... ] }
// The arguments can be any valid expression as long as they resolve to either all numbers or to numbers and a date.
func Add[T expression.DateNumberExpression](n []T) arithmeticOperator {
	data := bsonx.Array()
	for _, t := range n {
		data.Append(t)
	}
	return arithmeticOperator{doc: bsonx.BsonDoc("$add", data)}
}

// Ceil Returns the smallest integer greater than or equal to the specified number.
// $ceil has the following syntax:
// { $ceil: <number> }
// The <number> expression can be any valid expression as long as it resolves to a number.
func Ceil[T expression.NumberExpression](n T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$ceil", n)}
}

// Divide Divides one number by another and returns the result. Pass the arguments to
// $divide in an array.
// The $divide expression has the following syntax:
// { $divide: [ <expression1>, <expression2> ] }
// The first argument is the dividend, and the second argument is the divisor; i.e. the first argument is divided by the second argument.
// The arguments can be any valid expression as long as they resolve to numbers.
func Divide[T expression.NumberExpression](n1, n2 T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$divide", bsonx.Array(n1, n2))}
}

// Exp Raises Euler's number (i.e. e ) to the specified exponent and returns the result.
// $exp has the following syntax:
// { $exp: <exponent> }
// The <exponent> expression can be any valid expression as long as it resolves to a number
func Exp[T expression.NumberExpression](n T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$exp", n)}
}

// Floor Returns the largest integer less than or equal to the specified number.
// $floor has the following syntax:
// { $floor: <number> }
// The <number> expression can be any valid expression as long as it resolves to a number.
func Floor[T expression.NumberExpression](n T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$floor", n)}
}

// Ln Calculates the natural logarithm ln (i.e log e) of a number and returns the result as a double.
// $ln has the following syntax:
// { $ln: <number> }
// The <number> expression can be any valid expression as long as it resolves to a non-negative number. For more information on expressions, see Expressions.
// $ln is equivalent to $log: [ <number>, Math.E ] expression, where Math.E is a JavaScript representation for Euler's number e.
func Ln[T expression.NumberExpression](n T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$ln", n)}
}

// Log Calculates the log of a number in the specified base and returns the result as a double.
// $log has the following syntax:
// { $log: [ <number>, <base> ] }
// The <number> expression can be any valid expression as long as it resolves to a non-negative number.
// The <base> expression can be any valid expression as long as it resolves to a positive number greater than 1.
func Log[T expression.NumberExpression](n, base T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$log", bsonx.Array(n, base))}
}

// Log10 Calculates the log base 10 of a number and returns the result as a double.
// $log10 has the following syntax:
// { $log10: <number> }
// The <number> expression can be any valid expression as long as it resolves to a non-negative number. For more information on expressions, see Expressions.
// $log10 is equivalent to $log: [ <number>, 10 ] expression.
func Log10[T expression.NumberExpression](n T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$log10", n)}
}

// Mod Divides one number by another and returns the remainder.
// The $mod expression has the following syntax:
// { $mod: [ <expression1>, <expression2> ] }
// The first argument is the dividend, and the second argument is the divisor; i.e. first argument is divided by the second argument.
// The arguments can be any valid expression as long as they resolve to numbers.
func Mod[T expression.NumberExpression](n1, n2 T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$mod", bsonx.Array(n1, n2))}
}

// Multiply Multiplies numbers together and returns the result. Pass the arguments to
// $multiply in an array.
// The $multiply expression has the following syntax:
// { $multiply: [ <expression1>, <expression2>, ... ] }
// The arguments can be any valid expression as long as they resolve to numbers.
func Multiply[T expression.NumberExpression](ns []T) arithmeticOperator {
	data := bsonx.Array()
	for _, t := range ns {
		data.Append(t)
	}
	return arithmeticOperator{doc: bsonx.BsonDoc("$multiply", data)}
}

// Pow  Raises a number to the specified exponent and returns the result.
// $pow has the following syntax:
// { $pow: [ <number>, <exponent> ] }
// The <number> expression can be any valid expression as long as it resolves to a number.
// The <exponent> expression can be any valid expression as long as it resolves to a number.
// You cannot raise 0 to a negative exponent.
func Pow[T expression.NumberExpression](n, e T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$pow", bsonx.Array(n, e))}
}

// Round New in version 4.2.
// $round rounds a number to a whole integer or to a specified decimal place.
// $round has the following syntax:
// { $round : [ <number>, <place> ] }
func Round[N expression.NumberExpression, I expression.IntExpression](n N, p I) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$round", bsonx.Array(n, p))}
}

// Sqrt Calculates the square root of a positive number and returns the result as a double.
// $sqrt has the following syntax:
// { $sqrt: <number> }
// The argument can be any valid expression as long as it resolves to a non-negative number
func Sqrt[T expression.NumberExpression](n T) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$sqrt", n)}
}

// Subtract Subtracts two numbers to return the difference, or two dates to return the difference in milliseconds, or a date and a number in milliseconds to return the resulting date.
// The $subtract expression has the following syntax:
// { $subtract: [ <expression1>, <expression2> ] }
// The second argument is subtracted from the first argument.
// The arguments can be any valid expression as long as they resolve to numbers and/or dates. To subtract a number from a date, the date must be the first argument.
func Subtract[N expression.DateNumberExpression](n1, n2 N) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$subtract", bsonx.Array(n1, n2))}
}

// Trunc Changed in version 4.2..
// $trunc truncates a number to a whole integer or to a specified decimal place.
// MongoDB 4.2 adds the following syntax for $trunc:
// { $trunc : [ <number>, <place> ] }
func Trunc[N expression.NumberExpression, I expression.IntExpression](n N, p I) arithmeticOperator {
	return arithmeticOperator{doc: bsonx.BsonDoc("$trunc", bsonx.Array(n, p))}
}
