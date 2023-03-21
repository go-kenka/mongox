package densify

import "github.com/go-kenka/mongox/examples/data/bsonx"

var (
	EmptyDoc                             = bsonx.NewDoc()
	DefaultDensifyOptions DensifyOptions = DensifyConstructibleBson{}.of(EmptyDoc)
)

type DensifyOptions interface {
	bsonx.Bson
	PartitionByFields(fields ...string) DensifyOptions
	Option(name string, value any) DensifyOptions
}
