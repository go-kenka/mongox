package bsonx

type Bson interface {
	ToBsonDocument() BsonDocument
	Document() Document
}

func document(v IBsonValue) any {
	if v.IsArray() {
		var values []any
		for _, v := range v.AsArray().data {
			values = append(values, document(v))
		}
		return values
	}
	if !v.IsDocument() {
		return v.Get()
	}

	mp := make(map[string]any)
	for k, v := range v.AsDocument().data {
		mp[k] = document(v)
	}
	return mp
}
