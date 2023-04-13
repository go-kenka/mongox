package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonMinKey struct {
	BsonValue
	data primitive.MinKey
}

func (a *BsonMinKey) Exp() IBsonValue {
	return a
}
func (a *BsonMinKey) GetBsonType() BsonType {
	return MIN_KEY
}

func (a *BsonMinKey) Get() any {
	return a.data
}
