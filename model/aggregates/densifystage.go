package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/densify"
	"go.mongodb.org/mongo-driver/bson"
)

type DensifyStage Stage

// Densify NewStage in version 5.1. Creates new documents in a sequence of documents
// where certain values in a field are missing. You can use $densify to: Fill
// gaps in time series data. Add missing values between groups of data. Populate
// your data with a specified range of values. The $densify stage has this
// syntax:
//
//	{
//	  $densify: {
//	     field: <fieldName>,
//	     partitionByFields: [ <field 1>, <field 2> ... <field n> ],
//	     range: {
//	        step: <number>,
//	        unit: <time unit>,
//	        bounds: < "full" || "partition" > || [ < lower bound >, < upper bound > ]
//	     }
//	  }
//	}
func Densify(field string, dRange densify.DensifyRange) DensifyStage {
	return NewDensifyStage(field, dRange, densify.DefaultDensifyOptions)
}

type densifyStage struct {
	field        string
	densifyRange densify.DensifyRange
	options      densify.DensifyOptions
}

func (f densifyStage) Bson() bsonx.Bson {
	return f.ToBsonDocument()
}

func NewDensifyStage(
	field string,
	densifyRange densify.DensifyRange,
	options densify.DensifyOptions,
) densifyStage {
	return densifyStage{
		field:        field,
		densifyRange: densifyRange,
		options:      options,
	}
}

func (f densifyStage) ToBsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonDoc("field", bsonx.String(f.field))
	doc.Append("range", f.densifyRange.ToBsonDocument())
	return bsonx.NewMerged(doc, f.options.ToBsonDocument())
}

func (f densifyStage) Document() bson.D {
	return f.ToBsonDocument().Document()
}
