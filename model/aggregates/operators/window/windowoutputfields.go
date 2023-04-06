package window

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/model/aggregates"
	"go.mongodb.org/mongo-driver/bson"
)

type outputOperator struct {
	doc bsonx.Bson
}

func (o outputOperator) Exp() bsonx.IBsonValue {
	return o.doc.ToBsonDocument()
}

// AddToSet Returns an array of all unique values that results from applying an
// expression to each document. Changed in version 5.0: Available in
// $setWindowFields stage.
func AddToSet[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$addToSet", e, window)
}

// Avg Returns the average for the specified expression . Ignores non-numeric
// values. Changed in version 5.0: Available in $setWindowFields stage.
func Avg[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$avg", e, window)
}

// Bottom Returns the bottom element within a group according to the specified
// sort order. NewStage in version 5.2. Available in $group and $setWindowFields
// stages.
func Bottom[O expression.AnyExpression](path string, sortBy bsonx.Bson, out O, window Window) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[SortBy] = sortBy.ToBsonDocument()
	args[Output] = out
	return compoundParameterWindowFunction(path, "$bottom", args, window)
}

// BottomN Returns an aggregation of the bottom n fields within a group,
// according to the specified sort order. NewStage in version 5.2. Available in $group
// and $setWindowFields stages.
func BottomN[I expression.AnyExpression, O expression.AnyExpression, N expression.IntExpression](path string, in I, out O, n N, window Window) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[Input] = in
	args[Output] = out
	args[NLowercase] = n
	return compoundParameterWindowFunction(path, "$bottomN", args, window)
}

// Count Returns the number of documents in the group or window. Distinct from
// the $count pipeline stage. NewStage in version 5.0.
func Count(path string, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$count", nil, window)
}

// CovariancePop Returns the population covariance of two numeric expressions.
// NewStage in version 5.0.
func CovariancePop[T expression.AnyExpression](path string, expr1, expr2 T, window Window) outputOperator {
	args := bsonx.Array()
	args.Append(expr1)
	args.Append(expr2)
	return simpleParameterWindowFunction(path, "$covariancePop", args, window)
}

// CovarianceSamp Returns the sample covariance of two numeric expressions. NewStage
// in version 5.0.
func CovarianceSamp[T expression.AnyExpression](path string, expr1, expr2 T, window Window) outputOperator {
	args := bsonx.Array()
	args.Append(expr1)
	args.Append(expr2)
	return simpleParameterWindowFunction(path, "$covarianceSamp", args, window)
}

// DenseRank Returns the document position (known as the rank) relative to other
// documents in the $setWindowFields stage partition. There are no gaps in the
// ranks. Ties receive the same rank. NewStage in version 5.0.
func DenseRank(path string) outputOperator {
	return simpleParameterWindowFunction(path, "$denseRank", nil, nil)
}

// Derivative Returns the average rate of change within the specified window. NewStage
// in version 5.0.
func Derivative[I expression.AnyExpression](path string, in I, window Window) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[Input] = in
	return compoundParameterWindowFunction(path, "$derivative", args, window)
}

// TimeDerivative Returns the average rate of change within the specified window. NewStage
// in version 5.0.
func TimeDerivative[T expression.AnyExpression](path string, e T, window Window, unit aggregates.MongoTimeUnit) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[Input] = e
	args[Unit] = bsonx.String(unit.GetValue())
	return compoundParameterWindowFunction(path, "$derivative", args, window)
}

// DocumentNumber Returns the position of a document (known as the document number) in the
// $setWindowFields stage partition. Ties result in different adjacent document
// numbers. NewStage in version 5.0.
func DocumentNumber(path string) outputOperator {
	return simpleParameterWindowFunction(path, "$documentNumber", nil, nil)
}

// ExpMovingAvg Returns the exponential moving average for the numeric
// expression. NewStage in version 5.0.
func ExpMovingAvg[T expression.AnyExpression](path string, e T, n int32) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[Input] = e
	args[NUppercase] = bsonx.Int32(n)
	return compoundParameterWindowFunction(path, "$covariancePop", args, nil)
}

// ExpMovingAvgAlpha Returns the exponential moving average for the numeric
// expression. NewStage in version 5.0.
func ExpMovingAvgAlpha[T expression.AnyExpression](path string, e T, alpha float64) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[Input] = e
	args[Alpha] = bsonx.Double(alpha)
	return compoundParameterWindowFunction(path, "$covariancePop", args, nil)
}

