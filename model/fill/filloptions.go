package fill

import (
	"github.com/go-kenka/mongox/bsonx"
)

var (
	EmptyDoc                       = bsonx.Empty()
	DefaultFillOptions FillOptions = FillConstructibleBson{}.of(EmptyDoc)
	PartitionByFields              = func(o FillConstructibleBson, fields ...string) FillOptions {
		return o.PartitionByFields(fields...)
	}
)

type FillOptions interface {
	bsonx.Bson
	PartitionBy(expression bsonx.Expression) FillOptions
	PartitionByFields(fields ...string) FillOptions
	SortBy(sortBy bsonx.Bson) FillOptions
	Option(name string, value any) FillOptions
}
