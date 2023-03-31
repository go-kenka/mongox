package search

import (
	"github.com/go-kenka/mongox/bsonx"
)

type SearchOperator interface {
	bsonx.Bson
}
