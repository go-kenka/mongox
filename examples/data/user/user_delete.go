package user

import (
	"context"
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDelete struct {
	cc *mongo.Collection
}

func (q UserDelete) DeleteMany(ctx context.Context, filter bsonx.Bson, opts ...*options.DeleteOptions) (int64, error) {
	result, err := q.cc.DeleteMany(ctx, filter.Document(), opts...)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (q UserDelete) DeleteOne(ctx context.Context, filter bsonx.Bson, opts ...*options.DeleteOptions) error {
	_, err := q.cc.DeleteOne(ctx, filter.Document(), opts...)
	if err != nil {
		return err
	}
	return nil
}
