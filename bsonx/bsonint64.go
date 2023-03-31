package bsonx

type BsonInt64 struct {
	BsonValue
	data int64
}

func Int64(data int64) *BsonInt64 {
	return &BsonInt64{
		data: data,
	}
}

func (a *BsonInt64) GetBsonType() BsonType {
	return INT64
}
func (a *BsonInt64) Get() any {
	return a.data
}
func (a *BsonInt64) Exp() IBsonValue {
	return a
}
func (a *BsonInt64) Int32Value() int32 {
	return int32(a.data)
}
func (a *BsonInt64) Int64Value() int64 {
	return a.data
}
func (a *BsonInt64) Float64Value() float64 {
	return float64(a.data)
}
func (a *BsonInt64) AsNumber() IBsonNumber {
	return a
}
func (a *BsonInt64) AsInt64() *BsonInt64 {
	return a
}
func (a *BsonInt64) IsNumber() bool {
	return true
}
func (a *BsonInt64) IsInt64() bool {
	return true
}
