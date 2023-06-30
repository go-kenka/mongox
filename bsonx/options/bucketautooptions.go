package options

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/aggregates"
)

type BucketAutoOptions struct {
	output      []bsonx.BsonField
	granularity aggregates.BucketGranularity
}

func (b BucketAutoOptions) Granularity(granularity aggregates.BucketGranularity) BucketAutoOptions {
	b.granularity = granularity
	return b
}

func (b BucketAutoOptions) Output(output ...bsonx.BsonField) BucketAutoOptions {
	b.output = output
	return b
}

func (b BucketAutoOptions) GetGranularity() string {
	return b.granularity.GetValue()
}

func (b BucketAutoOptions) GetOutPut() []bsonx.BsonField {
	return b.output
}
