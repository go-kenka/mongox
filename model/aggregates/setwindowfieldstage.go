package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/model/aggregates/operators/window"
	utils "github.com/go-kenka/mongox/uitls"
	"go.mongodb.org/mongo-driver/bson"
)

type SetWindowFieldsStage struct {
	stage bsonx.Bson
}

func (s SetWindowFieldsStage) Bson() bsonx.Bson {
	return s.stage
}

func (s SetWindowFieldsStage) Document() bson.D {
	return s.stage.Document()
}

// SetWindowFields NewDefaultStage in version 5.0. Performs operations on a specified span
// of documents in a collection, known as a window, and returns the results
// based on the chosen window operator. For example, you can use the
// $setWindowFields DefaultStage to output the: Difference in sales between two
// documents in a collection. Sales rankings. Cumulative sales totals. Analysis
// of complex time series information without exporting the data to an external
// Database. Syntax The $setWindowFields DefaultStage syntax:
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
	return SetWindowFieldsStage{stage: NewSetWindowFieldsStage(partitionBy, sortBy, output)}
}

type setWindowFieldsStage[T expression.AnyExpression, O window.WindowOutputField] struct {
	partitionBy T
	sortBy      bsonx.Bson
	output      []O
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

func (f setWindowFieldsStage[T, O]) BsonDocument() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonEmpty()
	if !utils.IsZero(f.partitionBy) {
		data.Append("partitionBy", f.partitionBy)
	}
	if f.sortBy != nil {
		data.Append("sortBy", f.sortBy.BsonDocument())
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
	return f.BsonDocument().Document()
}
