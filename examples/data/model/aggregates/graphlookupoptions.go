package aggregates

import "github.com/go-kenka/mongox/examples/data/bsonx"

type GraphLookupOptions struct {
	maxDepth                int32
	depthField              string
	restrictSearchWithMatch bsonx.Bson
}

func (g GraphLookupOptions) DepthField(field string) GraphLookupOptions {
	g.depthField = field
	return g
}

func (g GraphLookupOptions) GetDepthField() string {
	return g.depthField
}

func (g GraphLookupOptions) MaxDepth(max int32) GraphLookupOptions {
	g.maxDepth = max
	return g
}

func (g GraphLookupOptions) GetMaxDepth() int32 {
	return g.maxDepth
}

func (g GraphLookupOptions) RestrictSearchWithMatch(r bsonx.Bson) GraphLookupOptions {
	g.restrictSearchWithMatch = r
	return g
}

func (g GraphLookupOptions) GetRestrictSearchWithMatch() bsonx.Bson {
	return g.restrictSearchWithMatch
}
