package bsonx

type IBsonValue interface {
	GetBsonType() BsonType
	Get() any
	AsDocument() BsonDocument
	AsArray() BsonArray
	AsString() BsonString
	AsNumber() IBsonNumber
	AsInt32() BsonInt32
	AsInt64() BsonInt64
	AsDecimal128() BsonDecimal128
	AsDouble() BsonDouble
	AsBoolean() BsonBoolean
	AsObjectId() BsonObjectId
	AsDBPointer() BsonDBPointer
	AsTimestamp() BsonTimestamp
	AsBinary() BsonBinary
	AsDateTime() BsonDateTime
	AsSymbol() BsonSymbol
	AsRegularExpression() BsonRegularExpression
	AsJavaScript() BsonJavaScript
	AsJavaScriptWithScope() BsonJavaScriptWithScope
	IsNull() bool
	IsDocument() bool
	IsArray() bool
	IsString() bool
	IsNumber() bool
	IsInt32() bool
	IsInt64() bool
	IsDecimal128() bool
	IsDouble() bool
	IsBoolean() bool
	IsObjectId() bool
	IsDBPointer() bool
	IsTimestamp() bool
	IsBinary() bool
	IsDateTime() bool
	IsSymbol() bool
	IsRegularExpression() bool
	IsJavaScript() bool
	IsJavaScriptWithScope() bool
}
type BsonValue struct {
}

func (v BsonValue) GetBsonType() BsonType {
	return END_OF_DOCUMENT
}

func (v BsonValue) Get() any {
	return nil
}

func (v BsonValue) AsDocument() BsonDocument {
	return BsonDocument{}
}
func (v BsonValue) AsArray() BsonArray {
	return BsonArray{}
}
func (v BsonValue) AsString() BsonString {
	return BsonString{}
}
func (v BsonValue) AsNumber() IBsonNumber {
	return BsonNumber{}
}
func (v BsonValue) AsInt32() BsonInt32 {
	return BsonInt32{}
}
func (v BsonValue) AsInt64() BsonInt64 {
	return BsonInt64{}
}
func (v BsonValue) AsDecimal128() BsonDecimal128 {
	return BsonDecimal128{}
}
func (v BsonValue) AsDouble() BsonDouble {
	return BsonDouble{}
}
func (v BsonValue) AsBoolean() BsonBoolean {
	return BsonBoolean{}
}
func (v BsonValue) AsObjectId() BsonObjectId {
	return BsonObjectId{}
}
func (v BsonValue) AsDBPointer() BsonDBPointer {
	return BsonDBPointer{}
}
func (v BsonValue) AsTimestamp() BsonTimestamp {
	return BsonTimestamp{}
}
func (v BsonValue) AsBinary() BsonBinary {
	return BsonBinary{}
}
func (v BsonValue) AsDateTime() BsonDateTime {
	return BsonDateTime{}
}
func (v BsonValue) AsSymbol() BsonSymbol {
	return BsonSymbol{}
}

func (v BsonValue) AsRegularExpression() BsonRegularExpression {
	return BsonRegularExpression{}
}

func (v BsonValue) AsJavaScript() BsonJavaScript {
	return BsonJavaScript{}
}

func (v BsonValue) AsJavaScriptWithScope() BsonJavaScriptWithScope {
	return BsonJavaScriptWithScope{}
}

func (v BsonValue) IsNull() bool {
	return false
}
func (v BsonValue) IsDocument() bool {
	return false
}
func (v BsonValue) IsArray() bool {
	return false
}
func (v BsonValue) IsString() bool {
	return false
}
func (v BsonValue) IsNumber() bool {
	return false
}
func (v BsonValue) IsInt32() bool {
	return false
}
func (v BsonValue) IsInt64() bool {
	return false
}
func (v BsonValue) IsDecimal128() bool {
	return false
}
func (v BsonValue) IsDouble() bool {
	return false
}
func (v BsonValue) IsBoolean() bool {
	return false
}
func (v BsonValue) IsObjectId() bool {
	return false
}
func (v BsonValue) IsDBPointer() bool {
	return false
}
func (v BsonValue) IsTimestamp() bool {
	return false
}
func (v BsonValue) IsBinary() bool {
	return false
}
func (v BsonValue) IsDateTime() bool {
	return false
}
func (v BsonValue) IsSymbol() bool {
	return false
}
func (v BsonValue) IsRegularExpression() bool {
	return false
}
func (v BsonValue) IsJavaScript() bool {
	return false
}
func (v BsonValue) IsJavaScriptWithScope() bool {
	return false
}
