package updates

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/expressions"
)

func Combine(updates ...bsonx.Bson) bsonx.Bson {
	return NewCompositeUpdate(updates)
}

func Set(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$set")
}
func UnSet(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.NewBsonString(""), "$unset")
}

func SetOnInsert(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$setOnInsert")
}
func SetOnInsertBson(value bsonx.Bson) bsonx.Bson {
	return NewSimpleBsonKeyValue("$setOnInsert", value)
}
func Rename(fieldName, newFieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.NewBsonString(newFieldName), "$rename")
}
func Inc(fieldName string, number int32) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.NewBsonInt32(number), "$inc")
}
func Mul(fieldName string, number int32) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.NewBsonInt32(number), "$mul")
}
func Min(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$min")
}
func Max(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$max")
}
func CurrentDate(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.NewBsonBoolean(true), "$currentDate")
}
func CurrentTimestamp(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.NewBsonDocument("$type", bsonx.NewBsonString("timestamp")), "$currentDate")
}
func AddToSet(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$addToSet")
}
func AddEachToSet(fieldName string, values []expressions.TItem) bsonx.Bson {
	return NewWithEachUpdate(fieldName, values, "$addToSet")
}

func Push(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$push")
}

func PushEach(fieldName string, values []expressions.TItem, options PushOptions) bsonx.Bson {
	return NewPushUpdate(fieldName, values, options)
}

func Pull(fieldName string, value expressions.TItem) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$pull")
}
func PullByFilter(filter bsonx.Bson) bsonx.Bson {
	return bsonx.NewBsonDocument("$pull", filter.ToBsonDocument())
}
func PullAll(fieldName string, values []expressions.TItem) bsonx.Bson {
	return NewPullAllUpdate(fieldName, values)
}
func PopFirst(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.NewBsonInt32(-1), "$pop")
}

func PopLast(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.NewBsonInt32(1), "$pop")
}
func BitwiseAnd(fieldName string, value int64) bsonx.Bson {
	return createBitUpdateDocument(fieldName, "and", bsonx.NewBsonInt64(value))
}

func BitwiseOr(fieldName string, value int64) bsonx.Bson {
	return createBitUpdateDocument(fieldName, "or", bsonx.NewBsonInt64(value))
}

func BitwiseXor(fieldName string, value int64) bsonx.Bson {
	return createBitUpdateDocument(fieldName, "xor", bsonx.NewBsonInt64(value))
}

func createBitUpdateDocument(fieldName, bitwiseOperator string, value bsonx.IBsonValue) bsonx.Bson {
	return bsonx.NewBsonDocument("$bit", bsonx.NewBsonDocument(fieldName, bsonx.NewBsonDocument(bitwiseOperator, value)))
}

type SimpleBsonKeyValue struct {
	fieldName string
	value     bsonx.Bson
}

func NewSimpleBsonKeyValue(fieldName string, value bsonx.Bson) SimpleBsonKeyValue {
	return SimpleBsonKeyValue{
		fieldName: fieldName,
		value:     value,
	}
}

func (s SimpleBsonKeyValue) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewBsonDocument(s.fieldName, s.value.ToBsonDocument())
}
func (s SimpleBsonKeyValue) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

type SimpleUpdate[T expressions.TItem] struct {
	fieldName string
	value     T
	operator  string
}

func NewSimpleUpdate[T expressions.TItem](fieldName string, value T, operator string) SimpleUpdate[T] {
	return SimpleUpdate[T]{
		fieldName: fieldName,
		value:     value,
		operator:  operator,
	}
}

func (s SimpleUpdate[T]) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewBsonDocument(s.operator, bsonx.NewBsonDocument(s.fieldName, s.value))
}

func (s SimpleUpdate[T]) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

type WithEachUpdate[T expressions.TItem] struct {
	fieldName string
	values    []T
	operator  string
}

func NewWithEachUpdate[T expressions.TItem](fieldName string, values []T, operator string) WithEachUpdate[T] {
	return WithEachUpdate[T]{
		fieldName: fieldName,
		values:    values,
		operator:  operator,
	}
}

func (w WithEachUpdate[T]) writeAdditionalFields(document *bsonx.BsonDocument) {
}
func (w WithEachUpdate[T]) additionalFieldsToString() string {
	return ""
}

func (w WithEachUpdate[T]) ToBsonDocument() bsonx.BsonDocument {
	doc := bsonx.NewEmptyDoc()
	values := bsonx.NewBsonArray()
	for _, value := range w.values {
		values.Append(value)
	}
	each := bsonx.NewBsonDocument("$each", values)
	w.writeAdditionalFields(&each)

	return doc.Append(w.operator, each)
}

func (w WithEachUpdate[T]) Document() bsonx.Document {
	return w.ToBsonDocument().Document()
}

type PushUpdate[T expressions.TItem] struct {
	WithEachUpdate[T]
	options PushOptions
}

func NewPushUpdate[T expressions.TItem](fieldName string, values []T, options PushOptions) PushUpdate[T] {
	return PushUpdate[T]{
		WithEachUpdate: NewWithEachUpdate(fieldName, values, "$push"),
		options:        options,
	}
}

func (p PushUpdate[T]) writeAdditionalFields(document *bsonx.BsonDocument) {
	if p.options.HasPosition() {
		document.Append("$position", bsonx.NewBsonInt32(p.options.GetPosition()))
	}
	if p.options.HasSlice() {
		document.Append("$slice", bsonx.NewBsonInt32(p.options.GetSlice()))
	}
	if p.options.HasSort() {
		document.Append("$sort", bsonx.NewBsonInt32(p.options.GetSort()))
	} else {
		sortDocument := p.options.GetSortDocument()
		if sortDocument != nil {
			document.Append("$sort", sortDocument.ToBsonDocument())
		}
	}
}
func (p PushUpdate[T]) additionalFieldsToString() string {
	return ", options=" + p.options.ToString()
}

type PullAllUpdate[T expressions.TItem] struct {
	fieldName string
	values    []T
}

func NewPullAllUpdate[T expressions.TItem](fieldName string, values []T) PullAllUpdate[T] {
	return PullAllUpdate[T]{
		fieldName: fieldName,
		values:    values,
	}
}

func (p PullAllUpdate[T]) ToBsonDocument() bsonx.BsonDocument {
	values := bsonx.NewBsonArray()
	for _, value := range p.values {
		values.Append(value)
	}
	return bsonx.NewBsonDocument("$pullAll", bsonx.NewBsonDocument(p.fieldName, values))
}

func (p PullAllUpdate[T]) Document() bsonx.Document {
	return p.ToBsonDocument().Document()
}

type CompositeUpdate struct {
	updates []bsonx.Bson
}

func NewCompositeUpdate(updates []bsonx.Bson) CompositeUpdate {
	return CompositeUpdate{
		updates: updates,
	}
}

func (p CompositeUpdate) ToBsonDocument() bsonx.BsonDocument {
	document := bsonx.NewEmptyDoc()
	for _, update := range p.updates {
		rendered := update.ToBsonDocument()
		for key, value := range rendered.Data() {
			if document.ContainsKey(key) {
				if value.IsDocument() {
					currentOperatorDocument := value.AsDocument()
					existingOperatorDocument := document.GetDocument(key)
					for ck, cv := range currentOperatorDocument.Data() {
						existingOperatorDocument.Append(ck, cv)
					}
				}
			} else {
				document.Append(key, value)
			}
		}
	}
	return document
}

func (p CompositeUpdate) Document() bsonx.Document {
	return p.ToBsonDocument().Document()
}
