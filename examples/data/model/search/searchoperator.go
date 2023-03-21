package search

import "github.com/go-kenka/mongox/examples/data/bsonx"

type SearchOperator interface {
	bsonx.Bson
}
