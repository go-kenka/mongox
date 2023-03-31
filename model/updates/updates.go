package updates

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

func Combine(updates ...bsonx.Bson) bsonx.Bson {
	return NewCompositeUpdate(updates)
}

func Set[I bsonx.Expression](fieldName string, value I) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$set")
}
func UnSet(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.String(""), "$unset")
}

func SetOnInsert[I bsonx.Expression](fieldName string, value I) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$setOnInsert")
}
func SetOnInsertBson(value bsonx.Bson) bsonx.Bson {
	return NewSimpleBsonKeyValue("$setOnInsert", value)
}
func Rename(fieldName, newFieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.String(newFieldName), "$rename")
}
func Inc(fieldName string, number int32) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.Int32(number), "$inc")
}
func Mul(fieldName string, number int32) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.Int32(number), "$mul")
}
func Min[I bsonx.Expression](fieldName string, value I) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$min")
}
func Max[I bsonx.Expression](fieldName string, value I) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$max")
}
func CurrentDate(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.Boolean(true), "$currentDate")
}
func CurrentTimestamp(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.BsonDoc("$type", bsonx.String("timestamp")), "$currentDate")
}
func AddToSet[I bsonx.Expression](fieldName string, value I) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$addToSet")
}
func AddEachToSet[I bsonx.Expression](fieldName string, values []I) bsonx.Bson {
	return NewWithEachUpdate(fieldName, values, "$addToSet")
}

func Push[I bsonx.Expression](fieldName string, value I) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$push")
}

func PushEach[I bsonx.Expression](fieldName string, values []I, options PushOptions) bsonx.Bson {
	return NewPushUpdate(fieldName, values, options)
}

func Pull[I bsonx.Expression](fieldName string, value I) bsonx.Bson {
	return NewSimpleUpdate(fieldName, value, "$pull")
}
func PullByFilter(filter bsonx.Bson) bsonx.Bson {
	return bsonx.BsonDoc("$pull", filter.ToBsonDocument())
}
func PullAll[I bsonx.Expression](fieldName string, values []I) bsonx.Bson {
	return NewPullAllUpdate(fieldName, values)
}
func PopFirst(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.Int32(-1), "$pop")
}

func PopLast(fieldName string) bsonx.Bson {
	return NewSimpleUpdate(fieldName, bsonx.Int32(1), "$pop")
}
func BitwiseAnd(fieldName string, value int64) bsonx.Bson {
	return createBitUpdateDocument(fieldName, "and", bsonx.Int64(value))
}

func BitwiseOr(fieldName string, value int64) bsonx.Bson {
	return createBitUpdateDocument(fieldName, "or", bsonx.Int64(value))
}

func BitwiseXor(fieldName string, value int64) bsonx.Bson {
	return createBitUpdateDocument(fieldName, "xor", bsonx.Int64(value))
}

func createBitUpdateDocument(fieldName, bitwiseOperator string, value bsonx.IBsonValue) bsonx.Bson {
	return bsonx.BsonDoc("$bit", bsonx.BsonDoc(fieldName, bsonx.BsonDoc(bitwiseOperator, value)))
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

func (s SimpleBsonKeyValue) ToBsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(s.fieldName, s.value.ToBsonDocument())
}
func (s SimpleBsonKeyValue) Document() bson.D {
	return s.ToBsonDocument().Document()
}

type SimpleUpdate[T bsonx.Expression] struct {
	fieldName string
	value     T
	operator  string
}

func NewSimpleUpdate[T bsonx.Expression](fieldName string, value T, operator string) SimpleUpdate[T] {
	return SimpleUpdate[T]{
		fieldName: fieldName,
		value:     value,
		operator:  operator,
	}
}

func (s SimpleUpdate[T]) ToBsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonDoc(s.operator, bsonx.BsonDoc(s.fieldName, s.value))
}

func (s SimpleUpdate[T]) Document() bson.D {
	return s.ToBsonDocument().Document()
}

type WithEachUpdate[T bsonx.Expression] struct {
	fieldName string
	values    []T
	operator  string
}

func NewWithEachUpdate[T bsonx.Expression](fieldName string, values []T, operator string) WithEachUpdate[T] {
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

func (w WithEachUpdate[T]) ToBsonDocument() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()
	values := bsonx.Array()
	for _, value := range w.values {
		values.Append(value)
	}
	each := bsonx.BsonDoc("$each", values)
	w.writeAdditionalFields(each)

	return doc.Append(w.operator, each)
}

func (w WithEachUpdate[T]) Document() bson.D {
	return w.ToBsonDocument().Document()
}

type PushUpdate[T bsonx.Expression] struct {
	WithEachUpdate[T]
	options PushOptions
}

func NewPushUpdate[T bsonx.Expression](fieldName string, values []T, options PushOptions) PushUpdate[T] {
	return PushUpdate[T]{
		WithEachUpdate: NewWithEachUpdate(fieldName, values, "$push"),
		options:        options,
	}
}

func (p PushUpdate[T]) writeAdditionalFields(document *bsonx.BsonDocument) {
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
			document.Append("$sort", sortDocument.ToBsonDocument())
		}
	}
}
func (p PushUpdate[T]) additionalFieldsToString() string {
	return ", options=" + p.options.ToString()
}

type PullAllUpdate[T bsonx.Expression] struct {
	fieldName string
	values    []T
}

func NewPullAllUpdate[T bsonx.Expression](fieldName string, values []T) PullAllUpdate[T] {
	return PullAllUpdate[T]{
		fieldName: fieldName,
		values:    values,
	}
}

func (p PullAllUpdate[T]) ToBsonDocument() *bsonx.BsonDocument {
	values := bsonx.Array()
	for _, value := range p.values {
		values.Append(value)
	}
	return bsonx.BsonDoc("$pullAll", bsonx.BsonDoc(p.fieldName, values))
}

func (p PullAllUpdate[T]) Document() bson.D {
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

func (p CompositeUpdate) ToBsonDocument() *bsonx.BsonDocument {
	document := bsonx.BsonEmpty()
	for _, update := range p.updates {
		rendered := update.ToBsonDocument()
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

func (p CompositeUpdate) Document() bson.D {
	return p.ToBsonDocument().Document()
}
