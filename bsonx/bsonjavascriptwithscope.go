package bsonx

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonJavaScriptWithScope struct {
	BsonValue
	data primitive.CodeWithScope
}

func JavaScriptWithScope(data primitive.CodeWithScope) *BsonJavaScriptWithScope {
	return &BsonJavaScriptWithScope{
		data: data,
	}
}

func (a *BsonJavaScriptWithScope) Exp() IBsonValue {
	return a
}

func (a *BsonJavaScriptWithScope) GetBsonType() BsonType {
	return JAVASCRIPT_WITH_SCOPE
}

func (a *BsonJavaScriptWithScope) Get() any {
	return a.data
}

func (a *BsonJavaScriptWithScope) AsJavaScriptWithScope() *BsonJavaScriptWithScope {
	return a
}

func (a *BsonJavaScriptWithScope) IsJavaScriptWithScope() bool {
	return true
}
