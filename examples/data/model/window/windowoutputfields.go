package window

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/aggregates"
	"github.com/go-kenka/mongox/examples/data/model/expressions"
)

type windowOutputFields struct {
}

func WindowOutputFields() windowOutputFields {
	return windowOutputFields{}
}

func (w windowOutputFields) of(windowOutputField bsonx.BsonField) WindowOutputField {
	return NewBsonFieldWindowOutputField(windowOutputField)
}
func (w windowOutputFields) sum(path string, expression expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$sum", expression, window)
}
func (w windowOutputFields) avg(path string, expression expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$avg", expression, window)
}
func (w windowOutputFields) stdDevSamp(path string, expression expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$stdDevSamp", expression, window)
}
func (w windowOutputFields) stdDevPop(path string, expression expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$stdDevPop", expression, window)
}
func (w windowOutputFields) min(path string, expression expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$min", expression, window)
}
func (w windowOutputFields) minN(path string, in expressions.InExpression, n expressions.NExpression, window Window) WindowOutputField {

	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = in
	args[N_LOWERCASE] = n

	return w.compoundParameterWindowFunction(path, "$minN", args, window)
}
func (w windowOutputFields) max(path string, expression expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$max", expression, window)
}
func (w windowOutputFields) maxN(path string, in expressions.InExpression, n expressions.NExpression, window Window) WindowOutputField {

	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = in
	args[N_LOWERCASE] = n

	return w.compoundParameterWindowFunction(path, "$maxN", args, window)
}
func (w windowOutputFields) count(path string, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$count", nil, window)
}
func (w windowOutputFields) derivative(path string, in expressions.TExpression, window Window) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = in
	return w.compoundParameterWindowFunction(path, "$derivative", args, window)
}
func (w windowOutputFields) timeDerivative(path string, expr expressions.TExpression, window Window, unit aggregates.MongoTimeUnit) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = expr
	args[UNIT] = bsonx.NewBsonString(unit.GetValue())
	return w.compoundParameterWindowFunction(path, "$derivative", args, window)
}
func (w windowOutputFields) integral(path string, expr expressions.TExpression, window Window) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = expr
	return w.compoundParameterWindowFunction(path, "$integral", args, window)
}
func (w windowOutputFields) timeIntegral(path string, expr expressions.TExpression, window Window, unit aggregates.MongoTimeUnit) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = expr
	args[UNIT] = bsonx.NewBsonString(unit.GetValue())
	return w.compoundParameterWindowFunction(path, "$integral", args, window)
}
func (w windowOutputFields) covarianceSamp(path string, expr1, expr2 expressions.TExpression, window Window) WindowOutputField {
	args := bsonx.NewBsonArray()
	args.Append(expr1)
	args.Append(expr2)
	return w.simpleParameterWindowFunction(path, "$covarianceSamp", args, window)
}
func (w windowOutputFields) covariancePop(path string, expr1, expr2 expressions.TExpression, window Window) WindowOutputField {
	args := bsonx.NewBsonArray()
	args.Append(expr1)
	args.Append(expr2)
	return w.simpleParameterWindowFunction(path, "$covariancePop", args, window)
}
func (w windowOutputFields) expMovingAvg(path string, expr expressions.TExpression, n int32) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = expr
	args[N_UPPERCASE] = bsonx.NewBsonInt32(n)
	return w.compoundParameterWindowFunction(path, "$covariancePop", args, nil)
}

