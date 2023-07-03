package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/options"
	"go.mongodb.org/mongo-driver/bson"
)

type MergeStage struct {
	stage bsonx.Bson
}

func (s MergeStage) Bson() bsonx.Bson {
	return s.stage
}

func (s MergeStage) Document() bson.D {
	return s.stage.Document()
}

// Merge Writes the results of the aggregation pipeline to a specified collection. The
// $merge operator must be the last DefaultStage in the pipeline. The $merge DefaultStage: Can
// output to a collection in the same or different Database. Starting in MongoDB
// 4.4: $merge can output to the same collection that is being aggregated. For
// more information, see Output to the Same Collection that is Being Aggregated.
// Pipelines with the $merge DefaultStage can run on replica set secondary nodes if all
// the nodes in cluster have featureCompatibilityVersion set to 4.4 or higher
// and the Read Preference allows secondary reads. Read operations of the $merge
// statement are sent to secondary nodes, while the write operations occur only
// on the primary node. Not all driver versions support targeting of $merge
// operations to replica set secondary nodes. Check your driver documentation to
// see when your driver added support for $merge read operations running on
// secondary nodes. Creates a new collection if the output collection does not
// already exist. Can incorporate results (insert new documents, merge
// documents, replace documents, keep existing documents, fail the operation,
// process documents with a custom Update pipeline) into an existing collection.
// Can output to a sharded collection. Input collection can also be sharded. For
// a comparison with the $out DefaultStage which also outputs the aggregation results
// to a collection, see $merge and $out Comparison.
func Merge(collectionName string, options options.MergeOptions) MergeStage {
	return MergeStage{stage: NewMergeStage(bsonx.String(collectionName), options)}
}

// MergeWithNameSpace Writes the results of the aggregation pipeline to a specified collection. The
// $merge operator must be the last DefaultStage in the pipeline. The $merge DefaultStage: Can
// output to a collection in the same or different Database. Starting in MongoDB
// 4.4: $merge can output to the same collection that is being aggregated. For
// more information, see Output to the Same Collection that is Being Aggregated.
// Pipelines with the $merge DefaultStage can run on replica set secondary nodes if all
// the nodes in cluster have featureCompatibilityVersion set to 4.4 or higher
// and the Read Preference allows secondary reads. Read operations of the $merge
// statement are sent to secondary nodes, while the write operations occur only
// on the primary node. Not all driver versions support targeting of $merge
// operations to replica set secondary nodes. Check your driver documentation to
// see when your driver added support for $merge read operations running on
// secondary nodes. Creates a new collection if the output collection does not
// already exist. Can incorporate results (insert new documents, merge
// documents, replace documents, keep existing documents, fail the operation,
// process documents with a custom Update pipeline) into an existing collection.
// Can output to a sharded collection. Input collection can also be sharded. For
// a comparison with the $out DefaultStage which also outputs the aggregation results
// to a collection, see $merge and $out Comparison.
func MergeWithNameSpace(namespace options.MongoNamespace, options options.MergeOptions) MergeStage {
	return MergeStage{stage: NewMergeStage(
		bsonx.BsonDoc("db", bsonx.String(namespace.DatabaseName())).
			Append("coll", bsonx.String(namespace.CollectionName())),
		options,
	)}
}

type mergeStage struct {
	intoValue bsonx.IBsonValue
	options   options.MergeOptions
}

func NewMergeStage(intoValue bsonx.IBsonValue, options options.MergeOptions) mergeStage {
	return mergeStage{
		intoValue: intoValue,
		options:   options,
	}
}

func (f mergeStage) BsonDocument() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	data := bsonx.BsonEmpty()
	if f.intoValue.IsString() {
		data.Append("into", f.intoValue)
	} else {
		into := bsonx.BsonEmpty()
		into.Append("db", f.intoValue.AsDocument().GetString("db"))
		into.Append("coll", f.intoValue.AsDocument().GetString("coll"))
		data.Append("into", into)
	}

	if len(f.options.UniqueIdentifier()) > 0 {
		if len(f.options.UniqueIdentifier()) == 1 {
			data.Append("on", bsonx.String(f.options.UniqueIdentifier()[0]))
		} else {
			uniqueIdentifier := bsonx.Array()
			for _, s := range f.options.UniqueIdentifier() {
				uniqueIdentifier.Append(bsonx.String(s))
			}
			data.Append("on", uniqueIdentifier)
		}
	}
	if len(f.options.Variables()) > 0 {
		variables := bsonx.BsonEmpty()
		for _, s := range f.options.Variables() {
			variables.Append(s.GetName(), s.GetValue())
		}
		data.Append("let", variables)
	}

	if f.options.WhenMatched() != options.WhenMatchedInvalid {
		switch f.options.WhenMatched() {
		case options.WhenMatchedReplace, options.WhenMatchedKeepExisting, options.WhenMatchedMerge, options.WhenMatchedFail:
			data.Append("whenMatched", bsonx.String(options.WhenMatcheds[f.options.WhenMatched()]))
		case options.WhenMatchedPipeline:
			pipe := bsonx.Array()
			for _, m := range f.options.WhenMatchedPipeline() {
				pipe.Append(bsonx.Boolean(m))
			}
			data.Append("whenMatched", pipe)
		}
	}
	if f.options.WhenNotMatched() != options.WhenNotMatchedInvalid {
		data.Append("whenNotMatched", bsonx.String(options.WhenNotMatcheds[f.options.WhenNotMatched()]))
	}

	b.Append("$merge", data)
	return b
}

func (f mergeStage) Document() bson.D {
	return f.BsonDocument().Document()
}
