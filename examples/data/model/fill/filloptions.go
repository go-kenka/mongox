package fill

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/expressions"
)

var (
	EmptyDoc                       = bsonx.NewDoc()
	DefaultFillOptions FillOptions = FillConstructibleBson{}.of(EmptyDoc)
	PartitionByFields              = func(o FillConstructibleBson, fields ...string) FillOptions {
		return o.PartitionByFields(fields...)
	}
)

type FillOptions interface {
	bsonx.Bson
	PartitionBy(expression expressions.TExpression) FillOptions
	PartitionByFields(fields ...string) FillOptions
	SortBy(sortBy bsonx.Bson) FillOptions
	Option(name string, value any) FillOptions
}
