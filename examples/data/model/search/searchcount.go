package search

import "github.com/go-kenka/mongox/examples/data/bsonx"

type ISearchCount interface {
	bsonx.Bson
	Total() TotalSearchCount
	LowerBound() LowerBoundSearchCount
	Of(count bsonx.Bson) SearchCount
}

type SearchCount struct {
}

func (s SearchCount) Total() TotalSearchCount {
	return TotalSearchCount{}
}

func (s SearchCount) LowerBound() LowerBoundSearchCount {
	return nil
}

func (s SearchCount) Of(count bsonx.Bson) SearchCount {
	return s
}

func (s SearchCount) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.BsonDocument{}
}

func (s SearchCount) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}
