package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"go.mongodb.org/mongo-driver/bson"
)

type SortByCountStage Stage

// SortByCount Groups incoming documents based on the value of a specified expression, then
// computes the count of documents in each distinct group. Each output document
// contains two fields: an _id field containing the distinct grouping value, and
// a count field containing the number of documents belonging to that grouping
// or category. The documents are sorted by count in descending order. The
// $sortByCount stage has the following prototype form: { $sortByCount:
// <expression> }
func SortByCount[T expression.AnyExpression](filter T) SortByCountStage {
	return NewSortByCountStage(filter)
}

type sortByCountStage[T expression.AnyExpression] struct {
	filter T
}

func (bs sortByCountStage[T]) Bson() bsonx.Bson {
	return bs.Pro()
}

func NewSortByCountStage[T expression.AnyExpression](filter T) sortByCountStage[T] {
	return sortByCountStage[T]{
		filter: filter,
	}
}

func (bs sortByCountStage[T]) Pro() *bsonx.BsonDocument {
	return bsonx.BsonDoc("$sortByCount", bs.filter)
}
func (bs sortByCountStage[T]) Document() bson.D {
	return bs.Pro().Document()
}
