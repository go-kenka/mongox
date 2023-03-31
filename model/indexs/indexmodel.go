package indexs

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IndexModel struct {
	keys    bsonx.Bson
	options options.IndexOptions
}

func NewIndexModel(keys bsonx.Bson, options options.IndexOptions) IndexModel {
	return IndexModel{
		keys:    keys,
		options: options,
	}
}

func (m IndexModel) GetKeys() bsonx.Bson {
	return m.keys
}

func (m IndexModel) GetOptions() options.IndexOptions {
	return m.options
}
