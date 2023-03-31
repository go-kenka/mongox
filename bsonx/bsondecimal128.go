package bsonx

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BsonDecimal128 struct {
	BsonValue
	data primitive.Decimal128
}

func Decimal128(data primitive.Decimal128) *BsonDecimal128 {
	return &BsonDecimal128{
		data: data,
	}
}

func (a *BsonDecimal128) GetBsonType() BsonType {
	return DECIMAL128
}

func (a *BsonDecimal128) Get() any {
	return a.data
}

func (a *BsonDecimal128) AsDecimal128() *BsonDecimal128 {
	return a
}
func (a *BsonDecimal128) IsDecimal128() bool {
	return true
}
