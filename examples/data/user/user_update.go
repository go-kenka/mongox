package user

import (
	"context"

	bsonx2 "github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserUpdate struct {
	cc *mongo.Collection
}

func (u UserUpdate) UpdateMany(ctx context.Context, filter, update bsonx2.Bson, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return u.cc.UpdateMany(ctx, filter.Document(), update.Document(), opts...)
}

func (u UserUpdate) UpdateOne(ctx context.Context, filter, update bsonx2.Bson, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return u.cc.UpdateOne(ctx, filter.Document(), update.Document(), opts...)
}

func (u UserUpdate) UpdateByID(ctx context.Context, id bsonx2.BsonObjectId, update bsonx2.Bson, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return u.cc.UpdateByID(ctx, id.Get(), update.Document(), opts...)
}

func (u UserUpdate) ReplaceOne(ctx context.Context, filter, replacement bsonx2.Bson, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return u.cc.ReplaceOne(ctx, filter.Document(), replacement.Document(), opts...)
}

func (u UserUpdate) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return u.cc.BulkWrite(ctx, models, opts...)
}
