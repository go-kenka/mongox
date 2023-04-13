package user

import (
	"context"

	"github.com/go-kenka/mongox/model/filters"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDeleteMany struct {
	cc     *mongo.Collection
	filter bson.D
	opts   *options.DeleteOptions
}

func NewUserDeleteMany(cc *mongo.Collection) *UserDeleteMany {
	return &UserDeleteMany{
		cc: cc,
	}
}

// SetFilter set filter
func (u *UserDeleteMany) SetFilter(f filters.Filter) *UserDeleteMany {
	u.filter = f.Document()
	return u
}

// SetCollation sets the value for the Collation field.
func (u *UserDeleteMany) SetCollation(c *options.Collation) *UserDeleteMany {
	u.opts.Collation = c
	return u
}

// SetComment sets the value for the Comment field.
func (u *UserDeleteMany) SetComment(comment interface{}) *UserDeleteMany {
	u.opts.Comment = comment
	return u
}

// SetHint sets the value for the Hint field.
func (u *UserDeleteMany) SetHint(h interface{}) *UserDeleteMany {
	u.opts.Hint = h
	return u
}

// SetLet sets the value for the Let field.
func (u *UserDeleteMany) SetLet(l interface{}) *UserDeleteMany {
	u.opts.Let = l
	return u
}

func (u *UserDeleteMany) Save(ctx context.Context) (*mongo.DeleteResult, error) {
	return u.cc.DeleteMany(ctx, u.filter, u.opts)
}

type UserDeleteOne struct {
	cc     *mongo.Collection
	filter bson.D
	opts   *options.DeleteOptions
}

func NewUserDeleteOne(cc *mongo.Collection) *UserDeleteOne {
	return &UserDeleteOne{
		cc: cc,
	}
}

// SetFilter set filter
func (u *UserDeleteOne) SetFilter(f filters.Filter) *UserDeleteOne {
	u.filter = f.Document()
	return u
}

// SetCollation sets the value for the Collation field.
func (u *UserDeleteOne) SetCollation(c *options.Collation) *UserDeleteOne {
	u.opts.Collation = c
	return u
}

// SetComment sets the value for the Comment field.
func (u *UserDeleteOne) SetComment(comment interface{}) *UserDeleteOne {
	u.opts.Comment = comment
	return u
}

// SetHint sets the value for the Hint field.
func (u *UserDeleteOne) SetHint(h interface{}) *UserDeleteOne {
	u.opts.Hint = h
	return u
}

// SetLet sets the value for the Let field.
func (u *UserDeleteOne) SetLet(l interface{}) *UserDeleteOne {
	u.opts.Let = l
	return u
}

func (u *UserDeleteOne) Save(ctx context.Context) (*mongo.DeleteResult, error) {
	return u.cc.DeleteOne(ctx, u.filter, u.opts)
}
