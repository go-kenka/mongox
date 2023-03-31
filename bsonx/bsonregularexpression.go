package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonRegularExpression struct {
	BsonValue
	data primitive.Regex
}

func RegularExpression(data primitive.Regex) *BsonRegularExpression {
	return &BsonRegularExpression{
		data: data,
	}
}

func (a *BsonRegularExpression) GetBsonType() BsonType {
	return REGULAR_EXPRESSION
}

func (a *BsonRegularExpression) Get() any {
	return a.data
}

func (a *BsonRegularExpression) AsRegularExpression() *BsonRegularExpression {
	return a
}

func (a *BsonRegularExpression) IsRegularExpression() bool {
	return true
}
