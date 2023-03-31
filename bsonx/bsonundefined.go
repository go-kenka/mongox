package bsonx

type BsonUndefined struct {
	BsonValue
}

func (a *BsonUndefined) GetBsonType() BsonType {
	return UNDEFINED
}

func (a *BsonUndefined) Get() any {
	return nil
}
