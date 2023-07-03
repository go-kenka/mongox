package accumulator

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type AccumulatorExpression interface {
	expression.Expression
	accumulatorOperator | CustomOperator
}

type AccumulatorBson bsonx.Bson

func New[T AccumulatorExpression](fieldName string, e T) AccumulatorBson {
	return bsonx.BsonDoc(fieldName, e.Exp())
}

type accumulatorOperator struct {
	doc bsonx.Bson
}

func (o accumulatorOperator) Exp() bsonx.IBsonValue {
	return o.doc.BsonDocument()
}

// AddToSet Changed in version 5.0. $addToSet $addToSet returns an array of all
// unique values that results from applying an expression to each document in a
// group. The order of the elements in the returned array is unspecified.
// $addToSet is available in these stages: $bucket $bucketAuto $group
// $setWindowFields (Available starting in MongoDB 5.0)
func AddToSet[T expression.AnyExpression](e T) accumulatorOperator {
	return accumulatorOperator{doc: bsonx.BsonDoc("$addToSet", e)}
}

// Avg Returns the average value of the numeric values. $avg ignores non-numeric
// values. $avg is available in these stages: $addFields (Available starting in
// MongoDB 3.4) $bucket $bucketAuto $group $match stage that includes an $expr
// expression $project $replaceRoot (Available starting in MongoDB 3.4)
// $replaceWith (Available starting in MongoDB 4.2) $set (Available starting in
// MongoDB 4.2) $setWindowFields (Available starting in MongoDB 5.0) In MongoDB
// 3.2 and earlier, $avg is available in the $group stage only.
func Avg[T expression.AnyExpression](e []T) accumulatorOperator {
	if len(e) == 1 {
		return accumulatorOperator{doc: bsonx.BsonDoc("$avg", e[0])}
	}
	doc := bsonx.Array()
	for _, t := range e {
		doc.Append(t)
	}
	return accumulatorOperator{doc: bsonx.BsonDoc("$avg", doc)}
}

// Bottom New in version 5.2. Returns the bottom element within a group according
// to the specified sort order.
//
//	{
//	  $bottom:
//	     {
//	        sortBy: { <field1>: <sort order>, <field2>: <sort order> ... },
//	        output: <expression>
//	     }
//	}
func Bottom[T expression.AnyExpression](sortBy bsonx.Bson, out T) accumulatorOperator {
	return accumulatorOperator{doc: sortingPickAccumulator("$bottom", sortBy, out)}
}

// BottomN New in version 5.2. Returns an aggregation of the bottom n elements within a
// group, according to the specified sort order. If the group contains fewer than
// n elements, $bottomN returns all elements in the group.
func BottomN[T expression.AnyExpression, N expression.IntExpression](sortBy bsonx.Bson, out T, n N) accumulatorOperator {
	return accumulatorOperator{doc: sortingPickNAccumulator("$bottom", sortBy, out, n)}
}

// Count New in version 5.0. $count Returns the number of documents in a group. $count
// is available in these stages: $bucket $bucketAuto $group $setWindowFields
// (Available starting in MongoDB 5.0)
func Count() accumulatorOperator {
	return accumulatorOperator{doc: bsonx.BsonDoc("$count", bsonx.BsonEmpty())}
}

// First Changed in version 5.0. Returns the value that results from applying an
// expression to the first document in a group of documents. Only meaningful when
// documents are in a defined order. $first is available in these stages: $bucket
// $bucketAuto $group $setWindowFields (Available starting in MongoDB 5.0)
func First[T expression.AnyExpression](e T) accumulatorOperator {
	return accumulatorOperator{doc: bsonx.BsonDoc("$first", e)}
}

// FirstN New in version 5.2. Returns an aggregation of the first n elements
// within a group. The elements returned are meaningful only if in a specified
// sort order. If the group contains fewer than n elements, $firstN returns all
// elements in the group.
func FirstN[T expression.AnyExpression, N expression.IntExpression](in T, n N) accumulatorOperator {
	return accumulatorOperator{doc: pickNAccumulator("$firstN", in, n)}
}

