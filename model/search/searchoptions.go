package search

import (
	"github.com/go-kenka/mongox/bsonx"
)

type SearchOptions interface {
	bsonx.Bson

	Index(name string) SearchOptions
	Highlight(option SearchHighlight) SearchOptions
	Count(option SearchCount) SearchOptions
	ReturnStoredSource(returnStoredSource bool) SearchOptions
	Option(name string, value any) SearchOptions
}

func DefaultSearchOptions() SearchOptions {
	return nil
}
