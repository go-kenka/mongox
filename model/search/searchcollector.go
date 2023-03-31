package search

import (
	"github.com/go-kenka/mongox/bsonx"
)

type SearchCollector interface {
	bsonx.Bson
}
