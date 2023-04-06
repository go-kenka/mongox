package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/internal/options"
	"github.com/go-kenka/mongox/model/filters"
	"go.mongodb.org/mongo-driver/bson"
)

type Stage interface {
	Bson() bsonx.Bson
	Document() bson.D
}

type stage struct {
	doc *bsonx.BsonDocument
}

func NewStage(doc *bsonx.BsonDocument) stage {
	return stage{doc: doc}
}

func (s stage) Bson() bsonx.Bson {
	return s.doc
}

func (s stage) Document() bson.D {
	return s.doc.Document()
}

// Count Passes a document to the next stage that contains a count of the number of
// documents input to the stage. $count has the following prototype form: {
// $count: <string> } <string> is the name of the output field which has the
// count as its value. <string> must be a non-empty string, must not start with
// $ and must not contain the . character.
func Count(field string) Stage {
	return NewStage(bsonx.BsonDoc("$count", bsonx.String(field)))
}

// Limit Limits the number of documents passed to the next stage in the pipeline. The
// $limit stage has the following prototype form: { $limit: <positive 64-bit
// integer> } $limit takes a positive integer that specifies the maximum number
// of documents to pass along.
func Limit(limit int64) Stage {
	return NewStage(bsonx.BsonDoc("$limit", bsonx.Int64(limit)))
}

// Match Filters the documents to pass only the documents that match the
// specified condition(s) to the next pipeline stage. The $match stage has the
// following prototype form: { $match: { <query> } } $match takes a document
// that specifies the query conditions. The query syntax is identical to the
// read operation query syntax; i.e. $match does not accept raw aggregation
// expressions. Instead, use a $expr query expression to include aggregation
// expression in $match.
func Match[T filters.MatchFilter](filter T) Stage {
	return NewStage(bsonx.BsonDoc("$match", filter.Value()))
}

// Out Takes the documents returned by the aggregation pipeline and writes them to a
// specified collection. Starting in MongoDB 4.4, you can specify the output
// database. The $out stage must be the last stage in the pipeline. The $out
// operator lets the aggregation framework return result sets of any size.
// Starting in MongoDB 4.4, $out can take a document to specify the output
// database as well as the output collection: { $out: { db: "<output-db>", coll:
// "<output-collection>" } }
func Out(databaseName, collectionName string) Stage {
	if len(databaseName) == 0 {
		return NewStage(bsonx.BsonDoc("$out", bsonx.String(collectionName)))
	}
	return NewStage(bsonx.BsonDoc("$out", bsonx.BsonDoc("db", bsonx.String(databaseName)).
		Append("coll", bsonx.String(collectionName))))
}

// Sample Randomly selects the specified number of documents from the input documents.
// The $sample stage has the following syntax: { $sample: { size: <positive
// integer N> } } N is the number of documents to randomly select.
func Sample(size int32) Stage {
	return NewStage(bsonx.BsonDoc("$sample", bsonx.BsonDoc("size", bsonx.Int32(size))))
}

// Skip Skips over the specified number of documents that pass into the stage and
// passes the remaining documents to the next stage in the pipeline. The $skip
// stage has the following prototype form: { $skip: <positive 64-bit integer> }
// $skip takes a positive integer that specifies the maximum number of documents
// to skip.
func Skip(skip int64) Stage {
	return NewStage(bsonx.BsonDoc("$skip", bsonx.Int64(skip)))
}

// Sort Sorts all input documents and returns them to the pipeline in sorted order.
// The $sort stage has the following prototype form: { $sort: { <field1>: <sort
// order>, <field2>: <sort order> ... } } $sort takes a document that specifies
// the field(s) to sort by and the respective sort order.
func Sort[T expression.DocumentExpression](sort T) Stage {
	return NewStage(bsonx.BsonDoc("$sort", sort))
}

// Unwind Deconstructs an array field from the input documents to output a document for
// each element. Each output document is the input document with the value of
// the array field replaced by the element. You can pass a field path operand or
// a document operand to unwind an array field. Field Path Operand You can pass
// the array field path to $unwind . When using this syntax, $unwind does not
// output a document if the field value is null, missing, or an empty array.
// { $unwind: <field path> }
// When you specify the field path, prefix the field
// name with a dollar sign $ and enclose in quotes. Document Operand with
// Options You can pass a document to $unwind to specify various behavior
// options.
//
//	{
//	 $unwind:
//	   {
//	     path: <field path>,
//	     includeArrayIndex: <string>,
//	     preserveNullAndEmptyArrays: <boolean>
//	   }
//	}
func Unwind(fieldName string, unwindOptions *options.UnwindOptions) Stage {
	if unwindOptions == nil {
		return NewStage(bsonx.BsonDoc("$unwind", bsonx.String(fieldName)))
	}
	options := bsonx.BsonDoc("path", bsonx.String(fieldName))
	if unwindOptions.HasPreserveNullAndEmptyArrays() {
		options.Append("preserveNullAndEmptyArrays", bsonx.Boolean(unwindOptions.PreserveNullAndEmptyArrays()))
	}
	if unwindOptions.HasIncludeArrayIndex() {
		options.Append("includeArrayIndex", bsonx.String(unwindOptions.IncludeArrayIndex()))
	}
	return NewStage(bsonx.BsonDoc("$unwind", options))
}
