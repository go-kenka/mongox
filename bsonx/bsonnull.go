package bsonx

type BsonNull struct {
	BsonValue
}

func (a *BsonNull) GetBsonType() BsonType {
	return NULL
}

func (a *BsonNull) Get() any {
	return nil
}

func (a *BsonNull) IsNull() bool {
	return true
}
