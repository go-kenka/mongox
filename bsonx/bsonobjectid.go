package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonObjectId struct {
	BsonValue
	data primitive.ObjectID
}

func ObjectId(data primitive.ObjectID) *BsonObjectId {
	return &BsonObjectId{
		data: data,
	}
}

func (a *BsonObjectId) GetBsonType() BsonType {
	return OBJECT_ID
}

func (a *BsonObjectId) Get() any {
	return a.data
}

func (a *BsonObjectId) AsObjectId() *BsonObjectId {
	return a
}
func (a *BsonObjectId) IsObjectId() bool {
	return true
}
