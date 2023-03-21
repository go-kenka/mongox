package bsonx

type BsonString struct {
	BsonValue
	data string
}

func NewBsonString(data string) BsonString {
	return BsonString{
		data: data,
	}
}

func (a BsonString) GetBsonType() BsonType {
	return STRING
}

func (a BsonString) Get() any {
	return a.data
}

func (a BsonString) isEmpty() bool {
	return len(a.data) == 0
}

func (a BsonString) AsString() BsonString {
	return a
}
func (a BsonString) IsString() bool {
	return true
}
