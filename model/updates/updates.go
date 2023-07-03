package updates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/model/filters"
	"go.mongodb.org/mongo-driver/bson"
)

type Update interface {
	Document() any
	Update()
}

type updateStage struct {
	update bsonx.Bson
}

func (s updateStage) Value() bsonx.IBsonValue {
	return s.update.BsonDocument()
}

func (s updateStage) Document() any {
	return s.update.Document()
}

func (s updateStage) Update() {
}

// Combine combine multiple different update contents
func Combine(updates ...updateStage) Update {
	return updateStage{update: NewCompositeUpdate(updates)}
}

// Set The $set operator replaces the value of a field with the specified value. The
// $set operator expression has the following form: { $set: { <field1>:
// <value1>, ... } }
// To specify a <field> in an embedded document or in an array, use dot notation.
func Set[I expression.AnyExpression](fieldName string, value I) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, value, "$set")}
}

// UnSet The $unset operator deletes a particular field. Consider the following
// syntax: { $unset: { <field1>: "", ... } } The specified value in the $unset
// expression (i.e. "") does not impact the operation. To specify a <field> in
// an embedded document or in an array, use dot notation.
func UnSet(fieldName string) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, bsonx.String(""), "$unset")}
}

// SetOnInsert If an update operation with upsert: true results in an insert of a document,
// then $setOnInsert assigns the specified values to the fields in the document.
// If the update operation does not result in an insert, $setOnInsert does
// nothing.
// db.collection.updateOne(
//
//	<query>,
//	{ $setOnInsert: { <field1>: <value1>, ... } },
//	{ upsert: true }
//
// )
func SetOnInsert[I expression.AnyExpression](fieldName string, value I) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, value, "$setOnInsert")}
}

// Rename The $rename operator updates the name of a field and has the following form:
// {$rename: { <field1>: <newName1>, <field2>: <newName2>, ... } } The new field
// name must differ from the existing field name. To specify a <field> in an
// embedded document, use dot notation. Consider the following example:
// db.students.updateOne(
//
//	{ _id: 1 },
//	{ $rename: { 'nickname': 'alias', 'cell': 'mobile' } }
//
// ) This operation renames the field nickname to alias, and the field cell to
// mobile.
func Rename(fieldName, newFieldName string) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, bsonx.String(newFieldName), "$rename")}
}

// Inc The $inc operator increments a field by a specified value and has the
// following form: { $inc: { <field1>: <amount1>, <field2>: <amount2>, ... } }
// To specify a <field> in an embedded document or in an array, use dot
// notation.
func Inc(fieldName string, number int32) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, bsonx.Int32(number), "$inc")}
}

// Mul Multiply the value of a field by a number. To specify a $mul expression, use
// the following prototype: { $mul: { <field1>: <number1>, ... } } The field to
// update must contain a numeric value. To specify a <field> in an embedded
// document or in an array, use dot notation.
func Mul(fieldName string, number int32) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, bsonx.Int32(number), "$mul")}
}

// Min The $min updates the value of the field to a specified value if the specified
// value is less than the current value of the field. The $min operator can
// compare values of different types, using the BSON comparison order. { $min: {
// <field1>: <value1>, ... } }
// To specify a <field> in an embedded document or in an array, use dot notation.
func Min[I expression.AnyExpression](fieldName string, value I) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, value, "$min")}
}

// Max The $max operator updates the value of the field to a specified value if the
// specified value is greater than the current value of the field. The $max
// operator can compare values of different types, using the BSON comparison
// order. The $max operator expression has the form: { $max: { <field1>:
// <value1>, ... } } To specify a <field> in an embedded document or in an
// array, use dot notation.
func Max[I expression.AnyExpression](fieldName string, value I) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, value, "$max")}
}

// CurrentDate The $currentDate operator sets the value of a field to the
// current date, either as a Date or a timestamp. The default type is Date. The
// $currentDate operator has the form: { $currentDate: { <field1>:
// <typeSpecification1>, ... } } <typeSpecification> can be either: a boolean
// true to set the field value to the current date as a Date, or a document {
// $type: "timestamp" } or { $type: "date" } which explicitly specifies the
// type. The operator is case-sensitive and accepts only the lowercase
// "timestamp" or the lowercase "date". To specify a <field> in an embedded
// document or in an array, use dot notation.
func CurrentDate(fieldName string) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, bsonx.Boolean(true), "$currentDate")}
}

// CurrentTimestamp The $currentDate operator sets the value of a field to the
// current date, either as a Date or a timestamp. The default type is Date. The
// $currentDate operator has the form: { $currentDate: { <field1>:
// <typeSpecification1>, ... } } <typeSpecification> can be either: a boolean
// true to set the field value to the current date as a Date, or a document {
// $type: "timestamp" } or { $type: "date" } which explicitly specifies the
// type. The operator is case-sensitive and accepts only the lowercase
// "timestamp" or the lowercase "date". To specify a <field> in an embedded
// document or in an array, use dot notation.
func CurrentTimestamp(fieldName string) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, bsonx.BsonDoc("$type", bsonx.String("timestamp")), "$currentDate")}
}

