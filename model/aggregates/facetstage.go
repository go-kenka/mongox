package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type FacetStage struct {
	stage bsonx.Bson
}

func (s FacetStage) Bson() bsonx.Bson {
	return s.stage
}

func (s FacetStage) Document() bson.D {
	return s.stage.Document()
}

// Facets Processes multiple aggregation pipelines within a single DefaultStage on the same
// set of input documents. Each sub-pipeline has its own field in the output
// document where its results are stored as an array of documents. The $facet
// DefaultStage allows you to create multi-faceted aggregations which characterize data
// across multiple dimensions, or facets, within a single aggregation DefaultStage.
// Multi-faceted aggregations provide multiple filters and categorizations to
// guide data browsing and analysis. Retailers commonly use faceting to narrow
// search results by creating filters on product price, manufacturer, size, etc.
// Input documents are passed to the $facet DefaultStage only once. $facet enables
// various aggregations on the same set of input documents, without needing to
// retrieve the input documents multiple times. The $facet DefaultStage has the
// following form: { $facet:
//
//	  {
//	     <outputField1>: [ <stage1>, <stage2>, ... ],
//	     <outputField2>: [ <stage1>, <stage2>, ... ],
//	     ...
//	  }
//	}
//
// Specify the output field name for each specified pipeline.
func Facets(facets ...Facet) FacetStage {
	return FacetStage{stage: NewFacetStage(facets)}
}

type facetStage struct {
	facets []Facet
}

func NewFacetStage(facets []Facet) facetStage {
	return facetStage{
		facets: facets,
	}
}

func (bs facetStage) BsonDocument() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonEmpty()

	if len(bs.facets) > 0 {
		for _, f := range bs.facets {
			pipeline := bsonx.Array()
			for _, p := range f.pipeline {
				pipeline.Append(p.Bson().BsonDocument())
			}

			data.Append(f.name, pipeline)
		}
	}
	b.Append("$facet", data)
	return b
}

func (bs facetStage) Document() bson.D {
	return bs.BsonDocument().Document()
}