// First Returns the value that results from applying an expression to the first
// document in a group or window. Changed in version 5.0: Available in
// $setWindowFields stage.
func First[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$first", e, window)
}

// Integral Returns the approximation of the area under a curve.
// NewStage in version 5.0.
func Integral[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[Input] = e
	return compoundParameterWindowFunction(path, "$integral", args, window)
}

// TimeIntegral Returns the approximation of the area under a curve.
// NewStage in version 5.0.
func TimeIntegral[T expression.AnyExpression](path string, e T, window Window, unit aggregates.MongoTimeUnit) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[Input] = e
	args[Unit] = bsonx.String(unit.GetValue())
	return compoundParameterWindowFunction(path, "$integral", args, window)
}

// Last Returns the value that results from applying an
// expression
// to the last document in a group or window.
// Changed in version 5.0: Available in $setWindowFields stage.
func Last[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$last", e, window)
}

// LinearFill Fills null and missing fields in a window using linear interpolation based on
// surrounding field values. Available in $setWindowFields stage. NewStage in version
// 5.3.
func LinearFill[T expression.AnyExpression](path string, e T) outputOperator {
	return simpleParameterWindowFunction(path, "$linearFill", e, nil)
}

// Locf Last observation carried forward. Sets values for null and missing fields in a
// window to the last non-null value for the field. Available in $setWindowFields
// stage. NewStage in version 5.2.
func Locf[T expression.AnyExpression](path string, e T) outputOperator {
	return simpleParameterWindowFunction(path, "$locf", e, nil)
}

// Max Returns the maximum value that results from applying an
// expression
// to each document.
// Changed in version 5.0: Available in $setWindowFields stage.
func Max[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$max", e, window)
}

// Min Returns the minimum value that results from applying an expression to each
// document. Changed in version 5.0: Available in $setWindowFields stage.
func Min[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$min", e, window)
}

// MinN Returns an aggregation of the n minimum valued elements in a group.
// Distinct from the $minN array operator. NewStage in version 5.2. Available in
// $group, $setWindowFields and as an expression.
func MinN[I expression.AnyExpression, N expression.IntExpression](path string, in I, n N, window Window) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[Input] = in
	args[NLowercase] = n

	return compoundParameterWindowFunction(path, "$minN", args, window)
}

// Push Returns an array of values that result from applying an expression to
// each document.Changed Changed in version 5.0: Available in $setWindowFields
// stage.
func Push[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$push", e, window)
}

// Rank Returns the document position (known as the rank) relative to other
// documents in the $setWindowFields stage partition.
// NewStage in version 5.0.
func Rank(path string) outputOperator {
	return simpleParameterWindowFunction(path, "$rank", nil, nil)
}

// Shift Returns the value from an expression applied to a document in a specified
// position relative to the current document in the $setWindowFields stage
// partition. NewStage in version 5.0.
func Shift[T expression.AnyExpression](path string, e, defaultExpression T, by int32) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[Output] = e
	args[By] = bsonx.Int32(by)
	if defaultExpression != nil {
		args[Default] = defaultExpression
	}
	return compoundParameterWindowFunction(path, "$shift", args, nil)
}

// StdDevPop Returns the population standard deviation that results from applying
// a numeric expression to each document. Changed in version 5.0: Available in
// $setWindowFields stage.
func StdDevPop[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$stdDevPop", e, window)
}

// StdDevSamp Returns the sample standard deviation that results from applying a
// numeric expression to each document. Changed in version 5.0: Available in
// $setWindowFields stage.
func StdDevSamp[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$stdDevSamp", e, window)
}

// Sum Returns the sum that results from applying a numeric expression to each
// document. Changed in version 5.0: Available in $setWindowFields stage.
func Sum[T expression.AnyExpression](path string, e T, window Window) outputOperator {
	return simpleParameterWindowFunction(path, "$sum", e, window)
}

// Top Returns the top element within a group according to the specified sort
// order. NewStage in version 5.2. Available in $group and $setWindowFields stages.
func Top[O expression.AnyExpression](path string, sortBy bsonx.Bson, out O, window Window) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[SortBy] = sortBy.ToBsonDocument()
	args[Output] = out
	return compoundParameterWindowFunction(path, "$top", args, window)
}

