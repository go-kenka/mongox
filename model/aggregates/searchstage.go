package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/search"
	"go.mongodb.org/mongo-driver/bson"
)

type SearchStage Stage

// Search The $search stage performs a full-text search on the specified field or
// fields which must be covered by an Atlas Search index. $search A $search
// pipeline stage has the following prototype form:
//
//	{
//	 $search: {
//	   "index": "<index-name>",
//	   "<operator-name>"|"<collector-name>": {
//	     <operator-specification>|<collector-specification>
//	   },
//	   "highlight": {
//	     <highlight-options>
//	   },
//	   "count": {
//	     <count-options>
//	   },
//	   "returnStoredSource": true | false
//	 }
//	}
func Search(operator search.SearchOperator, options search.SearchOptions) SearchStage {
	return NewSearchStage("$search", operator, options)
}

// SearchWithCollector The $search stage performs a full-text search on the specified field or
// fields which must be covered by an Atlas Search index. $search A $search
// pipeline stage has the following prototype form:
//
//	{
//	 $search: {
//	   "index": "<index-name>",
//	   "<operator-name>"|"<collector-name>": {
//	     <operator-specification>|<collector-specification>
//	   },
//	   "highlight": {
//	     <highlight-options>
//	   },
//	   "count": {
//	     <count-options>
//	   },
//	   "returnStoredSource": true | false
//	 }
//	}
func SearchWithCollector(collector search.SearchCollector, options search.SearchOptions) SearchStage {
	return NewSearchStage("$search", collector, options)
}

// SearchMeta The $searchMeta stage returns different types of metadata result documents.
// $searchMeta A $searchMeta pipeline stage has the following prototype form:
//
//	{
//	 $searchMeta: {
//	   "index": "<index-name>",
//	   "<collector-name>"|"<operator-name>": {
//	     <collector-specification>|<operator-specification>
//	   },
//	   "count": {
//	     <count-options>
//	   }
//	 }
//	}
func SearchMeta(operator search.SearchOperator, options search.SearchOptions) SearchStage {
	return NewSearchStage("$searchMeta", operator, options)
}

// SearchMetaWithCollector The $searchMeta stage returns different types of metadata result documents.
// $searchMeta A $searchMeta pipeline stage has the following prototype form:
//
//	{
//	 $searchMeta: {
//	   "index": "<index-name>",
//	   "<collector-name>"|"<operator-name>": {
//	     <collector-specification>|<operator-specification>
//	   },
//	   "count": {
//	     <count-options>
//	   }
//	 }
//	}
func SearchMetaWithCollector(collector search.SearchCollector, options search.SearchOptions) SearchStage {
	return NewSearchStage("$searchMeta", collector, options)
}

// searchStage TODO: SearchStage未完成
type searchStage struct {
	name                string
	operatorOrCollector bsonx.Bson
	options             search.SearchOptions
}

func (f searchStage) Bson() bsonx.Bson {
	return f.Pro()
}

func NewSearchStage(
	name string,
	operatorOrCollector bsonx.Bson,
	options search.SearchOptions,
) searchStage {
	return searchStage{
		name:                name,
		operatorOrCollector: operatorOrCollector,
		options:             options,
	}
}

func (f searchStage) Pro() *bsonx.BsonDocument {
	return bsonx.BsonEmpty()
}

func (f searchStage) Document() bson.D {
	return f.Pro().Document()
}
