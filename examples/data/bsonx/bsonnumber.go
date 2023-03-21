package bsonx

import (
	"math/big"
)

type IBsonNumber interface {
	IBsonValue
	Int32Value() int32
	Int64Value() int64
	Float64Value() float64
	BigIntValue() *big.Int
}

type BsonNumber struct {
	BsonValue
}

func (a BsonNumber) Int32Value() int32 {
	return 0
}

func (a BsonNumber) Int64Value() int64 {
	return 0
}

func (a BsonNumber) Float64Value() float64 {
	return 0
}

func (a BsonNumber) BigIntValue() *big.Int {
	return big.NewInt(0)
}
