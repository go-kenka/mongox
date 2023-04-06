package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/model/aggregates/operators/window"
	"go.mongodb.org/mongo-driver/bson"
)

type SetWindowFieldsStage Stage

// SetWindowFields NewStage in version 5.0. Performs operations on a specified span
// of documents in a collection, known as a window, and returns the results
// based on the chosen window operator. For example, you can use the
// $setWindowFields stage to output the: Difference in sales between two
// documents in a collection. Sales rankings. Cumulative sales totals. Analysis
// of complex time series information without exporting the data to an external
// database. Syntax The $setWindowFields stage syntax:
//
//	{
//	  $setWindowFields: {
//	     partitionBy: <expression>,
//	     sortBy: {
//	        <sort field 1>: <sort order>,
//	        <sort field 2>: <sort order>,
//	        ...,
//	        <sort field n>: <sort order>
//	     },
//	     output: {
//	        <output field 1>: {
//	           <window operator>: <window operator parameters>,
//	           window: {
//	              documents: [ <lower boundary>, <upper boundary> ],
//	              range: [ <lower boundary>, <upper boundary> ],
//	              unit: <time unit>
//	           }
//	        },
//	        <output field 2>: { ... },
//	        ...
//	        <output field n>: { ... }
//	     }
//	  }
//	}
func SetWindowFields[T expression.AnyExpression, O window.WindowOutputField](partitionBy T, sortBy bsonx.Bson, output []O) SetWindowFieldsStage {
	return NewSetWindowFieldsStage(partitionBy, sortBy, output)
}

type setWindowFieldsStage[T expression.AnyExpression, O window.WindowOutputField] struct {
	partitionBy T
	sortBy      bsonx.Bson
	output      []O
}

func (f setWindowFieldsStage[T, O]) Bson() bsonx.Bson {
	return f.Pro()
}

func NewSetWindowFieldsStage[T expression.AnyExpression, O window.WindowOutputField](
	partitionBy T,
	sortBy bsonx.Bson,
	output []O,
) setWindowFieldsStage[T, O] {
	return setWindowFieldsStage[T, O]{
		partitionBy: partitionBy,
		sortBy:      sortBy,
		output:      output,
	}
}

func (f setWindowFieldsStage[T, O]) Pro() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonEmpty()
	if f.partitionBy != nil {
		data.Append("partitionBy", f.partitionBy)
	}
	if f.sortBy != nil {
		data.Append("sortBy", f.sortBy.Pro())
	}
	output := bsonx.BsonEmpty()
	for _, s := range f.output {
		field := s.Exp()
		output = bsonx.NewMerged(output, field.AsDocument())
	}
	data.Append("output", output)
	b.Append("$setWindowFields", data)
	return b
}

func (f setWindowFieldsStage[T, O]) Document() bson.D {
	return f.Pro().Document()
}
