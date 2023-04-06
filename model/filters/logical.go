package filters

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type logicalFilter struct {
	filter bsonx.Bson
}

func (f logicalFilter) Value() bsonx.IBsonValue {
	return f.filter.ToBsonDocument()
}

func (f logicalFilter) Document() bson.D {
	return f.filter.Document()
}

// And Joins query clauses with a logical AND returns all documents that match the
// conditions of both clauses.
func And[T Filter](filters ...T) logicalFilter {
	return logicalFilter{filter: newAddFilter(filters)}
}

// Or Inverts the effect of a query expression and returns documents that do not
// match the query expression.
func Or[T Filter](filters ...T) logicalFilter {
	return logicalFilter{filter: newOrNorFilter(OR, filters)}
}

// Nor Joins query clauses with a logical NOR returns all documents that fail to
// match both clauses.
func Nor[T Filter](filters ...T) logicalFilter {
	return logicalFilter{filter: newOrNorFilter(NOR, filters)}
}

// Not Joins query clauses with a logical OR returns all documents that match the
// conditions of either clause.
func Not[T Filter](f T) logicalFilter {
	return logicalFilter{filter: newNotFilter(f)}
}
