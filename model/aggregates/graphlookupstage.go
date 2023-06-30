package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/bsonx/expression"
	"github.com/go-kenka/mongox/bsonx/options"
	"go.mongodb.org/mongo-driver/bson"
)

type GraphLookupStage struct {
	stage bsonx.Bson
}

func (s GraphLookupStage) Bson() bsonx.Bson {
	return s.stage
}

func (s GraphLookupStage) Document() bson.D {
	return s.stage.Document()
}

// GraphLookup Changed in version 5.1. Performs a recursive search on a
// collection, with options for restricting the search by recursion depth and
// query filter. The $graphLookup search process is summarized below: Input
// documents flow into the $graphLookup DefaultStage of an aggregation operation.
// $graphLookup targets the search to the collection designated by the from
// parameter (see below for full list of search parameters). For each input
// document, the search begins with the value designated by startWith.
// $graphLookup matches the startWith value against the field designated by
// connectToField in other documents in the from collection. For each matching
// document, $graphLookup takes the value of the connectFromField and checks
// every document in the from collection for a matching connectToField value.
// For each match, $graphLookup adds the matching document in the from
// collection to an array field named by the as parameter. This step continues
// recursively until no more matching documents are found, or until the
// operation reaches a recursion depth specified by the maxDepth parameter.
// $graphLookup then appends the array field to the input document. $graphLookup
// returns results after completing its search on all input documents.
// $graphLookup has the following prototype form:
//
//	{
//	  $graphLookup: {
//	     from: <collection>,
//	     startWith: <expression>,
//	     connectFromField: <string>,
//	     connectToField: <string>,
//	     as: <string>,
//	     maxDepth: <number>,
//	     depthField: <string>,
//	     restrictSearchWithMatch: <document>
//	  }
//	}
func GraphLookup[T expression.AnyExpression](
	from string,
	startWith T,
	connectFromField string,
	connectToField string,
	as string,
	options options.GraphLookupOptions,
) GraphLookupStage {
	return GraphLookupStage{stage: NewGraphLookupStage(from, startWith, connectFromField, connectToField, as, options)}
}

type graphLookupStage[T expression.AnyExpression] struct {
	from             string
	startWith        T
	connectFromField string
	connectToField   string
	as               string
	options          options.GraphLookupOptions
}

func NewGraphLookupStage[T expression.AnyExpression](
	from string,
	startWith T,
	connectFromField string,
	connectToField string,
	as string,
	options options.GraphLookupOptions,
) graphLookupStage[T] {
	return graphLookupStage[T]{
		from:             from,
		startWith:        startWith,
		connectFromField: connectFromField,
		connectToField:   connectToField,
		as:               as,
		options:          options,
	}
}

func (bs graphLookupStage[T]) BsonDocument() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonDoc("form", bsonx.String(bs.from))

	data.Append("startWith", bs.startWith)
	data.Append("connectFromField", bsonx.String(bs.connectFromField))
	data.Append("connectToField", bsonx.String(bs.connectToField))
	data.Append("as", bsonx.String(bs.as))

	if bs.options.GetMaxDepth() > 0 {
		data.Append("maxDepth", bsonx.Int32(bs.options.GetMaxDepth()))
	}
	if bs.options.GetDepthField() != "" {
		data.Append("depthField", bsonx.String(bs.options.GetDepthField()))
	}
	if bs.options.GetRestrictSearchWithMatch() != nil {
		data.Append("restrictSearchWithMatch", bs.options.GetRestrictSearchWithMatch().BsonDocument())
	}
	b.Append("$graphLookup", data)
	return b
}

func (bs graphLookupStage[T]) Document() bson.D {
	return bs.BsonDocument().Document()
}
