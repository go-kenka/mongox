package aggregates

import "github.com/go-kenka/mongox/examples/data/bsonx"

type AggBucketOptions struct {
	defaultBucket string
	output        []bsonx.BsonField
}

func (b AggBucketOptions) DefaultBucket(name string) AggBucketOptions {
	b.defaultBucket = name
	return b
}

func (b AggBucketOptions) Output(output ...bsonx.BsonField) AggBucketOptions {
	b.output = output
	return b
}

func (b AggBucketOptions) GetDefaultBucket() string {
	return b.defaultBucket
}

func (b AggBucketOptions) GetOutPut() []bsonx.BsonField {
	return b.output
}