func (w windowOutputFields) expMovingAvgAlpha(path string, expr expressions.TExpression, alpha float64) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = expr
	args[ALPHA] = bsonx.NewBsonDouble(alpha)
	return w.compoundParameterWindowFunction(path, "$covariancePop", args, nil)
}
func (w windowOutputFields) push(path string, expr expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$push", expr, window)
}
func (w windowOutputFields) addToSet(path string, expr expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$addToSet", expr, window)
}
func (w windowOutputFields) first(path string, expr expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$first", expr, window)
}
func (w windowOutputFields) firstN(path string, in expressions.InExpression, n expressions.NExpression, window Window) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = in
	args[N_LOWERCASE] = n
	return w.compoundParameterWindowFunction(path, "$firstN", args, window)
}
func (w windowOutputFields) top(path string, sortBy bsonx.Bson, out expressions.OutExpression, window Window) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[SORT_BY] = sortBy.ToBsonDocument()
	args[OUTPUT] = out
	return w.compoundParameterWindowFunction(path, "$top", args, window)
}
func (w windowOutputFields) topN(path string, sortBy bsonx.Bson, out expressions.OutExpression, n expressions.NExpression, window Window) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[SORT_BY] = sortBy.ToBsonDocument()
	args[OUTPUT] = out
	args[N_LOWERCASE] = n
	return w.compoundParameterWindowFunction(path, "$topN", args, window)
}
func (w windowOutputFields) last(path string, expr expressions.TExpression, window Window) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$last", expr, window)
}
func (w windowOutputFields) lastN(path string, in expressions.InExpression, n expressions.NExpression, window Window) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = in
	args[N_LOWERCASE] = n
	return w.compoundParameterWindowFunction(path, "$lastN", args, window)
}
func (w windowOutputFields) bottom(path string, sortBy bsonx.Bson, out expressions.OutExpression, window Window) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[SORT_BY] = sortBy.ToBsonDocument()
	args[OUTPUT] = out
	return w.compoundParameterWindowFunction(path, "$bottom", args, window)
}
func (w windowOutputFields) bottomN(path string, in expressions.InExpression, out expressions.OutExpression, n expressions.NExpression, window Window) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[INPUT] = in
	args[OUTPUT] = out
	args[N_LOWERCASE] = n
	return w.compoundParameterWindowFunction(path, "$bottomN", args, window)
}
func (w windowOutputFields) shift(path string, expression, defaultExpression expressions.TExpression, by int32) WindowOutputField {
	args := make(map[ParamName]expressions.TExpression)
	args[OUTPUT] = expression
	args[BY] = bsonx.NewBsonInt32(by)
	if defaultExpression != nil {
		args[DEFAULT] = defaultExpression
	}
	return w.compoundParameterWindowFunction(path, "$shift", args, nil)
}
func (w windowOutputFields) documentNumber(path string) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$documentNumber", nil, nil)
}
func (w windowOutputFields) rank(path string) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$rank", nil, nil)
}
func (w windowOutputFields) denseRank(path string) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$denseRank", nil, nil)
}
func (w windowOutputFields) locf(path string, expression expressions.TExpression) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$locf", expression, nil)
}
func (w windowOutputFields) linearFill(path string, expression expressions.TExpression) WindowOutputField {
	return w.simpleParameterWindowFunction(path, "$linearFill", expression, nil)
}
func (w windowOutputFields) simpleParameterWindowFunction(path, functionName string, expression expressions.TExpression, window Window) WindowOutputField {
	return NewBsonFieldWindowOutputField(bsonx.NewBsonField(path, NewSimpleParameterFunctionAndWindow[expressions.TExpression](functionName, expression, window)))
}
func (w windowOutputFields) compoundParameterWindowFunction(path, functionName string, args map[ParamName]expressions.TExpression, window Window) WindowOutputField {
	return NewBsonFieldWindowOutputField(bsonx.NewBsonField(path, NewCompoundParameterFunctionAndWindow[expressions.TExpression](functionName, args, window)))
}

type BsonFieldWindowOutputField struct {
	wrapped bsonx.BsonField
}

func NewBsonFieldWindowOutputField(field bsonx.BsonField) BsonFieldWindowOutputField {
	return BsonFieldWindowOutputField{
		field,
	}
}

func (b BsonFieldWindowOutputField) ToBsonField() bsonx.BsonField {
	return b.wrapped
}

type SimpleParameterFunctionAndWindow[T expressions.TExpression] struct {
	AbstractFunctionAndWindow
	expression T
}

func NewSimpleParameterFunctionAndWindow[T expressions.TExpression](functionName string, expression T, window Window) SimpleParameterFunctionAndWindow[T] {
	return SimpleParameterFunctionAndWindow[T]{
		AbstractFunctionAndWindow: NewAbstractFunctionAndWindow(functionName, window),
		expression:                expression,
	}
}

func (a SimpleParameterFunctionAndWindow[T]) ToBsonDocument() bsonx.BsonDocument {
	doc := bsonx.NewEmptyDoc()
	if a.expression != nil {
		doc.Append(a.functionName, a.expression)
	} else {
		doc.Append(a.functionName, bsonx.NewEmptyDoc())
	}
	a.writeWindow(doc)
	return doc
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

func (a AbstractFunctionAndWindow) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewBsonDocument("window", a.window.ToBsonDocument())
}

func (a AbstractFunctionAndWindow) Document() bsonx.Document {
	return a.ToBsonDocument().Document()
}

func (a AbstractFunctionAndWindow) writeWindow(doc bsonx.BsonDocument) {
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
	INPUT       = NewParamName("input")
	UNIT        = NewParamName("unit")
	N_UPPERCASE = NewParamName("N")
	N_LOWERCASE = NewParamName("n")
	ALPHA       = NewParamName("alpha")
	OUTPUT      = NewParamName("output")
	BY          = NewParamName("by")
	DEFAULT     = NewParamName("default")
	SORT_BY     = NewParamName("sortBy")
)

type CompoundParameterFunctionAndWindow[T expressions.TExpression] struct {
	AbstractFunctionAndWindow
	args map[ParamName]T
}

func NewCompoundParameterFunctionAndWindow[T expressions.TExpression](functionName string, args map[ParamName]T, window Window) CompoundParameterFunctionAndWindow[T] {
	return CompoundParameterFunctionAndWindow[T]{
		AbstractFunctionAndWindow: NewAbstractFunctionAndWindow(functionName, window),
		args:                      args,
	}
}

func (a CompoundParameterFunctionAndWindow[T]) ToBsonDocument() bsonx.BsonDocument {
	doc := bsonx.NewEmptyDoc()
	args := bsonx.NewEmptyDoc()
	for name, value := range a.args {
		args.Append(name.value, value)
	}

	a.writeWindow(doc)
	return doc
}
