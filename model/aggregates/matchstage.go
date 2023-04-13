package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/filters"
	"go.mongodb.org/mongo-driver/bson"
)

type MatchStage struct {
	stage bsonx.Bson
}

func (s MatchStage) Bson() bsonx.Bson {
	return s.stage
}

func (s MatchStage) Document() bson.D {
	return s.stage.Document()
}
func (s MatchStage) Watch() {
}

// Match Filters the documents to pass only the documents that match the
// specified condition(s) to the next pipeline DefaultStage. The $match DefaultStage has the
// following prototype form: { $match: { <query> } } $match takes a document
// that specifies the query conditions. The query syntax is identical to the
// read operation query syntax; i.e. $match does not accept raw aggregation
// expressions. Instead, use a $expr query expression to include aggregation
// expression in $match.
func Match[T filters.MatchFilter](filter T) MatchStage {
	return MatchStage{stage: bsonx.BsonDoc("$match", filter.Value())}
}
