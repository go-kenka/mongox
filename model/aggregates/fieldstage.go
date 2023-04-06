package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"go.mongodb.org/mongo-driver/bson"
)

type FieldsStage Stage

// AddFields Adds new fields to documents. $addFields outputs documents that
// contain all existing fields from the input documents and newly added fields.
// The $addFields stage is equivalent to a $project stage that explicitly
// specifies all existing fields in the input documents and adds the new fields.
// NOTE Starting in version 4.2, MongoDB adds a new aggregation pipeline stage
// $set that is an alias for $addFields. $addFields has the following form: {
// $addFields: { <newField>: <expression>, ... } } Specify the name of each
// field to add and set its value to an aggregation expression
func AddFields[T expression.AnyExpression](fields ...Field[T]) FieldsStage {
	return NewFieldsStage("$addFields", fields)
}

// Set NewStage in version 4.2. Adds new fields to documents. $set outputs documents
// that contain all existing fields from the input documents and newly added
// fields. The $set stage is an alias for $addFields. Both stages are equivalent
// to a $project stage that explicitly specifies all existing fields in the
// input documents and adds the new fields. $set has the following form: { $set:
// { <newField>: <expression>, ... } } Specify the name of each field to add and
// set its value to an aggregation expression.
func Set[T expression.AnyExpression](fields ...Field[T]) FieldsStage {
	return NewFieldsStage("$set", fields)
}

// UnSet NewStage in version 4.2. Removes/excludes fields from documents. The $unset stage
// has the following syntax: To remove a single field, the $unset takes a string
// that specifies the field to remove: { $unset: "<field>" } To remove multiple
// fields, the $unset takes an array of fields to remove. { $unset: [
// "<field1>", "<field2>", ... ] }
func UnSet(fields ...string) FieldsStage {
	if len(fields) == 1 {
		return NewStage(bsonx.BsonDoc("$unset", bsonx.String(fields[0])))
	}

	array := bsonx.Array()
	for _, field := range fields {
		array.Append(bsonx.String(field))
	}

	return NewStage(bsonx.BsonDoc("$unset", array))
}

type fieldsStage[T expression.AnyExpression] struct {
	fields    []Field[T]
	stageName string
}

func (f fieldsStage[T]) Bson() bsonx.Bson {
	return f.Pro()
}

func NewFieldsStage[T expression.AnyExpression](stageName string, fields []Field[T]) fieldsStage[T] {
	return fieldsStage[T]{
		fields:    fields,
		stageName: stageName,
	}
}

func (f fieldsStage[T]) Pro() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonEmpty()
	for _, field := range f.fields {
		data.Append(field.name, field.value)
	}
	b.Append(f.stageName, data)
	return b
}

func (f fieldsStage[T]) Document() bson.D {
	return f.Pro().Document()
}
