package densify

import (
	bsonx2 "github.com/go-kenka/mongox/bsonx"
)

var (
	EmptyDoc                             = bsonx2.Empty()
	DefaultDensifyOptions DensifyOptions = DensifyConstructibleBson{}.of(EmptyDoc)
)

type DensifyOptions interface {
	bsonx2.Bson
	PartitionByFields(fields ...string) DensifyOptions
	Option(name string, value any) DensifyOptions
}
