package bsonx

type BsonDocument struct {
	BsonValue
	data map[string]IBsonValue
}

func NewEmptyDoc() BsonDocument {
	return BsonDocument{
		data: map[string]IBsonValue{},
	}
}

func NewBsonDocument(key string, value IBsonValue) BsonDocument {
	return BsonDocument{
		data: map[string]IBsonValue{
			key: value,
		},
	}
}

func (d BsonDocument) ToBsonDocument() BsonDocument {
	return d
}
func (d BsonDocument) Document() Document {
	data := make(map[string]any)
	for k, v := range d.data {
		data[k] = document(v)
	}
	return data
}

func (d BsonDocument) Append(key string, value IBsonValue) BsonDocument {
	d.data[key] = value
	return d
}
func (d BsonDocument) Remove(key string) BsonDocument {
	delete(d.data, key)
	return d
}
func (d BsonDocument) GetBsonType() BsonType {
	return DOCUMENT
}
func (d BsonDocument) Get() any {
	return d.data
}
func (d BsonDocument) GetValue(key string) IBsonValue {
	return d.data[key]
}
func (d BsonDocument) Keys() []string {
	var keys []string
	for k := range d.data {
		keys = append(keys, k)
	}
	return keys
}
func (d BsonDocument) Data() map[string]IBsonValue {
	return d.data
}
func (d BsonDocument) AsDocument() BsonDocument {
	return d
}
func (d BsonDocument) IsDocument() bool {
	return true
}
func (d BsonDocument) Size() int {
	return len(d.data)
}
func (d BsonDocument) IsEmpty() bool {
	return d.Size() == 0
}
func (d BsonDocument) ContainsKey(key string) bool {
	_, ok := d.data[key]
	return ok
}
func (d BsonDocument) ContainsValue(val any) bool {
	for _, a := range d.data {
		if a == val {
			return true
		}
	}
	return false
}

func (d BsonDocument) GetDocument(key string) BsonDocument {
	return d.data[key].(BsonDocument)
}
func (d BsonDocument) GetArray(key string) BsonArray {
	return d.data[key].(BsonArray)
}

func (d BsonDocument) GetNumber(key string) IBsonNumber {
	return d.data[key].(BsonNumber)
}

func (d BsonDocument) GetInt32(key string) BsonInt32 {
	return d.data[key].(BsonInt32)
}

func (d BsonDocument) GetInt64(key string) BsonInt64 {
	return d.data[key].(BsonInt64)
}

func (d BsonDocument) GetDecimal128(key string) BsonDecimal128 {
	return d.data[key].(BsonDecimal128)
}

func (d BsonDocument) GetDouble(key string) BsonDouble {
	return d.data[key].(BsonDouble)
}

func (d BsonDocument) GetBoolean(key string) BsonBoolean {
	return d.data[key].(BsonBoolean)
}

func (d BsonDocument) GetString(key string) BsonString {
	return d.data[key].(BsonString)
}

func (d BsonDocument) GetDateTime(key string) BsonDateTime {
	return d.data[key].(BsonDateTime)
}

func (d BsonDocument) GetTimestamp(key string) BsonTimestamp {
	return d.data[key].(BsonTimestamp)
}

func (d BsonDocument) GetObjectId(key string) BsonObjectId {
	return d.data[key].(BsonObjectId)
}

func (d BsonDocument) GetRegularExpression(key string) BsonRegularExpression {
	return d.data[key].(BsonRegularExpression)
}

func (d BsonDocument) GetBinary(key string) BsonBinary {
	return d.data[key].(BsonBinary)
}

func (d BsonDocument) IsBsonNull(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsNull()
}
func (d BsonDocument) IsBsonDocument(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsDocument()
}
func (d BsonDocument) IsBsonArray(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsArray()
}
func (d BsonDocument) IsBsonNumber(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsNumber()
}
func (d BsonDocument) IsBsonInt32(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsInt32()
}
func (d BsonDocument) IsBsonInt64(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsInt64()
}
func (d BsonDocument) IsBsonDecimal128(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsDecimal128()
}
func (d BsonDocument) IsBsonDouble(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsDouble()
}
func (d BsonDocument) IsBsonBoolean(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsBoolean()
}
func (d BsonDocument) IsBsonString(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsString()
}
func (d BsonDocument) IsBsonDateTime(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsDateTime()
}
func (d BsonDocument) IsBsonTimestamp(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsTimestamp()
}
func (d BsonDocument) IsBsonObjectId(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsObjectId()
}
func (d BsonDocument) IsBsonBinary(key string) bool {
	if !d.ContainsKey(key) {
		return false
	}
	return d.data[key].(BsonValue).IsBinary()
}

func (d BsonDocument) GetBsonDocument(key string, defaultValue BsonDocument) BsonDocument {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsDocument()
}

func (d BsonDocument) GetBsonArray(key string, defaultValue BsonArray) BsonArray {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsArray()
}

func (d BsonDocument) GetBsonNumber(key string, defaultValue BsonNumber) IBsonNumber {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsNumber()
}

func (d BsonDocument) GetBsonInt32(key string, defaultValue BsonInt32) BsonInt32 {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsInt32()
}

func (d BsonDocument) GetBsonInt64(key string, defaultValue BsonInt64) BsonInt64 {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsInt64()
}

func (d BsonDocument) GetBsonDecimal128(key string, defaultValue BsonDecimal128) BsonDecimal128 {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsDecimal128()
}

func (d BsonDocument) GetBsonDouble(key string, defaultValue BsonDouble) BsonDouble {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsDouble()
}

func (d BsonDocument) GetBsonBoolean(key string, defaultValue BsonBoolean) BsonBoolean {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsBoolean()
}

func (d BsonDocument) GetBsonString(key string, defaultValue BsonString) BsonString {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsString()
}

func (d BsonDocument) GetBsonDateTime(key string, defaultValue BsonDateTime) BsonDateTime {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsDateTime()
}

func (d BsonDocument) GetBsonTimestamp(key string, defaultValue BsonTimestamp) BsonTimestamp {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsTimestamp()
}

func (d BsonDocument) GetBsonObjectId(key string, defaultValue BsonObjectId) BsonObjectId {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsObjectId()
}

func (d BsonDocument) GetBsonBinary(key string, defaultValue BsonBinary) BsonBinary {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsBinary()
}

func (d BsonDocument) GetBsonRegularExpression(key string, defaultValue BsonRegularExpression) BsonRegularExpression {
	if !d.ContainsKey(key) {
		return defaultValue
	}
	return d.data[key].(BsonValue).AsRegularExpression()
}

func NewMerged(base, appended BsonDocument) BsonDocument {
	result := NewEmptyDoc()
	for k, v := range base.data {
		result.Append(k, v.(BsonValue))
	}
	for k, v := range appended.data {
		result.Append(k, v.(BsonValue))
	}
	return result

}
