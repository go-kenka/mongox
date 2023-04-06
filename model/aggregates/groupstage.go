package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/model/aggregates/operators/accumulator"
	"go.mongodb.org/mongo-driver/bson"
)

type GroupStage Stage

// Group The $group stage separates documents into groups according to a "group key".
// The output is one document for each unique group key. A group key is often a
// field, or group of fields. The group key can also be the result of an
// expression. Use the _id field in the $group pipeline stage to set the group
// key. See below for usage examples. In the $group stage output, the _id field
// is set to the group key for that document. The output documents can also
// contain additional fields that are set using accumulator expressions. The
// $group stage has the following prototype form:
//
//	{
//	 $group:
//	   {
//	     _id: <expression>, // Group key
//	     <field1>: { <accumulator1> : <expression1> },
//	     ...
//	   }
//	}
func Group[T expression.AnyExpression](id T, fieldAccumulators ...accumulator.AccumulatorBson) GroupStage {
	return NewGroupStage(id, fieldAccumulators)
}

type groupStage[T expression.AnyExpression] struct {
	id                T
	fieldAccumulators []accumulator.AccumulatorBson
}

func (bs groupStage[T]) Bson() bsonx.Bson {
	return bs.Pro()
}

func NewGroupStage[T expression.AnyExpression](id T, fieldAccumulators []accumulator.AccumulatorBson) groupStage[T] {
	return groupStage[T]{
		id:                id,
		fieldAccumulators: fieldAccumulators,
	}
}

func (bs groupStage[T]) Pro() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonDoc("_id", bs.id)

	if len(bs.fieldAccumulators) > 0 {
		for _, field := range bs.fieldAccumulators {
			data = bsonx.NewMerged(data, field.Pro())
		}
	}

	b.Append("$group", data)
	return b
}

func (bs groupStage[T]) Document() bson.D {
	return bs.Pro().Document()
}
