package bsonx

type BsonField struct {
	name  string
	value Bson
}

func NewBsonField(name string, value Bson) BsonField {
	return BsonField{
		name:  name,
		value: value,
	}
}

func (f BsonField) GetName() string {
	return f.name
}

func (f BsonField) GetValue() Bson {
	return f.value
}
