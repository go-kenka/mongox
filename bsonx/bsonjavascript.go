package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonJavaScript struct {
	BsonValue
	data primitive.JavaScript
}

func JavaScript(data primitive.JavaScript) *BsonJavaScript {
	return &BsonJavaScript{
		data: data,
	}
}

func (a *BsonJavaScript) Exp() IBsonValue {
	return a
}

func (a *BsonJavaScript) GetBsonType() BsonType {
	return JAVASCRIPT
}

func (a *BsonJavaScript) Get() any {
	return a.data
}

func (a *BsonJavaScript) AsJavaScript() *BsonJavaScript {
	return a
}

func (a *BsonJavaScript) IsJavaScript() bool {
	return true
}
