package user

import (
	"context"
	"time"

	"github.com/go-kenka/mongox/model/aggregates"
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

type UserAggregate struct {
	cc   *mongo.Collection
	pipe mongo.Pipeline
	opts *options.AggregateOptions
}

func NewUserAggregate(cc *mongo.Collection) *UserAggregate {
	return &UserAggregate{cc: cc}
}

// SetAllowDiskUse sets the value for the AllowDiskUse field.
func (a *UserAggregate) SetAllowDiskUse(b bool) *UserAggregate {
	a.opts.AllowDiskUse = &b
	return a
}

// SetBatchSize sets the value for the BatchSize field.
func (a *UserAggregate) SetBatchSize(i int32) *UserAggregate {
	a.opts.BatchSize = &i
	return a
}

// SetBypassDocumentValidation sets the value for the BypassDocumentValidation field.
func (a *UserAggregate) SetBypassDocumentValidation(b bool) *UserAggregate {
	a.opts.BypassDocumentValidation = &b
	return a
}

// SetCollation sets the value for the Collation field.
func (a *UserAggregate) SetCollation(c *options.Collation) *UserAggregate {
	a.opts.Collation = c
	return a
}

// SetMaxTime sets the value for the MaxTime field.
//
// NOTE(benjirewis): MaxTime will be deprecated in a future release. The more general Timeout
// option may be used in its place to control the amount of time that a single operation can
// run before returning an error. MaxTime is ignored if Timeout is set on the client.
func (a *UserAggregate) SetMaxTime(d time.Duration) *UserAggregate {
	a.opts.MaxTime = &d
	return a
}

// SetMaxAwaitTime sets the value for the MaxAwaitTime field.
func (a *UserAggregate) SetMaxAwaitTime(d time.Duration) *UserAggregate {
	a.opts.MaxAwaitTime = &d
	return a
}

// SetComment sets the value for the Comment field.
func (a *UserAggregate) SetComment(s string) *UserAggregate {
	a.opts.Comment = &s
	return a
}

// SetHint sets the value for the Hint field.
func (a *UserAggregate) SetHint(h interface{}) *UserAggregate {
	a.opts.Hint = h
	return a
}

// SetLet sets the value for the Let field.
func (a *UserAggregate) SetLet(let interface{}) *UserAggregate {
	a.opts.Let = let
	return a
}

// SetCustom sets the value for the Custom field. Key-value pairs of the BSON map should correlate
// with desired option names and values. Values must be Marshalable. Custom options may conflict
// with non-custom options, and custom options bypass client-side validation. Prefer using non-custom
// options where possible.
func (a *UserAggregate) SetCustom(c bson.M) *UserAggregate {
	a.opts.Custom = c
	return a
}

// SetPipe set aggregate pipe
func (a *UserAggregate) SetPipe(pipe aggregates.Pipe) *UserAggregate {
	a.pipe = pipe.Pipe()
	return a
}

// Save only save result,such as: $out,$merge
func (a *UserAggregate) Save(ctx context.Context) error {
	_, err := a.cc.Aggregate(ctx, a.pipe, a.opts)
	if err != nil {
		return err
	}
	return nil
}

// Find get result as slice
func (a *UserAggregate) Find(ctx context.Context, val []any) (err error) {
	cursor, err := a.cc.Aggregate(ctx, a.pipe, a.opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &val)
	if err != nil {
		return err
	}
	return
}

// FindOne get one data
func (a *UserAggregate) FindOne(ctx context.Context, val interface{}) (err error) {
	cursor, err := a.cc.Aggregate(ctx, a.pipe, a.opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	if cursor.Next(ctx) {
		err = cursor.Decode(&val)
		if err != nil {
			return err
		}
		return
	}
	return
}
