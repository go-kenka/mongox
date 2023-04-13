package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonUndefined struct {
	BsonValue
	data primitive.Undefined
}

func Undefined() *BsonUndefined {
	return &BsonUndefined{
		data: primitive.Undefined{},
	}
}

func (a *BsonUndefined) Exp() IBsonValue {
	return a
}

func (a *BsonUndefined) GetBsonType() BsonType {
	return UNDEFINED
}

func (a *BsonUndefined) Get() any {
	return a.data
}
