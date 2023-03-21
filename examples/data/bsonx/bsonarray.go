package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonArray struct {
	BsonValue
	data []IBsonValue
}

func NewBsonArray(values ...IBsonValue) BsonArray {
	return BsonArray{
		data: values,
	}
}

func (a BsonArray) Append(value IBsonValue) BsonArray {
	a.data = append(a.data, value)
	return a
}

func (a BsonArray) GetBsonType() BsonType {
	return ARRAY
}

func (a BsonArray) Get() any {
	return a.data
}

func (a BsonArray) List() primitive.A {
	var list []any
	for _, v := range a.data {
		list = append(list, v.Get())
	}
	return list
}

func (a BsonArray) AsArray() BsonArray {
	return a
}

func (a BsonArray) IsArray() bool {
	return true
}
