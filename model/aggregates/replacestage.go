package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/bsonx/expression"
	"go.mongodb.org/mongo-driver/bson"
)

type ReplaceStage struct {
	stage bsonx.Bson
}

func (s ReplaceStage) Bson() bsonx.Bson {
	return s.stage
}

func (s ReplaceStage) Document() bson.D {
	return s.stage.Document()
}
func (s ReplaceStage) Watch() {
}
func (s ReplaceStage) Update() {
}

// ReplaceRoot Replaces the input document with the specified document. The operation
// replaces all existing fields in the input document, including the _id field.
// You can promote an existing embedded document to the top level, or create a
// new document for promotion (see example ). NOTE Starting in version 4.2,
// MongoDB adds a new aggregation pipeline DefaultStage $replaceWith that is an alias
// for $replaceRoot. The $replaceRoot DefaultStage has the following form: {
// $replaceRoot: { newRoot: <replacementDocument> } } The replacement document
// can be any valid expression that resolves to a document. The DefaultStage errors and
// fails if <replacementDocument> is not a document
func ReplaceRoot[T expression.AnyExpression](value T) ReplaceStage {
	return ReplaceStage{stage: NewReplaceStage(value, false)}
}

// ReplaceWith NewDefaultStage in version 4.2. Replaces the input document with the
// specified document. The operation replaces all existing fields in the input
// document, including the _id field. With $replaceWith , you can promote an
// embedded document to the top-level. You can also specify a new document as
// the replacement. The $replaceWith is an alias for $replaceRoot. The
// $replaceWith DefaultStage has the following form: { $replaceWith:
// <replacementDocument> } The replacement document can be any valid expression
// that resolves to a document.
func ReplaceWith[T expression.AnyExpression](value T) ReplaceStage {
	return ReplaceStage{stage: NewReplaceStage(value, true)}
}

type replaceStage[T expression.AnyExpression] struct {
	value       T
	replaceWith bool
}

func NewReplaceStage[T expression.AnyExpression](value T, replaceWith bool) replaceStage[T] {
	return replaceStage[T]{
		value:       value,
		replaceWith: replaceWith,
	}
}

func (f replaceStage[T]) BsonDocument() *bsonx.BsonDocument {
	if f.replaceWith {
		return bsonx.BsonDoc("$replaceWith", f.value)
	}
	b := bsonx.BsonEmpty()
	data := bsonx.BsonDoc("newRoot", f.value)
	b.Append("$replaceRoot", data)
	return b
}

func (f replaceStage[T]) Document() bson.D {
	return f.BsonDocument().Document()
}