// AddToSet The $addToSet operator adds a value to an array unless the value is already
// present, in which case $addToSet does nothing to that array. The $addToSet
// operator has the form: { $addToSet: { <field1>: <value1>, ... } } To specify
// a <field> in an embedded document or in an array, use dot notation.
func AddToSet[I expression.AnyExpression](fieldName string, value I) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, value, "$addToSet")}
}

// AddEachToSet $addToSet with $each,The $each modifier is available for use
// with the $addToSet operator and the $push operator. Use with the $addToSet
// operator to add multiple values to an array <field> if the values do not
// exist in the <field>. { $addToSet: { <field>: { $each: [ <value1>, <value2>
// ... ] } } }
func AddEachToSet[I expression.AnyExpression](fieldName string, values []I) updateStage {
	return updateStage{update: NewWithEachUpdate(fieldName, values, "$addToSet")}
}

// Push The $push operator appends a specified value to an array. The $push operator
// has the form: { $push: { <field1>: <value1>, ... } } To specify a <field> in
// an embedded document or in an array, use dot notation.
func Push[I expression.AnyExpression](fieldName string, value I) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, value, "$push")}
}

// PushEach Use with the $push operator to append multiple values to an array <field>.
// { $push: { <field>: { $each: [ <value1>, <value2> ... ] } } }
// The $push operator can use
// $each
// modifier with other modifiers. For a list of modifiers available for $push, see Modifiers.
func PushEach[I expression.AnyExpression](fieldName string, values []I, options PushOptions) updateStage {
	return updateStage{update: NewPushUpdate(fieldName, values, options)}
}

// Pull The $pull operator removes from an existing array all instances of a value or
// values that match a specified condition. The $pull operator has the form: {
// $pull: { <field1>: <value|condition>, <field2>: <value|condition>, ... } } To
// specify a <field> in an embedded document or in an array, use dot notation.
func Pull[I expression.AnyExpression](fieldName string, value I) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, value, "$pull")}
}

// PullByFilter The $pull operator removes from an existing array all instances of a value or
// values that match a specified condition. The $pull operator has the form: {
// $pull: { <field1>: <value|condition>, <field2>: <value|condition>, ... } } To
// specify a <field> in an embedded document or in an array, use dot notation.
func PullByFilter(filter filters.Filter) updateStage {
	return updateStage{update: bsonx.BsonDoc("$pull", filter.Value())}
}

// PullAll The $pullAll operator removes all instances of the specified values from an
// existing array. Unlike the $pull operator that removes elements by specifying
// a query, $pullAll removes elements that match the listed values. The $pullAll
// operator has the form: { $pullAll: { <field1>: [ <value1>, <value2> ... ],
// ... } } To specify a <field> in an embedded document or in an array, use dot
// notation.
func PullAll[I expression.AnyExpression](fieldName string, values []I) updateStage {
	return updateStage{update: NewPullAllUpdate(fieldName, values)}
}

// PopFirst The $pop operator removes the first or last element of an array.
// Pass $pop a value of -1 to remove the first element of an array and 1 to
// remove the last element in an array. The $pop operator has the form: { $pop:
// { <field>: <-1 | 1>, ... } } To specify a <field> in an embedded document or
// in an array, use dot notation.
func PopFirst(fieldName string) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, bsonx.Int32(-1), "$pop")}
}

// PopLast The $pop operator removes the first or last element of an array.
// Pass $pop a value of -1 to remove the first element of an array and 1 to
// remove the last element in an array. The $pop operator has the form: { $pop:
// { <field>: <-1 | 1>, ... } } To specify a <field> in an embedded document or
// in an array, use dot notation.
func PopLast(fieldName string) updateStage {
	return updateStage{update: NewSimpleUpdate(fieldName, bsonx.Int32(1), "$pop")}
}

// BitwiseAnd Use a bitwise and in the updateOne() operation to update expdata.
// db.switches.updateOne(
//
//	{ _id: 1 }, { $bit: { expdata: { and: Int32( 10 ) } } }
//
// )
func BitwiseAnd(fieldName string, value int64) updateStage {
	return updateStage{update: createBitUpdateDocument(fieldName, "and", bsonx.Int64(value))}
}

// BitwiseOr Use a bitwise or in the updateOne() operation to update expdata.
//
// db.switches.updateOne(
//
//	{ _id: 2 },
//	{ $bit: { expdata: { or: Int32( 5 ) } } }
//
// )
func BitwiseOr(fieldName string, value int64) updateStage {
	return updateStage{update: createBitUpdateDocument(fieldName, "or", bsonx.Int64(value))}
}

// BitwiseXor Use a bitwise xor in the updateOne() operation to update expdata.
//
// db.switches.updateOne(
//
//	{ _id: 3 },
//	{ $bit: { expdata: { xor: Int32( 5 ) } } }
//
// )
func BitwiseXor(fieldName string, value int64) updateStage {
	return updateStage{update: createBitUpdateDocument(fieldName, "xor", bsonx.Int64(value))}
}

