package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonDBPointer struct {
	BsonValue
	data primitive.DBPointer
}

func NewBsonDBPointer(data primitive.DBPointer) BsonDBPointer {
	return BsonDBPointer{
		data: data,
	}
}

func (a BsonDBPointer) GetBsonType() BsonType {
	return DB_POINTER
}

func (a BsonDBPointer) Get() any {
	return a.data
}

func (a BsonDBPointer) AsDBPointer() BsonDBPointer {
	return a
}
func (a BsonDBPointer) IsDBPointer() bool {
	return true
}
