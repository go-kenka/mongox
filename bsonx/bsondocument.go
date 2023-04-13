package bsonx

import (
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
)

type E struct {
	Key   string
	Value IBsonValue
}

type D []E

type BsonDocument struct {
	BsonValue
	data D
}

func BsonEmpty() *BsonDocument {
	return &BsonDocument{
		data: D{},
	}
}

func BsonDoc(key string, value IBsonValue) *BsonDocument {
	return &BsonDocument{
		data: D{{key, value}},
	}
}

func (d *BsonDocument) BsonDocument() *BsonDocument {
	return d
}

func (d *BsonDocument) Document() bson.D {
	data := Document{}
	for _, v := range d.data {
		data.Append(v.Key, document(v.Value))
	}
	return data.Document()
}

func (d *BsonDocument) Append(key string, value IBsonValue) *BsonDocument {
	d.data = append(d.data, E{Key: key, Value: value})
	return d
}
func (d *BsonDocument) Remove(key string) *BsonDocument {
	lo.DropWhile(d.data, func(item E) bool {
		return item.Key == key
	})
	return d
}
func (d *BsonDocument) GetBsonType() BsonType {
	return DOCUMENT
}
func (d *BsonDocument) Get() any {
	return d.data
}
func (d *BsonDocument) Exp() IBsonValue {
	return d
}
func (d *BsonDocument) Data() D {
	return d.data
}
func (d *BsonDocument) GetValue(key string) IBsonValue {
	v, ok := lo.Find(d.data, func(item E) bool {
		return item.Key == key
	})
	if !ok {
		return nil
	}
	return v.Value
}
func (d *BsonDocument) Keys() []string {
	var keys []string
	for _, v := range d.data {
		keys = append(keys, v.Key)
	}
	return keys
}

func (d *BsonDocument) AsDocument() *BsonDocument {
	return d
}
func (d *BsonDocument) IsDocument() bool {
	return true
}
func (d *BsonDocument) Size() int {
	return len(d.data)
}
func (d *BsonDocument) IsEmpty() bool {
	return d.Size() == 0
}
func (d *BsonDocument) ContainsKey(key string) bool {
	_, ok := lo.Find(d.data, func(item E) bool {
		return item.Key == key
	})
	return ok
}
func (d *BsonDocument) ContainsValue(val any) bool {
	for _, a := range d.data {
		if a == val {
			return true
		}
	}
	return false
}

func (d *BsonDocument) GetDocument(key string) *BsonDocument {
	return d.GetValue(key).(*BsonDocument)
}
func (d *BsonDocument) GetArray(key string) *BsonArray {
	return d.GetValue(key).(*BsonArray)
}

func (d *BsonDocument) GetNumber(key string) IBsonNumber {
	return d.GetValue(key).(*BsonNumber)
}

func (d *BsonDocument) GetInt32(key string) *BsonInt32 {
	return d.GetValue(key).(*BsonInt32)
}

func (d *BsonDocument) GetInt64(key string) *BsonInt64 {
	return d.GetValue(key).(*BsonInt64)
}

func (d *BsonDocument) GetDecimal128(key string) *BsonDecimal128 {
	return d.GetValue(key).(*BsonDecimal128)
}

func (d *BsonDocument) GetDouble(key string) *BsonDouble {
	return d.GetValue(key).(*BsonDouble)
}

func (d *BsonDocument) GetBoolean(key string) *BsonBoolean {
	return d.GetValue(key).(*BsonBoolean)
}

func (d *BsonDocument) GetString(key string) *BsonString {
	return d.GetValue(key).(*BsonString)
}

func (d *BsonDocument) GetDateTime(key string) *BsonDateTime {
	return d.GetValue(key).(*BsonDateTime)
}

func (d *BsonDocument) GetTimestamp(key string) *BsonTimestamp {
	return d.GetValue(key).(*BsonTimestamp)
}

func (d *BsonDocument) GetObjectId(key string) *BsonObjectId {
	return d.GetValue(key).(*BsonObjectId)
}

