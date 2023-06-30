package user

import (
	"context"

	"github.com/go-kenka/mongox/model/filters"
	"github.com/go-kenka/mongox/model/updates"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserQuery struct {
	cc        *mongo.Collection
	filter    bson.D
	update    any
	replace   *UserData
	fieldName string
}

func NewUserQuery(cc *mongo.Collection) *UserQuery {
	return &UserQuery{cc: cc}
}

// SetFilter set filter
func (q *UserQuery) SetFilter(f filters.Filter) *UserQuery {
	q.filter = f.Document()
	return q
}

// SetUpdate set update doc, use when FindOneAndUpdate
func (q *UserQuery) SetUpdate(up updates.Update) *UserQuery {
	q.update = up.Document()
	return q
}

// SetReplace set replace doc, use when FindOneAndReplace
func (q *UserQuery) SetReplace(replace *UserData) *UserQuery {
	q.replace = replace
	return q
}

// SetFieldName set distinct fieldName, use when Distinct
func (q *UserQuery) SetFieldName(replace *UserData) *UserQuery {
	q.replace = replace
	return q
}

func (q *UserQuery) Find(ctx context.Context, opts ...*options.FindOptions) (result []*UserData, err error) {
	cur, err := q.cc.Find(ctx, q.filter, opts...)
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

func (q *UserQuery) FindOne(ctx context.Context, opts ...*options.FindOneOptions) (result *UserData, err error) {
	err = q.cc.FindOne(ctx, q.filter, opts...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return
}

func (q *UserQuery) FindOneAndReplace(ctx context.Context, opts ...*options.FindOneAndReplaceOptions) (result *UserData, err error) {
	err = q.cc.FindOneAndReplace(ctx, q.filter, q.replace, opts...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return
}

func (q *UserQuery) FindOneAndUpdate(ctx context.Context, opts ...*options.FindOneAndUpdateOptions) (result *UserData, err error) {
	err = q.cc.FindOneAndUpdate(ctx, q.filter, q.update, opts...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return
}

func (q *UserQuery) FindOneAndDelete(ctx context.Context, opts ...*options.FindOneAndDeleteOptions) (result *UserData, err error) {
	err = q.cc.FindOneAndDelete(ctx, q.filter, opts...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return
}

func (q *UserQuery) CountDocuments(ctx context.Context, opts ...*options.CountOptions) (int64, error) {
	return q.cc.CountDocuments(ctx, q.filter, opts...)
}

func (q *UserQuery) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return q.cc.EstimatedDocumentCount(ctx, opts...)
}

func (q *UserQuery) Distinct(ctx context.Context, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return q.cc.Distinct(ctx, q.fieldName, q.filter, opts...)
}
