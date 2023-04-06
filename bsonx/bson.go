package bsonx

import "go.mongodb.org/mongo-driver/bson"

type Bson interface {
	Pro() *BsonDocument
	Document() bson.D
}

func document(v any) any {
	if v, ok := v.(IBsonValue); ok {
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

		// mp := bson.M{}
		// for _, v := range v.AsDocument().data {
		// 	mp[v.Key] = document(v.Value)
		// 	// mp = append(mp, bson.E{Key: v.Key, Value: document(v.Value)})
		// }
		mp := bson.D{}
		for _, v := range v.AsDocument().data {
			mp = append(mp, bson.E{Key: v.Key, Value: document(v.Value)})
		}
		return mp
	}

	return v
}