func (d *BsonDocument) GetRegularExpression(key string) *BsonRegularExpression {
	return d.GetValue(key).(*BsonRegularExpression)
}

func (d *BsonDocument) GetBinary(key string) *BsonBinary {
	return d.GetValue(key).(*BsonBinary)
}

func (d *BsonDocument) IsBsonNull(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsNull()
}
func (d *BsonDocument) IsBsonDocument(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsDocument()
}
func (d *BsonDocument) IsBsonArray(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsArray()
}
func (d *BsonDocument) IsBsonNumber(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsNumber()
}
func (d *BsonDocument) IsBsonInt32(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsInt32()
}
func (d *BsonDocument) IsBsonInt64(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsInt64()
}
func (d *BsonDocument) IsBsonDecimal128(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsDecimal128()
}
func (d *BsonDocument) IsBsonDouble(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsDouble()
}
func (d *BsonDocument) IsBsonBoolean(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsBoolean()
}
func (d *BsonDocument) IsBsonString(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsString()
}
func (d *BsonDocument) IsBsonDateTime(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsDateTime()
}
func (d *BsonDocument) IsBsonTimestamp(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsTimestamp()
}
func (d *BsonDocument) IsBsonObjectId(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsObjectId()
}
func (d *BsonDocument) IsBsonBinary(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.GetValue(key).(*BsonValue).IsBinary()
}

func (d *BsonDocument) GetBsonDocument(key string, defaultValue BsonDocument) *BsonDocument {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsDocument()
}

func (d *BsonDocument) GetBsonArray(key string, defaultValue BsonArray) *BsonArray {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsArray()
}

func (d *BsonDocument) GetBsonNumber(key string, defaultValue BsonNumber) IBsonNumber {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsNumber()
}

func (d *BsonDocument) GetBsonInt32(key string, defaultValue BsonInt32) *BsonInt32 {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsInt32()
}

func (d *BsonDocument) GetBsonInt64(key string, defaultValue BsonInt64) *BsonInt64 {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsInt64()
}

func (d *BsonDocument) GetBsonDecimal128(key string, defaultValue BsonDecimal128) *BsonDecimal128 {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsDecimal128()
}

func (d *BsonDocument) GetBsonDouble(key string, defaultValue BsonDouble) *BsonDouble {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsDouble()
}

func (d *BsonDocument) GetBsonBoolean(key string, defaultValue BsonBoolean) *BsonBoolean {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsBoolean()
}

func (d *BsonDocument) GetBsonString(key string, defaultValue BsonString) *BsonString {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsString()
}

func (d *BsonDocument) GetBsonDateTime(key string, defaultValue BsonDateTime) *BsonDateTime {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsDateTime()
}

func (d *BsonDocument) GetBsonTimestamp(key string, defaultValue BsonTimestamp) *BsonTimestamp {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsTimestamp()
}

func (d *BsonDocument) GetBsonObjectId(key string, defaultValue BsonObjectId) *BsonObjectId {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsObjectId()
}

func (d *BsonDocument) GetBsonBinary(key string, defaultValue BsonBinary) *BsonBinary {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsBinary()
}

func (d *BsonDocument) GetBsonRegularExpression(key string, defaultValue BsonRegularExpression) *BsonRegularExpression {
	if !d.ContainsKey(key) {
		return &defaultValue
	}
	return d.GetValue(key).(*BsonValue).AsRegularExpression()
}

func NewMerged(base, appended *BsonDocument) *BsonDocument {
	result := BsonEmpty()
	for _, v := range base.data {
		result.Append(v.Key, v.Value)
	}
	for _, v := range appended.data {
		_, index, ok := lo.FindIndexOf(result.data, func(item E) bool {
			return item.Key == v.Key
		})
		if ok {
			// 替换掉原有内容
			result.data[index].Value = v.Value
		} else {
			result.Append(v.Key, v.Value)
		}

	}
	return result

}
