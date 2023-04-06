// Package accumulator Custom Aggregation Expression Operators
package accumulator

import (
	"github.com/go-kenka/mongox/bsonx"
)

type CustomOperator struct {
	doc bsonx.Bson
}

func (o CustomOperator) Exp() bsonx.IBsonValue {
	return o.doc.Pro()
}

// Accumulator New in version 4.4. Defines a custom accumulator operator.
// Accumulators are operators that maintain their state (e.g. totals, maximums,
// minimums, and related data) as documents progress through the pipeline. Use
// the $accumulator operator to execute your own JavaScript functions to
// implement behavior not supported by the MongoDB Query Language. See also
// $function. $accumulator is available in these stages: $bucket $bucketAuto
// $group The $accumulator operator has this syntax:
//
//	{
//	 $accumulator: {
//	   init: <code>,
//	   initArgs: <array expression>,        // Optional
//	   accumulate: <code>,
//	   accumulateArgs: <array expression>,
//	   merge: <code>,
//	   finalize: <code>,                    // Optional
//	   lang: <string>
//	 }
//	}
func Accumulator(
	initFunction string,
	initArgs []string, accumulateFunction string,
	accumulateArgs []string, mergeFunction string,
	finalizeFunction string, lang string,
) CustomOperator {
	accumulatorStage := bsonx.BsonDoc("init", bsonx.String(initFunction))
	iArgs := bsonx.Array()
	for _, arg := range initArgs {
		iArgs.Append(bsonx.String(arg))
	}
	accumulatorStage.Append("initArgs", iArgs)
	accumulatorStage.Append("accumulate", bsonx.String(accumulateFunction))
	aArgs := bsonx.Array()
	for _, arg := range accumulateArgs {
		aArgs.Append(bsonx.String(arg))
	}
	accumulatorStage.Append("accumulateArgs", aArgs)
	accumulatorStage.Append("merge", bsonx.String(mergeFunction))
	accumulatorStage.Append("lang", bsonx.String(lang))
	if len(finalizeFunction) > 0 {
		accumulatorStage.Append("finalize", bsonx.String(finalizeFunction))
	}
	return CustomOperator{
		doc: bsonx.BsonDoc("$accumulator", accumulatorStage),
	}
}

// Function New in version 4.4. Defines a custom aggregation function or
// expression in JavaScript. You can use the $function operator to define custom
// functions to implement behavior not supported by the MongoDB Query Language.
// See also $accumulator. The $function operator has the following syntax:
//
//	{
//	 $function: {
//	   body: <code>, args: <array expression>, lang: "js"
//	 }
//	}
func Function(body string, initArgs []string, lang string) CustomOperator {
	accumulatorStage := bsonx.BsonDoc("body", bsonx.String(body))
	iArgs := bsonx.Array()
	for _, arg := range initArgs {
		iArgs.Append(bsonx.String(arg))
	}
	accumulatorStage.Append("lang", bsonx.String(lang))
	accumulatorStage.Append("args", iArgs)
	return CustomOperator{
		doc: bsonx.BsonDoc("$function", accumulatorStage),
	}
}
