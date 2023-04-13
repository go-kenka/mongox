package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonNull struct {
	BsonValue
	data primitive.Null
}

func Null() *BsonNull {
	return &BsonNull{
		data: primitive.Null{},
	}
}

func (a *BsonNull) Exp() IBsonValue {
	return a
}
func (a *BsonNull) GetBsonType() BsonType {
	return NULL
}

func (a *BsonNull) Get() any {
	return a.data
}

func (a *BsonNull) IsNull() bool {
	return true
}
