package miscellaneou

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	utils "github.com/go-kenka/mongox/utils"
)

type miscellaneousOperator struct {
	doc bsonx.Bson
}

func (o miscellaneousOperator) Exp() bsonx.IBsonValue {
	return o.doc.BsonDocument()
}

// GetField New in version 5.0.
// Returns the value of a specified field from a document. If you don't specify an object,
// $getField returns the value of the field from $$CURRENT.
// You can use $getField to retrieve the value of fields with names that contain periods (.) or start with dollar signs ($).
// $getField has the following syntax:
//
//	{
//	 $getField: {
//	   field: <String>,
//	   input: <Object>
//	 }
//	}
func GetField[T expression.ObjectExpression](filed string, input T) miscellaneousOperator {
	if !utils.IsZero(input) {
		doc := bsonx.BsonEmpty()
		doc.Append("field", bsonx.String(filed))
		doc.Append("input", input)
		return miscellaneousOperator{doc: bsonx.BsonDoc("$getField", doc)}
	}

	return miscellaneousOperator{doc: bsonx.BsonDoc("$getField", bsonx.String(filed))}
}

// Rand New in version 4.4.2.
// Returns a random float between 0 and 1 each time it is called.
// $rand  has the following syntax:
// { $rand: {} }
// The $rand operator doesn't take any arguments.
func Rand() miscellaneousOperator {
	return miscellaneousOperator{doc: bsonx.BsonDoc("$rand", bsonx.BsonEmpty())}
}

// SampleRate New in version 4.4.2.
// Matches a random selection of input documents. The number of documents selected approximates the sample rate expressed as a percentage of the total number of documents.
// The $sampleRate operator has the following syntax:
// { $sampleRate: <non-negative float> }
func SampleRate(f float64) miscellaneousOperator {
	return miscellaneousOperator{doc: bsonx.BsonDoc("$sampleRate", bsonx.Double(f))}
}
