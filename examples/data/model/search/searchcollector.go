package search

import "github.com/go-kenka/mongox/examples/data/bsonx"

type SearchCollector interface {
	bsonx.Bson
}