// TopN Returns an aggregation of the top n fields within a group, according to
// the specified sort order. NewStage in version 5.2. Available in $group and
// $setWindowFields stages.
func TopN[O expression.AnyExpression, N expression.IntExpression](path string, sortBy bsonx.Bson, out O, n N, window Window) outputOperator {
	args := make(map[ParamName]expression.AnyExpression)
	args[SortBy] = sortBy.ToBsonDocument()
	args[Output] = out
	args[NLowercase] = n
	return compoundParameterWindowFunction(path, "$topN", args, window)
}

func simpleParameterWindowFunction[T expression.Expression](path, functionName string, e T, window Window) outputOperator {
	return outputField(path, NewSimpleParameterFunctionAndWindow[T](functionName, e, window))
}
func compoundParameterWindowFunction[T expression.Expression](path, functionName string, args map[ParamName]T, window Window) outputOperator {
	return outputField(path, NewCompoundParameterFunctionAndWindow[T](functionName, args, window))
}

func outputField[T expression.Expression](path string, win T) outputOperator {
	return outputOperator{
		doc: bsonx.BsonDoc(path, win.Exp()),
	}
}

type SimpleParameterFunctionAndWindow[T expression.Expression] struct {
	AbstractFunctionAndWindow
	expression T
}

func NewSimpleParameterFunctionAndWindow[T expression.Expression](functionName string, expression T, window Window) SimpleParameterFunctionAndWindow[T] {
	return SimpleParameterFunctionAndWindow[T]{
		AbstractFunctionAndWindow: NewAbstractFunctionAndWindow(functionName, window),
		expression:                expression,
	}
}

func (a SimpleParameterFunctionAndWindow[T]) Exp() bsonx.IBsonValue {
	return a.ToBsonDocument()
}

func (a SimpleParameterFunctionAndWindow[T]) ToBsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	if a.expression != nil {
		doc.Append(a.functionName, a.expression.Exp())
	} else {
		doc.Append(a.functionName, bsonx.BsonEmpty())
	}
	a.writeWindow(doc)
	return doc
}

func (a SimpleParameterFunctionAndWindow[T]) Document() bson.D {
	return a.ToBsonDocument().Document()
}

type AbstractFunctionAndWindow struct {
	functionName string
	window       Window
}

func NewAbstractFunctionAndWindow(functionName string, window Window) AbstractFunctionAndWindow {
	return AbstractFunctionAndWindow{
		functionName: functionName,
		window:       window,
	}
}

func (a AbstractFunctionAndWindow) Exp() bsonx.IBsonValue {
	return a.ToBsonDocument()
}

func (a AbstractFunctionAndWindow) ToBsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc("window", a.window.ToBsonDocument())
}

func (a AbstractFunctionAndWindow) Document() bson.D {
	return a.ToBsonDocument().Document()
}

func (a AbstractFunctionAndWindow) writeWindow(doc *bsonx.BsonDocument) {
	if a.window != nil {
		doc.Append("window", a.window.ToBsonDocument())
	}
}

type ParamName struct {
	value string
}

func NewParamName(value string) ParamName {
	return ParamName{value: value}
}

var (
	Input      = NewParamName("input")
	Unit       = NewParamName("unit")
	NUppercase = NewParamName("N")
	NLowercase = NewParamName("n")
	Alpha      = NewParamName("alpha")
	Output     = NewParamName("output")
	By         = NewParamName("by")
	Default    = NewParamName("default")
	SortBy     = NewParamName("sortBy")
)

type CompoundParameterFunctionAndWindow[T expression.Expression] struct {
	AbstractFunctionAndWindow
	args map[ParamName]T
}

func NewCompoundParameterFunctionAndWindow[T expression.Expression](functionName string, args map[ParamName]T, window Window) CompoundParameterFunctionAndWindow[T] {
	return CompoundParameterFunctionAndWindow[T]{
		AbstractFunctionAndWindow: NewAbstractFunctionAndWindow(functionName, window),
		args:                      args,
	}
}

func (a CompoundParameterFunctionAndWindow[T]) ToBsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	args := bsonx.BsonEmpty()
	for name, value := range a.args {
		args.Append(name.value, value.Exp())
	}

	a.writeWindow(doc)
	return doc
}

func (a CompoundParameterFunctionAndWindow[T]) Exp() bsonx.IBsonValue {
	return a.ToBsonDocument()
}

func (a CompoundParameterFunctionAndWindow[T]) Document() bson.D {
	return a.ToBsonDocument().Document()
}
