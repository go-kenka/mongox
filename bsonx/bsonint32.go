package bsonx

type BsonInt32 struct {
	BsonValue
	data int32
}

func Int32(data int32) *BsonInt32 {
	return &BsonInt32{
		data: data,
	}
}

func (a *BsonInt32) GetBsonType() BsonType {
	return INT32
}

func (a *BsonInt32) Get() any {
	return a.data
}

func (a *BsonInt32) Exp() IBsonValue {
	return a
}

func (a *BsonInt32) Int32Value() int32 {
	return a.data
}

func (a *BsonInt32) Int64Value() int64 {
	return int64(a.data)
}

func (a *BsonInt32) Float64Value() float64 {
	return float64(a.data)
}

func (a *BsonInt32) AsNumber() IBsonNumber {
	return a
}
func (a *BsonInt32) AsInt32() *BsonInt32 {
	return a
}
func (a *BsonInt32) IsNumber() bool {
	return true
}
func (a *BsonInt32) IsInt32() bool {
	return true
}
