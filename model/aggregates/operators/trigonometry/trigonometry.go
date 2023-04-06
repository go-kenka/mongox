// Package trigonometry rigonometry expressions perform trigonometric operations
// on numbers. Values that represent angles are always input or output in
// radians. Use $degreesToRadians and $radiansToDegrees to convert between degree
// and radian measurements.
package trigonometry

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type trigonometryOperator struct {
	doc bsonx.Bson
}

func (o trigonometryOperator) Exp() bsonx.IBsonValue {
	return o.doc.Pro()
}

// Sin New in version 4.2. Returns the sine of a value that is measured in
// radians. $sin has the following syntax: { $sin: <expression> } $sin takes any
// valid expression that resolves to a number. If the expression returns a value
// in degrees, use the $degreesToRadians operator to convert the result to
// radians. By default $sin returns values as a double. $sin can also return
// values as a 128-bit decimal as long as the <expression> resolves to a 128-bit
// decimal value.
func Sin[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$sin", n)}
}

// Cos New in version 4.2. Returns the cosine of a value that is measured in
// radians. $cos has the following syntax: { $cos: <expression> } $cos takes any
// valid expression that resolves to a number. If the expression returns a value
// in degrees, use the $degreesToRadians operator to convert the result to
// radians. By default $cos returns values as a double. $cos can also return
// values as a 128-bit decimal as long as the <expression> resolves to a 128-bit
// decimal value.
func Cos[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$cos", n)}
}

// Tan New in version 4.2. Returns the tangent of a value that is measured in
// radians. $tan has the following syntax: { $tan: <expression> } $tan takes any
// valid expression that resolves to a number. If the expression returns a value
// in degrees, use the $degreesToRadians operator to convert the result to
// radians. By default $tan returns values as a double. $tan can also return
// values as a 128-bit decimal as long as the <expression> resolves to a 128-bit
// decimal value.
func Tan[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$tan", n)}
}

// Asin New in version 4.2. Returns the inverse sine (arc sine) of a value. $asin
// has the following syntax: { $asin: <expression> } $asin takes any valid
// expression that resolves to a number between -1 and 1, e.g. -1 <= value <= 1.
// $asin returns values in radians. Use $radiansToDegrees operator to convert the
// output value from radians to degrees. By default $asin returns values as a
// double. $asin can also return values as a 128-bit decimal as long as the
// <expression> resolves to a 128-bit decimal value.
func Asin[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$asin", n)}
}

// Acos New in version 4.2.
// Returns the inverse cosine (arc cosine) of a value.
// $acos has the following syntax:
// { $acos: <expression> }
// $acos takes any valid expression that resolves to a number between -1 and 1, e.g. -1 <= value <= 1.
// $acos returns values in radians. Use $radiansToDegrees operator to convert the output value from radians to degrees.
// By default $acos returns values as a double.
// $acos can also return values as a 128-bit decimal as long as the <expression> resolves to a 128-bit decimal value.
func Acos[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$acos", n)}
}

// Atan New in version 4.2. Returns the inverse tangent (arc tangent) of a value.
// $atan has the following syntax: { $atan: <expression> } $atan takes any valid
// expression that resolves to a number. $atan returns values in radians. Use
// $radiansToDegrees operator to convert the output value from radians to
// degrees. By default $atan returns values as a double. $atan can also return
// values as a 128-bit decimal as long as the <expression> resolves to a 128-bit
// decimal value.
func Atan[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$atan", n)}
}

// Atan2 New in version 4.2. Returns the inverse tangent (arc tangent) of y / x,
// where y and x are the first and second values passed to the expression
// respectively. $atan2 has the following syntax: { $atan2: [ <expression 1>,
// <expression 2> ] } $atan2 takes any valid expression that resolves to a
// number. $atan2 returns values in radians. Use $radiansToDegrees operator to
// convert the output value from radians to degrees. By default $atan2 returns
// values as a double. $atan2 can also return values as a 128-bit decimal as long
// as the <expression> resolves to a 128-bit decimal value.
func Atan2[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$atan2", n)}
}

