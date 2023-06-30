package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/bsonx/expression"
	"github.com/go-kenka/mongox/bsonx/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Stage interface {
	Bson() bsonx.Bson
	Document() bson.D
}

type Pipe interface {
	Pipe() mongo.Pipeline
}

type pipe struct {
	stages []Stage
}

func NewPipe(s ...Stage) *pipe {
	return &pipe{
		stages: s,
	}
}

func (p *pipe) Append(stage Stage) *pipe {
	p.stages = append(p.stages, stage)
	return p
}

func (p *pipe) Pipe() mongo.Pipeline {
	pipe := make(mongo.Pipeline, 0)
	for _, stage := range p.stages {
		pipe = append(pipe, stage.Document())
	}
	return pipe
}

type DatabaseStage interface {
	Stage
	Database()
}

type WatchStage interface {
	Stage
	Watch()
}

type UpdateStage interface {
	Stage
	Update()
}

type DefaultStage struct {
	doc bsonx.Bson
}

func NewDefaultStage(doc bsonx.Bson) DefaultStage {
	return DefaultStage{doc: doc}
}

func (s DefaultStage) Bson() bsonx.Bson {
	return s.doc
}

func (s DefaultStage) Document() bson.D {
	return s.doc.Document()
}

// Count Passes a document to the next DefaultStage that contains a count of the number of
// documents input to the DefaultStage. $count has the following prototype form: {
// $count: <string> } <string> is the name of the output field which has the
// count as its value. <string> must be a non-empty string, must not start with
// $ and must not contain the . character.
func Count(field string) Stage {
	return NewDefaultStage(bsonx.BsonDoc("$count", bsonx.String(field)))
}

// Limit Limits the number of documents passed to the next DefaultStage in the pipeline. The
// $limit DefaultStage has the following prototype form: { $limit: <positive 64-bit
// integer> } $limit takes a positive integer that specifies the maximum number
// of documents to pass along.
func Limit(limit int64) Stage {
	return NewDefaultStage(bsonx.BsonDoc("$limit", bsonx.Int64(limit)))
}

// Out Takes the documents returned by the aggregation pipeline and writes them to a
// specified collection. Starting in MongoDB 4.4, you can specify the output
// Database. The $out DefaultStage must be the last DefaultStage in the pipeline. The $out
// operator lets the aggregation framework return result sets of any size.
// Starting in MongoDB 4.4, $out can take a document to specify the output
// Database as well as the output collection: { $out: { db: "<output-db>", coll:
// "<output-collection>" } }
func Out(databaseName, collectionName string) Stage {
	if len(databaseName) == 0 {
		return NewDefaultStage(bsonx.BsonDoc("$out", bsonx.String(collectionName)))
	}
	return NewDefaultStage(bsonx.BsonDoc("$out", bsonx.BsonDoc("db", bsonx.String(databaseName)).
		Append("coll", bsonx.String(collectionName))))
}

// Sample Randomly selects the specified number of documents from the input documents.
// The $sample DefaultStage has the following syntax: { $sample: { size: <positive
// integer N> } } N is the number of documents to randomly select.
func Sample(size int32) Stage {
	return NewDefaultStage(bsonx.BsonDoc("$sample", bsonx.BsonDoc("size", bsonx.Int32(size))))
}

// Skip Skips over the specified number of documents that pass into the DefaultStage and
// passes the remaining documents to the next DefaultStage in the pipeline. The $skip
// DefaultStage has the following prototype form: { $skip: <positive 64-bit integer> }
// $skip takes a positive integer that specifies the maximum number of documents
// to skip.
func Skip(skip int64) Stage {
	return NewDefaultStage(bsonx.BsonDoc("$skip", bsonx.Int64(skip)))
}

// Sort Sorts all input documents and returns them to the pipeline in sorted order.
// The $sort DefaultStage has the following prototype form: { $sort: { <field1>: <sort
// order>, <field2>: <sort order> ... } } $sort takes a document that specifies
// the field(s) to sort by and the respective sort order.
func Sort[T expression.DocumentExpression](sort T) Stage {
	return NewDefaultStage(bsonx.BsonDoc("$sort", sort))
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
		return NewDefaultStage(bsonx.BsonDoc("$unwind", bsonx.String(fieldName)))
	}
	options := bsonx.BsonDoc("path", bsonx.String(fieldName))
	if unwindOptions.HasPreserveNullAndEmptyArrays() {
		options.Append("preserveNullAndEmptyArrays", bsonx.Boolean(unwindOptions.PreserveNullAndEmptyArrays()))
	}
	if unwindOptions.HasIncludeArrayIndex() {
		options.Append("includeArrayIndex", bsonx.String(unwindOptions.IncludeArrayIndex()))
	}
	return NewDefaultStage(bsonx.BsonDoc("$unwind", options))
}
