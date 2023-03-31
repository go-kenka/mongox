package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"go.mongodb.org/mongo-driver/bson"
)

type BucketStage Stage

// Bucket Categorizes incoming documents into groups, called buckets, based on a
// specified expression and bucket boundaries and outputs a document per each
// bucket. Each output document contains an _id field whose value specifies the
// inclusive lower bound of the bucket. The output option specifies the fields
// included in each output document. $bucket only produces output documents for
// buckets that contain at least one input document.
func Bucket[T expression.AnyExpression, B expression.NumberExpression](groupBy T, boundaries []B, options AggBucketOptions) BucketStage {
	return NewBucketStage(groupBy, boundaries, options)
}

type bucketStage[T expression.AnyExpression, B expression.NumberExpression] struct {
	groupBy    T
	boundaries []B
	options    AggBucketOptions
}

func (bs bucketStage[T, B]) Bson() bsonx.Bson {
	return bs.ToBsonDocument()
}

func NewBucketStage[T expression.AnyExpression, B expression.NumberExpression](groupBy T, boundaries []B, options AggBucketOptions) bucketStage[T, B] {
	return bucketStage[T, B]{
		groupBy:    groupBy,
		boundaries: boundaries,
		options:    options,
	}
}

func (bs bucketStage[T, B]) ToBsonDocument() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonDoc("groupBy", bs.groupBy)

	var boundaries bsonx.BsonArray
	for _, boundary := range bs.boundaries {
		boundaries.Append(boundary)
	}

	data.Append("boundaries", boundaries)

	defaultBucket := bs.options.GetDefaultBucket()
	if defaultBucket != "" {
		data.Append("default", bsonx.String(defaultBucket))
	}

	output := bs.options.GetOutPut()

	if len(output) > 0 {
		out := bsonx.BsonEmpty()
		for _, field := range output {
			out.Append(field.GetName(), field.GetValue().ToBsonDocument())
		}
		data.Append("output", out)
	}

	b.Append("$bucket", data)
	return b
}
func (bs bucketStage[T, B]) Document() bson.D {
	return bs.ToBsonDocument().Document()
}
