package aggregates

import "github.com/go-kenka/mongox/examples/data/bsonx"

type Facet struct {
	name     string
	pipeline []bsonx.Bson
}

func NewFacet(name string, pipeline []bsonx.Bson) Facet {
	return Facet{
		name:     name,
		pipeline: pipeline,
	}
}