func createBitUpdateDocument(fieldName, bitwiseOperator string, value bsonx.IBsonValue) *bsonx.BsonDocument {
	return bsonx.BsonDoc("$bit", bsonx.BsonDoc(fieldName, bsonx.BsonDoc(bitwiseOperator, value)))
}

type simpleBsonKeyValue struct {
	fieldName string
	value     bsonx.IBsonValue
}

func NewSimpleBsonKeyValue(fieldName string, value bsonx.IBsonValue) simpleBsonKeyValue {
	return simpleBsonKeyValue{
		fieldName: fieldName,
		value:     value,
	}
}

func (s simpleBsonKeyValue) BsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(s.fieldName, s.value)
}
func (s simpleBsonKeyValue) Document() bson.D {
	return s.BsonDocument().Document()
}

type simpleUpdate[T expression.AnyExpression] struct {
	fieldName string
	value     T
	operator  string
}

func NewSimpleUpdate[T expression.AnyExpression](fieldName string, value T, operator string) simpleUpdate[T] {
	return simpleUpdate[T]{
		fieldName: fieldName,
		value:     value,
		operator:  operator,
	}
}

func (s simpleUpdate[T]) BsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(s.operator, bsonx.BsonDoc(s.fieldName, s.value))
}

func (s simpleUpdate[T]) Document() bson.D {
	return s.BsonDocument().Document()
}

type withEachUpdate[T expression.AnyExpression] struct {
	fieldName string
	values    []T
	operator  string
}

func NewWithEachUpdate[T expression.AnyExpression](fieldName string, values []T, operator string) withEachUpdate[T] {
	return withEachUpdate[T]{
		fieldName: fieldName,
		values:    values,
		operator:  operator,
	}
}

func (w withEachUpdate[T]) writeAdditionalFields(document *bsonx.BsonDocument) {
}
func (w withEachUpdate[T]) additionalFieldsToString() string {
	return ""
}

func (w withEachUpdate[T]) BsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	values := bsonx.Array()
	for _, value := range w.values {
		values.Append(value)
	}
	each := bsonx.BsonDoc("$each", values)
	w.writeAdditionalFields(each)

	return doc.Append(w.operator, each)
}

func (w withEachUpdate[T]) Document() bson.D {
	return w.BsonDocument().Document()
}

type pushUpdate[T expression.AnyExpression] struct {
	withEachUpdate[T]
	options PushOptions
}

func NewPushUpdate[T expression.AnyExpression](fieldName string, values []T, options PushOptions) pushUpdate[T] {
	return pushUpdate[T]{
		withEachUpdate: NewWithEachUpdate(fieldName, values, "$push"),
		options:        options,
	}
}

func (p pushUpdate[T]) writeAdditionalFields(document *bsonx.BsonDocument) {
	if p.options.HasPosition() {
		document.Append("$position", bsonx.Int32(p.options.GetPosition()))
	}
	if p.options.HasSlice() {
		document.Append("$slice", bsonx.Int32(p.options.GetSlice()))
	}
	if p.options.HasSort() {
		document.Append("$sort", bsonx.Int32(p.options.GetSort()))
	} else {
		sortDocument := p.options.GetSortDocument()
		if sortDocument != nil {
			document.Append("$sort", sortDocument.BsonDocument())
		}
	}
}
func (p pushUpdate[T]) additionalFieldsToString() string {
	return ", options=" + p.options.ToString()
}

type pullAllUpdate[T expression.AnyExpression] struct {
	fieldName string
	values    []T
}

func NewPullAllUpdate[T expression.AnyExpression](fieldName string, values []T) pullAllUpdate[T] {
	return pullAllUpdate[T]{
		fieldName: fieldName,
		values:    values,
	}
}

func (p pullAllUpdate[T]) BsonDocument() *bsonx.BsonDocument {
	values := bsonx.Array()
	for _, value := range p.values {
		values.Append(value)
	}
	return bsonx.BsonDoc("$pullAll", bsonx.BsonDoc(p.fieldName, values))
}

func (p pullAllUpdate[T]) Document() bson.D {
	return p.BsonDocument().Document()
}

type compositeUpdate struct {
	updates []updateStage
}

func NewCompositeUpdate(updates []updateStage) compositeUpdate {
	return compositeUpdate{
		updates: updates,
	}
}

func (p compositeUpdate) BsonDocument() *bsonx.BsonDocument {
	document := bsonx.BsonEmpty()
	for _, update := range p.updates {
		rendered := update.Value().AsDocument()
		for _, v := range rendered.Data() {
			if document.ContainsKey(v.Key) {
				if v.Value.IsDocument() {
					currentOperatorDocument := v.Value.AsDocument()
					existingOperatorDocument := document.GetDocument(v.Key)
					for _, cv := range currentOperatorDocument.Data() {
						existingOperatorDocument.Append(cv.Key, cv.Value)
					}
				}
			} else {
				document.Append(v.Key, v.Value)
			}
		}
	}
	return document
}

func (p compositeUpdate) Document() bson.D {
	return p.BsonDocument().Document()
}