// Last Changed in version 5.0. Returns the value that results from applying an
// expression to the last document in a group of documents. Only meaningful when
// documents are in a defined order. $last is available in these stages: $bucket
// $bucketAuto $group $setWindowFields (Available starting in MongoDB 5.0)
func Last[T expression.AnyExpression](e T) accumulatorOperator {
	return accumulatorOperator{doc: bsonx.BsonDoc("$top", e)}
}

// LastN New in version 5.2.
// Returns an aggregation of the last n elements within a group. The elements
// returned are meaningful only if in a specified sort order. If the group
// contains fewer than n elements, $lastN returns all elements in the group.
func LastN[T expression.AnyExpression, N expression.IntExpression](in T, n N) accumulatorOperator {
	return accumulatorOperator{doc: pickNAccumulator("$lastN", in, n)}
}

// Max Returns the maximum value. $max compares both value and type, using the
// specified BSON comparison order for values of different types. $max is
// available in these stages: $addFields (Available starting in MongoDB 3.4)
// $bucket $bucketAuto $group $match stage that includes an $expr expression
// $project $replaceRoot (Available starting in MongoDB 3.4) $replaceWith
// (Available starting in MongoDB 4.2) $set (Available starting in MongoDB 4.2)
// $setWindowFields (Available starting in MongoDB 5.0) In MongoDB 3.2 and
// earlier, $max is available in the $group stage only.
func Max[T expression.AnyExpression](e T) accumulatorOperator {
	return accumulatorOperator{doc: bsonx.BsonDoc("$max", e)}
}

// MaxN New in version 5.2.
// Returns an aggregation of the maxmimum value n elements within a group. If the
// group contains fewer than n elements, $maxN returns all elements in the group.
func MaxN[T expression.AnyExpression, N expression.IntExpression](in T, n N) accumulatorOperator {
	return accumulatorOperator{doc: pickNAccumulator("$maxN", in, n)}
}

// Min Returns the minimum value. $min compares both value and type, using the
// specified BSON comparison order for values of different types. $min is
// available in these stages: $addFields (Available starting in MongoDB 3.4)
// $bucket $bucketAuto $group $match stage that includes an $expr expression
// $project $replaceRoot (Available starting in MongoDB 3.4) $replaceWith
// (Available starting in MongoDB 4.2) $set (Available starting in MongoDB 4.2)
// $setWindowFields (Available starting in MongoDB 5.0) In MongoDB 3.2 and
// earlier, $min is available in the $group stage only.
func Min[T expression.AnyExpression](e T) accumulatorOperator {
	return accumulatorOperator{doc: bsonx.BsonDoc("$min", e)}
}

// Push returns an array of all values that result from applying an expression
// to documents. $push is available in these stages: $bucket $bucketAuto $group
// $setWindowFields (Available starting in MongoDB 5.0)
func Push[T expression.AnyExpression](e T) accumulatorOperator {
	return accumulatorOperator{doc: bsonx.BsonDoc("$push", e)}
}

// StdDevPop Changed in version 5.0. Calculates the population standard deviation
// of the input values. Use if the values encompass the entire population of data
// you want to represent and do not wish to generalize about a larger population.
// $stdDevPop ignores non-numeric values. If the values represent only a sample
// of a population of data from which to generalize about the population, use
// $stdDevSamp instead. $stdDevPop is available in these stages: $addFields
// (Available starting in MongoDB 3.4) $group $match stage that includes an $expr
// expression $project $replaceRoot (Available starting in MongoDB 3.4)
// $replaceWith (Available starting in MongoDB 4.2) $set (Available starting in
// MongoDB 4.2) $setWindowFields (Available starting in MongoDB 5.0)
func StdDevPop[T expression.AnyExpression](e []T) accumulatorOperator {
	if len(e) == 1 {
		return accumulatorOperator{doc: bsonx.BsonDoc("$stdDevPop", e[0])}
	}
	doc := bsonx.Array()
	for _, t := range e {
		doc.Append(t)
	}
	return accumulatorOperator{doc: bsonx.BsonDoc("$stdDevPop", doc)}
}

