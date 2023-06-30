package user

import (
	"context"
	"time"

	"github.com/go-kenka/mongox/model/aggregates"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
