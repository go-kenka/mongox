package search

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

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

func (s SearchCount) ToBsonDocument() *bsonx.BsonDocument {
	return bsonx.BsonEmpty()
}

func (s SearchCount) Document() bson.D {
	return s.ToBsonDocument().Document()
}
