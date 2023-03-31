package bsonx

type BsonBoolean struct {
	BsonValue
	data bool
}

func (a *BsonBoolean) Exp() IBsonValue {
	return a
}

func Boolean(v bool) *BsonBoolean {
	return &BsonBoolean{
		data: v,
	}
}

func (a *BsonBoolean) GetBsonType() BsonType {
	return BOOLEAN
}

func (a *BsonBoolean) Get() any {
	return a.data
}

func (a *BsonBoolean) AsBoolean() *BsonBoolean {
	return a
}
func (a *BsonBoolean) IsBoolean() bool {
	return true
}
