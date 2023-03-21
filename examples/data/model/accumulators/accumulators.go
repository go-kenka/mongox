package accumulators

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/expressions"
)

func Sum(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$sum", fieldName, expr)
}
func Avg(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$avg", fieldName, expr)
}
func First(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$first", fieldName, expr)
}
func FirstN(fieldName string, in expressions.InExpression, n expressions.NExpression) bsonx.BsonField {
	return pickNAccumulator(fieldName, "$firstN", in, n)
}
func Top(fieldName string, sortBy bsonx.Bson, out expressions.OutExpression) bsonx.BsonField {
	return sortingPickAccumulator(fieldName, "$top", sortBy, out)
}
func TopN(fieldName string, sortBy bsonx.Bson, out expressions.OutExpression, n expressions.NExpression) bsonx.BsonField {
	return sortingPickNAccumulator(fieldName, "$top", sortBy, out, n)
}
func Last(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$top", fieldName, expr)
}
func LastN(fieldName string, in expressions.InExpression, n expressions.NExpression) bsonx.BsonField {
	return pickNAccumulator(fieldName, "$lastN", in, n)
}
func Bottom(fieldName string, sortBy bsonx.Bson, out expressions.OutExpression) bsonx.BsonField {
	return sortingPickAccumulator(fieldName, "$bottom", sortBy, out)
}
func BottomN(fieldName string, sortBy bsonx.Bson, out expressions.OutExpression, n expressions.NExpression) bsonx.BsonField {
	return sortingPickNAccumulator(fieldName, "$bottom", sortBy, out, n)
}
func Max(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$max", fieldName, expr)
}
func MaxN(fieldName string, in expressions.InExpression, n expressions.NExpression) bsonx.BsonField {
	return pickNAccumulator(fieldName, "$maxN", in, n)
}
func Min(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$min", fieldName, expr)
}
func MinN(fieldName string, in expressions.InExpression, n expressions.NExpression) bsonx.BsonField {
	return pickNAccumulator(fieldName, "$minN", in, n)
}
func Push(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$push", fieldName, expr)
}
func AddToSet(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$addToSet", fieldName, expr)
}
func MergeObjects(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$mergeObjects", fieldName, expr)
}
func StdDevPop(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$stdDevPop", fieldName, expr)
}
func StdDevSamp(fieldName string, expr expressions.TExpression) bsonx.BsonField {
	return accumulatorOperator("$stdDevSamp", fieldName, expr)
}
func Accumulator(
	fieldName, initFunction string,
	initArgs []string, accumulateFunction string,
	accumulateArgs []string, mergeFunction string,
	finalizeFunction string, lang string,
) bsonx.BsonField {
	accumulatorStage := bsonx.NewBsonDocument("init", bsonx.NewBsonString(initFunction))
	iArgs := bsonx.NewBsonArray()
	for _, arg := range initArgs {
		iArgs.Append(bsonx.NewBsonString(arg))
	}
	accumulatorStage.Append("initArgs", iArgs)
	accumulatorStage.Append("accumulate", bsonx.NewBsonString(accumulateFunction))
	aArgs := bsonx.NewBsonArray()
	for _, arg := range accumulateArgs {
		aArgs.Append(bsonx.NewBsonString(arg))
	}
	accumulatorStage.Append("accumulateArgs", aArgs)
	accumulatorStage.Append("merge", bsonx.NewBsonString(mergeFunction))
	accumulatorStage.Append("lang", bsonx.NewBsonString(lang))
	if len(finalizeFunction) > 0 {
		accumulatorStage.Append("finalize", bsonx.NewBsonString(finalizeFunction))
	}
	return accumulatorOperator[bsonx.BsonDocument]("$accumulator", fieldName, accumulatorStage)
}
func accumulatorOperator[T expressions.TExpression](name, fieldName string, expr T) bsonx.BsonField {
	return bsonx.NewBsonField(fieldName, expressions.NewSimpleExpression(name, expr))
}
func pickNAccumulator[I expressions.InExpression, N expressions.NExpression](fieldName, accumulatorName string, in I, n N) bsonx.BsonField {
	return bsonx.NewBsonField(fieldName, bsonx.NewDocument(accumulatorName, bsonx.NewDocument("input", in).Append("n", n)))
}
func sortingPickAccumulator[O expressions.OutExpression](fieldName, accumulatorName string, sort bsonx.Bson, out O) bsonx.BsonField {
	return bsonx.NewBsonField(fieldName, bsonx.NewDocument(accumulatorName,
		bsonx.NewDocument("sortBy", sort).Append("output", out),
	))
}
func sortingPickNAccumulator[O expressions.OutExpression, N expressions.NExpression](fieldName, accumulatorName string, sort bsonx.Bson, out O, n N) bsonx.BsonField {
	return bsonx.NewBsonField(fieldName, bsonx.NewDocument(accumulatorName,
		bsonx.NewDocument("sortBy", sort).
			Append("output", out).
			Append("n", n),
	))
}
