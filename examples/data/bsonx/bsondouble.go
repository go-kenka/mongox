package bsonx

import "math/big"

type BsonDouble struct {
	BsonNumber
	data float64
}

func NewBsonDouble(data float64) BsonDouble {
	return BsonDouble{
		data: data,
	}
}

func (a BsonDouble) GetBsonType() BsonType {
	return DOUBLE
}

func (a BsonDouble) Get() any {
	return a.data
}

func (a BsonDouble) Int32Value() int32 {
	return int32(a.data)
}

func (a BsonDouble) Int64Value() int64 {
	return int64(a.data)
}

func (a BsonDouble) Float64Value() float64 {
	return a.data
}

func (a BsonDouble) BigIntValue() *big.Int {
	return big.NewInt(int64(a.data))
}

func (a BsonDouble) AsNumber() IBsonNumber {
	return a
}
func (a BsonDouble) AsDouble() BsonDouble {
	return a
}
func (a BsonDouble) IsNumber() bool {
	return true
}
func (a BsonDouble) IsDouble() bool {
	return true
}
