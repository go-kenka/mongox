package bsonx

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
	"strconv"
)

type BsonDecimal128 struct {
	BsonValue
	data primitive.Decimal128
}

func NewBsonDecimal128(data primitive.Decimal128) BsonDecimal128 {
	return BsonDecimal128{
		data: data,
	}
}

func (a BsonDecimal128) GetBsonType() BsonType {
	return DECIMAL128
}

func (a BsonDecimal128) Get() any {
	return a.data
}

func (a BsonDecimal128) Int32Value() int32 {
	ds := a.data.String()
	data, _ := strconv.ParseInt(ds, 10, 64)
	return int32(data)
}

func (a BsonDecimal128) Int64Value() int64 {
	ds := a.data.String()
	data, _ := strconv.ParseInt(ds, 10, 64)
	return data
}

func (a BsonDecimal128) Float64Value() float64 {
	ds := a.data.String()
	data, _ := strconv.ParseFloat(ds, 64)
	return data
}

func (a BsonDecimal128) BigIntValue() *big.Int {
	data, _, _ := a.data.BigInt()
	return data
}

func (a BsonDecimal128) AsDecimal128() BsonDecimal128 {
	return a
}
func (a BsonDecimal128) IsDecimal128() bool {
	return true
}