// Asinh New in version 4.2. Returns the inverse hyperbolic sine (hyperbolic arc
// sine) of a value. $asinh has the following syntax: { $asinh: <expression> }
// $asinh takes any valid expression that resolves to a number. $asinh returns
// values in radians. Use $radiansToDegrees operator to convert the output value
// from radians to degrees. By default $asinh returns values as a double. $asinh
// can also return values as a 128-bit decimal as long as the <expression>
// resolves to a 128-bit decimal value.
func Asinh[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$asinh", n)}
}

// Acosh New in version 4.2. Returns the inverse hyperbolic cosine (hyperbolic
// arc cosine) of a value. $acosh has the following syntax: { $acosh:
// <expression> } $acosh takes any valid expression that resolves to a number
// between 1 and +Infinity, e.g. 1 <= value <= +Infinity. $acosh returns values
// in radians. Use $radiansToDegrees operator to convert the output value from
// radians to degrees. By default $acosh returns values as a double. $acosh can
// also return values as a 128-bit decimal as long as the <expression> resolves
// to a 128-bit decimal value.
func Acosh[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$acosh", n)}
}

// Atanh New in version 4.2. Returns the inverse hyperbolic tangent (hyperbolic
// arc tangent) of a value. $atanh has the following syntax: { $atanh:
// <expression> } $atanh takes any valid expression that resolves to a number
// between -1 and 1, e.g. -1 <= value <= 1. $atanh returns values in radians. Use
// $radiansToDegrees operator to convert the output value from radians to
// degrees. By default $atanh returns values as a double. $atanh can also return
// values as a 128-bit decimal as long as the <expression> resolves to a 128-bit
// decimal value.
func Atanh[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$atanh", n)}
}

// Sinh New in version 4.2. Returns the hyperbolic sine of a value that is
// measured in radians. $sinh has the following syntax: { $sinh: <expression> }
// $sinh takes any valid expression that resolves to a number, measured in
// radians. If the expression returns a value in degrees, use the
// $degreesToRadians operator to convert the value to radians. By default $sinh
// returns values as a double. $sinh can also return values as a 128-bit decimal
// if the <expression> resolves to a 128-bit decimal value.
func Sinh[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$sinh", n)}
}

// Cosh New in version 4.2. Returns the hyperbolic cosine of a value that is
// measured in radians. $cosh has the following syntax: { $cosh: <expression> }
// $cosh takes any valid expression that resolves to a number, measured in
// radians. If the expression returns a value in degrees, use the
// $degreesToRadians operator to convert the value to radians. By default $cosh
// returns values as a double. $cosh can also return values as a 128-bit decimal
// if the <expression> resolves to a 128-bit decimal value.
func Cosh[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$cosh", n)}
}

// Tanh New in version 4.2. Returns the hyperbolic tangent of a value that is
// measured in radians. $tanh has the following syntax: { $tanh: <expression> }
// $tanh takes any valid expression that resolves to a number, measured in
// radians. If the expression returns a value in degrees, use the
// $degreesToRadians operator to convert the value to radians. By default $tanh
// returns values as a double. $tanh can also return values as a 128-bit decimal
// if the <expression> resolves to a 128-bit decimal value.
func Tanh[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$tanh", n)}
}

// DegreesToRadians New in version 4.2. Converts an input value measured in
// degrees to radians. $degreesToRadians has the following syntax: {
// $degreesToRadians: <expression> } $degreesToRadians takes any valid expression
// that resolves to a number. By default $degreesToRadians returns values as a
// double. $degreesToRadians can also return values as a 128-bit decimal as long
// as the <expression> resolves to a 128-bit decimal value.
func DegreesToRadians[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$degreesToRadians", n)}
}

// RadiansToDegrees New in version 4.2. Converts an input value measured in
// radians to degrees. $radiansToDegrees has the following syntax: {
// $radiansToDegrees: <expression> } $radiansToDegrees takes any valid expression
// that resolves to a number. By default $radiansToDegrees returns values as a
// double. $radiansToDegrees can also return values as a 128-bit decimal as long
// as the <expression> resolves to a 128-bit decimal value.
func RadiansToDegrees[T expression.NumberExpression](n T) trigonometryOperator {
	return trigonometryOperator{doc: bsonx.BsonDoc("$radiansToDegrees", n)}
}
