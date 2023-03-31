package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonBinary struct {
	BsonValue
	data primitive.Binary
}

func Binary(data primitive.Binary) *BsonBinary {
	return &BsonBinary{
		data: data,
	}
}

func (a *BsonBinary) GetBsonType() BsonType {
	return BINARY
}

func (a *BsonBinary) Get() any {
	return a.data
}

func (a *BsonBinary) AsBinary() *BsonBinary {
	return a
}
func (a *BsonBinary) IsBinary() bool {
	return true
}
