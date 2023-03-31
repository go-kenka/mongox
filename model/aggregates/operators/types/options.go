package types

import "github.com/go-kenka/mongox/bsonx"

type ConvertOptions struct {
	onError bsonx.Bson
	onNull  bsonx.Bson
}