// StdDevSamp Changed in version 5.0. Calculates the sample standard deviation of the input
// values. Use if the values encompass a sample of a population of data from
// which to generalize about the population. $stdDevSamp ignores non-numeric
// values. If the values represent the entire population of data or you do not
// wish to generalize about a larger population, use $stdDevPop instead.
// $stdDevSamp is available in these stages: $addFields (Available starting in
// MongoDB 3.4) $group $match stage that includes an $expr expression $project
// $replaceRoot (Available starting in MongoDB 3.4) $replaceWith (Available
// starting in MongoDB 4.2) $set (Available starting in MongoDB 4.2)
// $setWindowFields (Available starting in MongoDB 5.0)
func StdDevSamp[T expression.AnyExpression](e []T) accumulatorOperator {
	if len(e) == 1 {
		return accumulatorOperator{doc: bsonx.BsonDoc("$stdDevSamp", e[0])}
	}
	doc := bsonx.Array()
	for _, t := range e {
		doc.Append(t)
	}
	return accumulatorOperator{doc: bsonx.BsonDoc("$stdDevSamp", doc)}
}

// Sum Changed in version 5.0. Calculates and returns the collective sum of numeric
// values. $sum ignores non-numeric values. $sum is available in these stages:
// $addFields (Available starting in MongoDB 3.4) $bucket $bucketAuto $group
// $match stage that includes an $expr expression $project $replaceRoot
// (Available starting in MongoDB 3.4) $replaceWith (Available starting in
// MongoDB 4.2) $set (Available starting in MongoDB 4.2) $setWindowFields
// (Available starting in MongoDB 5.0) In MongoDB 3.2 and earlier, $sum is
// available in the $group stage only.
func Sum[T expression.AnyExpression](e []T) accumulatorOperator {
	if len(e) == 1 {
		return accumulatorOperator{doc: bsonx.BsonDoc("$sum", e[0])}
	}
	doc := bsonx.Array()
	for _, t := range e {
		doc.Append(t)
	}
	return accumulatorOperator{doc: bsonx.BsonDoc("$sum", doc)}
}

// Top New in version 5.2. Returns the top element within a group according to the
// specified sort order.
func Top[T expression.AnyExpression](sortBy bsonx.Bson, out T) accumulatorOperator {
	return accumulatorOperator{doc: sortingPickAccumulator("$top", sortBy, out)}
}

// TopN New in version 5.2.
// Returns an aggregation of the top n elements within a group, according to the
// specified sort order. If the group contains fewer than n elements, $topN
// returns all elements in the group.
func TopN[T expression.AnyExpression, N expression.IntExpression](sortBy bsonx.Bson, out T, n N) accumulatorOperator {
	return accumulatorOperator{doc: sortingPickNAccumulator("$top", sortBy, out, n)}
}

func pickNAccumulator[T expression.AnyExpression, N expression.IntExpression](accumulatorName string, in T, n N) bsonx.Bson {
	return bsonx.BsonDoc(accumulatorName,
		bsonx.BsonDoc("input", in).Append("n", n))
}
func sortingPickAccumulator[T expression.AnyExpression](accumulatorName string, sort bsonx.Bson, out T) bsonx.Bson {
	return bsonx.BsonDoc(accumulatorName,
		bsonx.BsonDoc("sortBy", sort.BsonDocument()).Append("output", out),
	)
}
func sortingPickNAccumulator[T expression.AnyExpression, N expression.IntExpression](accumulatorName string, sort bsonx.Bson, out T, n N) bsonx.Bson {
	return bsonx.BsonDoc(accumulatorName,
		bsonx.BsonDoc("sortBy", sort.BsonDocument()).
			Append("output", out).
			Append("n", n),
	)
}
