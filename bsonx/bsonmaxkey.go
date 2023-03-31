package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonMaxKey struct {
	BsonValue
	data primitive.MaxKey
}

func (a *BsonMaxKey) GetBsonType() BsonType {
	return MAX_KEY
}

func (a *BsonMaxKey) Get() any {
	return a.data
}
