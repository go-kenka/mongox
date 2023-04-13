package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserCreate struct {
	cc *mongo.Collection
}

func NewUserCreate(cc *mongo.Collection) *UserCreate {
	return &UserCreate{cc: cc}
}

func (q *UserCreate) InsertOne(ctx context.Context, doc UserData, opts ...*options.InsertOneOptions) (*UserData, error) {
	doc.Id = primitive.NewObjectID()
	_, err := q.cc.InsertOne(ctx, doc, opts...)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func (q *UserCreate) InsertMany(ctx context.Context, docs []UserData, opts ...*options.InsertManyOptions) ([]UserData, error) {
	for _, doc := range docs {
		doc.Id = primitive.NewObjectID()
	}

	_, err := q.cc.InsertMany(ctx, toSlice(docs), opts...)
	if err != nil {
		return nil, err
	}

	return docs, err
}

func toSlice(docs []UserData) (v []interface{}) {
	for _, doc := range docs {
		v = append(v, doc)
	}
	return
}
