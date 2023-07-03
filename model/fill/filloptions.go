package fill

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
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
	PartitionBy(expression expression.AnyExpression) FillOptions
	PartitionByFields(fields ...string) FillOptions
	SortBy(sortBy bsonx.Bson) FillOptions
	Option(name string, value any) FillOptions
}
