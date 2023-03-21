package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonTimestamp struct {
	BsonValue
	data primitive.Timestamp
}

func NewBsonTimestamp() BsonTimestamp {
	return BsonTimestamp{
		data: primitive.Timestamp{},
	}
}

func (a BsonTimestamp) GetBsonType() BsonType {
	return TIMESTAMP
}

func (a BsonTimestamp) Get() any {
	return a.data
}

func (a BsonTimestamp) AsTimestamp() BsonTimestamp {
	return a
}
func (a BsonTimestamp) IsTimestamp() bool {
	return true
}
