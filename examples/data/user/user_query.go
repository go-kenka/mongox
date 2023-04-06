package user

import (
	"context"

	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/aggregates"
	"github.com/go-kenka/mongox/model/filters"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserQuery struct {
	cc *mongo.Collection
}

func test() {
	q := UserQuery{cc: nil}

	q.Find(context.TODO(), filters.Eq("aaa", bsonx.String("000")))
}

func (q UserQuery) Find(ctx context.Context, filter filters.Filter, opts ...*options.FindOptions) (result []*UserData, err error) {
	cur, err := q.cc.Find(ctx, filter.Document(), opts...)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return
}

func (q UserQuery) FindOne(ctx context.Context, filter filters.Filter, opts ...*options.FindOneOptions) (result *UserData, err error) {
	err = q.cc.FindOne(ctx, filter.Document(), opts...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return
}

func (q UserQuery) FindOneAndReplace(ctx context.Context, filter filters.Filter, replacement bsonx.Bson, opts ...*options.FindOneAndReplaceOptions) (result *UserData, err error) {
	err = q.cc.FindOneAndReplace(ctx, filter.Document(), replacement.Document(), opts...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return
}

func (q UserQuery) FindOneAndUpdate(ctx context.Context, filter filters.Filter, update bsonx.Bson, opts ...*options.FindOneAndUpdateOptions) (result *UserData, err error) {
	err = q.cc.FindOneAndUpdate(ctx, filter.Document(), update.Document(), opts...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return
}

func (q UserQuery) FindOneAndDelete(ctx context.Context, filter filters.Filter, opts ...*options.FindOneAndDeleteOptions) (result *UserData, err error) {
	err = q.cc.FindOneAndDelete(ctx, filter.Document(), opts...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return
}

func (q UserQuery) CountDocuments(ctx context.Context, filter filters.Filter, opts ...*options.CountOptions) (int64, error) {
	return q.cc.CountDocuments(ctx, filter.Document(), opts...)
}

func (q UserQuery) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return q.cc.EstimatedDocumentCount(ctx, opts...)
}

func (q UserQuery) Distinct(ctx context.Context, fieldName string, filter filters.Filter, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return q.cc.Distinct(ctx, fieldName, filter.Document(), opts...)
}

type UserAggregate struct {
	cc *mongo.Collection
}

func NewUserAggregate() *UserAggregate {
	return &UserAggregate{}
}

func (q UserAggregate) Save(ctx context.Context, pipe []aggregates.Stage, opts ...*options.AggregateOptions) error {
	var stages mongo.Pipeline
	for _, stage := range pipe {
		stages = append(stages, stage.Document())
	}
	_, err := q.cc.Aggregate(ctx, stages, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (q UserAggregate) All(ctx context.Context, pipe []aggregates.Stage, result any, opts ...*options.AggregateOptions) (err error) {
	var stages mongo.Pipeline
	for _, stage := range pipe {
		stages = append(stages, stage.Document())
	}
	cursor, err := q.cc.Aggregate(ctx, stages, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &result)
	if err != nil {
		return err
	}
	return
}

func (q UserAggregate) Get(ctx context.Context, pipe []aggregates.Stage, result any, opts ...*options.AggregateOptions) (err error) {
	var stages mongo.Pipeline
	for _, stage := range pipe {
		stages = append(stages, stage.Document())
	}
	cursor, err := q.cc.Aggregate(ctx, stages, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	if cursor.Next(ctx) {
		err = cursor.Decode(&result)
		if err != nil {
			return err
		}
		return
	}
	return
}
