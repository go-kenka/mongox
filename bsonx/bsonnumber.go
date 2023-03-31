package bsonx

type IBsonNumber interface {
	IBsonValue
	Int32Value() int32
	Int64Value() int64
	Float64Value() float64
}

type BsonNumber struct {
	BsonValue
}

func (a *BsonNumber) Int32Value() int32 {
	return 0
}

func (a *BsonNumber) Int64Value() int64 {
	return 0
}

func (a *BsonNumber) Float64Value() float64 {
	return 0
}
