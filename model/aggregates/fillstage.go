package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/fill"
	"go.mongodb.org/mongo-driver/bson"
)

type FillStage Stage

// Fill NewStage in version 5.3. Populates null and missing field values within documents.
// You can use $fill to populate missing data points: In a sequence based on
// surrounding values. With a fixed value. The $fill stage has this syntax:
//
//	{
//	  $fill: {
//	     partitionBy: <expression>,
//	     partitionByFields: [ <field 1>, <field 2>, ... , <field n> ],
//	     sortBy: {
//	        <sort field 1>: <sort order>,
//	        <sort field 2>: <sort order>,
//	        ...,
//	        <sort field n>: <sort order>
//	     },
//	     output: {
//	        <field 1>: { value: <expression> },
//	        <field 2>: { method: <string> },
//	        ...
//	     }
//	  }
//	}
//
// The $fill stage takes a document with these fields:
func Fill(options fill.FillOptions, output []fill.FillOutputField) FillStage {
	return NewFillStage(options, output)
}

type fillStage struct {
	output  []fill.FillOutputField
	options fill.FillOptions
}

func (f fillStage) Bson() bsonx.Bson {
	return f.Pro()
}

func NewFillStage(
	options fill.FillOptions,
	output []fill.FillOutputField,
) fillStage {
	return fillStage{
		output:  output,
		options: options,
	}
}

func (f fillStage) Pro() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	doc = bsonx.NewMerged(doc, f.options.Pro())
	outputDoc := bsonx.BsonEmpty()
	for _, computation := range f.output {
		computationDoc := computation.Pro()
		if computationDoc.Size() == 1 {
			outputDoc = bsonx.NewMerged(outputDoc, computationDoc)
		}
	}

	doc.Append("output", outputDoc)
	return bsonx.BsonDoc("$fill", doc)
}
func (f fillStage) Document() bson.D {
	return f.Pro().Document()
}
