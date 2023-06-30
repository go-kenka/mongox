package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/bsonx/expression"
	"github.com/go-kenka/mongox/bsonx/options"
	"go.mongodb.org/mongo-driver/bson"
)

type BucketAutoStage struct {
	stage bsonx.Bson
}

func (s BucketAutoStage) Bson() bsonx.Bson {
	return s.stage
}

func (s BucketAutoStage) Document() bson.D {
	return s.stage.Document()
}

// BucketAuto Categorizes incoming documents into a specific number of groups, called
// buckets, based on a specified expression. Bucket boundaries are automatically
// determined in an attempt to evenly distribute the documents into the
// specified number of buckets. Each bucket is represented as a document in the
// output. The document for each bucket contains: An _id object that specifies
// the bounds of the bucket. The _id.min field specifies the inclusive lower
// bound for the bucket. The _id.max field specifies the upper bound for the
// bucket. This bound is exclusive for all buckets except the final bucket in
// the series, where it is inclusive. A count field that contains the number of
// documents in the bucket. The count field is included by default when the
// output document is not specified. The $bucketAuto DefaultStage has the following
// form:
//
//	{
//	 $bucketAuto: {
//	     groupBy: <expression>,
//	     buckets: <number>,
//	     output: {
//	        <output1>: { <$accumulator expression> },
//	        ...
//	     }
//	     granularity: <string>
//	 }
//	}
func BucketAuto[T expression.AnyExpression](groupBy T, buckets int32, options options.BucketAutoOptions) BucketAutoStage {
	return BucketAutoStage{
		stage: NewBucketAutoStage(groupBy, buckets, options),
	}
}

type bucketAutoStage[T expression.AnyExpression] struct {
	groupBy T
	buckets int32
	options options.BucketAutoOptions
}

func NewBucketAutoStage[T expression.AnyExpression](groupBy T, buckets int32, options options.BucketAutoOptions) bucketAutoStage[T] {
	return bucketAutoStage[T]{
		groupBy: groupBy,
		buckets: buckets,
		options: options,
	}
}

func (bs bucketAutoStage[T]) BsonDocument() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonDoc("groupBy", bs.groupBy)

	data.Append("buckets", bsonx.Int32(bs.buckets))

	output := bs.options.GetOutPut()

	if len(output) > 0 {
		out := bsonx.BsonEmpty()
		for _, field := range output {
			out.Append(field.GetName(), field.GetValue().BsonDocument())
		}
		data.Append("output", out)
	}

	granularity := bs.options.GetGranularity()
	if len(granularity) > 0 {
		data.Append("granularity", bsonx.String(granularity))
	}

	b.Append("$bucketAuto", data)
	return b
}
func (bs bucketAutoStage[T]) Document() bson.D {
	return bs.BsonDocument().Document()
}
