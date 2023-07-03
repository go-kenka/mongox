package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/internal/options"
	"go.mongodb.org/mongo-driver/bson"
)

type LookupStage struct {
	stage bsonx.Bson
}

func (s LookupStage) Bson() bsonx.Bson {
	return s.stage
}

func (s LookupStage) Document() bson.D {
	return s.stage.Document()
}

// Lookup Changed in version 5.1. Performs a left outer join to a collection in the
// same Database to filter in documents from the "joined" collection for
// processing. The $lookup DefaultStage adds a new array field to each input document.
// The new array field contains the matching documents from the "joined"
// collection. The $lookup DefaultStage passes these reshaped documents to the next
// DefaultStage. Starting in MongoDB 5.1, $lookup works across sharded collections. To
// combine elements from two different collections, use the $unionWith pipeline
// DefaultStage. The $lookup DefaultStage has the following syntaxes: Equality Match with a
// Single Join Condition To perform an equality match between a field from the
// input documents with a field from the documents of the "joined" collection,
// the $lookup DefaultStage has this syntax:
//
//	{
//	  $lookup:
//	    {
//	      from: <collection to join>,
//	      localField: <field from the input documents>,
//	      foreignField: <field from the documents of the "from" collection>,
//	      as: <output array field>
//	    }
//	}
func Lookup(from, localField, foreignField, as string) LookupStage {
	return LookupStage{stage: bsonx.BsonDoc("$lookup",
		bsonx.BsonDoc("from", bsonx.String(from)).
			Append("localField", bsonx.String(localField)).
			Append("foreignField", bsonx.String(foreignField)).
			Append("as", bsonx.String(as)),
	)}
}

// LookupWithPipe Changed in version 5.1. Performs a left outer join to a collection in the
// same Database to filter in documents from the "joined" collection for
// processing. The $lookup DefaultStage adds a new array field to each input document.
// The new array field contains the matching documents from the "joined"
// collection. The $lookup DefaultStage passes these reshaped documents to the next
// DefaultStage. Starting in MongoDB 5.1, $lookup works across sharded collections. To
// combine elements from two different collections, use the $unionWith pipeline
// DefaultStage. Join Conditions and Subqueries on a Joined Collection MongoDB
// supports: Executing a pipeline on a joined collection. Multiple join
// conditions. Correlated and uncorrelated subqueries. In MongoDB, a correlated
// subquery is a pipeline in a $lookup DefaultStage that references document fields
// from a joined collection. An uncorrelated subquery does not reference joined
// fields. MongoDB correlated subqueries are comparable to SQL correlated
// subqueries, where the inner query references outer query values. An SQL
// uncorrelated subquery does not reference outer query values. MongoDB 5.0 also
// supports concise correlated subqueries. To perform correlated and
// uncorrelated subqueries with two collections, and perform other join
// conditions besides a single equality match, use this $lookup syntax:
//
//	{
//	  $lookup:
//	     {
//	        from: <joined collection>,
//	        let: { <var_1>: <expression>, …, <var_n>: <expression> },
//	        pipeline: [ <pipeline to run on joined collection> ],
//	        as: <output array field>
//	     }
//	}
func LookupWithPipe[T expression.AnyExpression](from, as string, let []options.Variable[T], pipeline []bsonx.Bson) LookupStage {
	return LookupStage{stage: NewLookupStage(from, let, pipeline, as)}
}

type lookupStage[T expression.AnyExpression] struct {
	from     string
	let      []options.Variable[T]
	pipeline []bsonx.Bson
	as       string
}

func NewLookupStage[T expression.AnyExpression](
	from string,
	let []options.Variable[T],
	pipeline []bsonx.Bson,
	as string,
) lookupStage[T] {
	return lookupStage[T]{
		from:     from,
		let:      let,
		pipeline: pipeline,
		as:       as,
	}
}

func (bs lookupStage[T]) BsonDocument() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonDoc("form", bsonx.String(bs.from))

	if len(bs.let) > 0 {
		let := bsonx.BsonEmpty()
		for _, field := range bs.let {
			let.Append(field.GetName(), field.GetValue())
		}
		data.Append("let", let)
	}

	if len(bs.pipeline) > 0 {
		pipeline := bsonx.Array()
		for _, p := range bs.pipeline {
			pipeline.Append(p.BsonDocument())
		}
		data.Append("pipeline", pipeline)
	}
	data.Append("as", bsonx.String(bs.as))
	b.Append("$lookup", data)
	return b
}

func (bs lookupStage[T]) Document() bson.D {
	return bs.BsonDocument().Document()
}
