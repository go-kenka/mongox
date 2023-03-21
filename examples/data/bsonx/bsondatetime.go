package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonDateTime struct {
	BsonValue
	data primitive.DateTime
}

func NewBsonDateTime(data primitive.DateTime) BsonDateTime {
	return BsonDateTime{
		data: data,
	}
}

func (a BsonDateTime) GetBsonType() BsonType {
	return DATE_TIME
}

func (a BsonDateTime) Get() any {
	return a.data
}

func (a BsonDateTime) AsDateTime() BsonDateTime {
	return a
}
func (a BsonDateTime) IsDateTime() bool {
	return true
}
