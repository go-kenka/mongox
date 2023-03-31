package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonSymbol struct {
	BsonValue
	data primitive.Symbol
}

func Symbol(data primitive.Symbol) *BsonSymbol {
	return &BsonSymbol{
		data: data,
	}
}

func (a *BsonSymbol) GetBsonType() BsonType {
	return SYMBOL
}

func (a *BsonSymbol) Get() any {
	return a.data
}

func (a *BsonSymbol) AsSymbol() *BsonSymbol {
	return a
}

func (a *BsonSymbol) IsSymbol() bool {
	return true
}
