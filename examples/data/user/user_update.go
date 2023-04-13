package user

import (
	"context"

	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/filters"
	"github.com/go-kenka/mongox/model/updates"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserUpdateMany struct {
	cc     *mongo.Collection
	filter bson.D
	update any
	opts   *options.UpdateOptions
}

func NewUserUpdateMany(cc *mongo.Collection) *UserUpdateMany {
	return &UserUpdateMany{cc: cc}
}

// SetFilter set filter
func (u *UserUpdateMany) SetFilter(f filters.Filter) *UserUpdateMany {
	u.filter = f.Document()
	return u
}

// SetUpdate set update doc
func (u *UserUpdateMany) SetUpdate(up updates.Update) *UserUpdateMany {
	u.update = up.Document()
	return u
}

// SetArrayFilters sets the value for the ArrayFilters field.
func (u *UserUpdateMany) SetArrayFilters(af options.ArrayFilters) *UserUpdateMany {
	u.opts.ArrayFilters = &af
	return u
}

// SetBypassDocumentValidation sets the value for the BypassDocumentValidation field.
func (u *UserUpdateMany) SetBypassDocumentValidation(b bool) *UserUpdateMany {
	u.opts.BypassDocumentValidation = &b
	return u
}

// SetCollation sets the value for the Collation field.
func (u *UserUpdateMany) SetCollation(c *options.Collation) *UserUpdateMany {
	u.opts.Collation = c
	return u
}

// SetComment sets the value for the Comment field.
func (u *UserUpdateMany) SetComment(comment interface{}) *UserUpdateMany {
	u.opts.Comment = comment
	return u
}

// SetHint sets the value for the Hint field.
func (u *UserUpdateMany) SetHint(h interface{}) *UserUpdateMany {
	u.opts.Hint = h
	return u
}

// SetUpsert sets the value for the Upsert field.
func (u *UserUpdateMany) SetUpsert(b bool) *UserUpdateMany {
	u.opts.Upsert = &b
	return u
}

// SetLet sets the value for the Let field.
func (u *UserUpdateMany) SetLet(l interface{}) *UserUpdateMany {
	u.opts.Let = l
	return u
}

func (u *UserUpdateMany) Save(ctx context.Context) (*mongo.UpdateResult, error) {
	return u.cc.UpdateMany(ctx, u.filter, u.update, u.opts)
}

type UserUpdateOne struct {
	cc     *mongo.Collection
	filter bson.D
	update any
	opts   *options.UpdateOptions
}

func NewUserUpdateOne(cc *mongo.Collection) *UserUpdateOne {
	return &UserUpdateOne{cc: cc}
}

// SetFilter set filter
func (u *UserUpdateOne) SetFilter(f filters.Filter) *UserUpdateOne {
	u.filter = f.Document()
	return u
}

// SetUpdate set update doc
func (u *UserUpdateOne) SetUpdate(up updates.Update) *UserUpdateOne {
	u.update = up.Document()
	return u
}

// SetArrayFilters sets the value for the ArrayFilters field.
func (u *UserUpdateOne) SetArrayFilters(af options.ArrayFilters) *UserUpdateOne {
	u.opts.ArrayFilters = &af
	return u
}

// SetBypassDocumentValidation sets the value for the BypassDocumentValidation field.
func (u *UserUpdateOne) SetBypassDocumentValidation(b bool) *UserUpdateOne {
	u.opts.BypassDocumentValidation = &b
	return u
}

// SetCollation sets the value for the Collation field.
func (u *UserUpdateOne) SetCollation(c *options.Collation) *UserUpdateOne {
	u.opts.Collation = c
	return u
}

// SetComment sets the value for the Comment field.
func (u *UserUpdateOne) SetComment(comment interface{}) *UserUpdateOne {
	u.opts.Comment = comment
	return u
}

// SetHint sets the value for the Hint field.
func (u *UserUpdateOne) SetHint(h interface{}) *UserUpdateOne {
	u.opts.Hint = h
	return u
}

// SetUpsert sets the value for the Upsert field.
func (u *UserUpdateOne) SetUpsert(b bool) *UserUpdateOne {
	u.opts.Upsert = &b
	return u
}

// SetLet sets the value for the Let field.
func (u *UserUpdateOne) SetLet(l interface{}) *UserUpdateOne {
	u.opts.Let = l
	return u
}

func (u *UserUpdateOne) Save(ctx context.Context) (*mongo.UpdateResult, error) {
	return u.cc.UpdateOne(ctx, u.filter, u.update, u.opts)
}

type UserUpdateOneID struct {
	cc     *mongo.Collection
	filter bson.D
	update any
	opts   *options.UpdateOptions
}

func NewUserUpdateOneID(id primitive.ObjectID, cc *mongo.Collection) *UserUpdateOneID {
	return &UserUpdateOneID{
		cc:     cc,
		filter: filters.Eq("_id", bsonx.ObjectId(id)).Document(),
	}
}

// SetUpdate set update doc
func (u *UserUpdateOneID) SetUpdate(up updates.Update) *UserUpdateOneID {
	u.update = up.Document()
	return u
}

// SetArrayFilters sets the value for the ArrayFilters field.
func (u *UserUpdateOneID) SetArrayFilters(af options.ArrayFilters) *UserUpdateOneID {
	u.opts.ArrayFilters = &af
	return u
}

// SetBypassDocumentValidation sets the value for the BypassDocumentValidation field.
func (u *UserUpdateOneID) SetBypassDocumentValidation(b bool) *UserUpdateOneID {
	u.opts.BypassDocumentValidation = &b
	return u
}

// SetCollation sets the value for the Collation field.
func (u *UserUpdateOneID) SetCollation(c *options.Collation) *UserUpdateOneID {
	u.opts.Collation = c
	return u
}

// SetComment sets the value for the Comment field.
func (u *UserUpdateOneID) SetComment(comment interface{}) *UserUpdateOneID {
	u.opts.Comment = comment
	return u
}

// SetHint sets the value for the Hint field.
func (u *UserUpdateOneID) SetHint(h interface{}) *UserUpdateOneID {
	u.opts.Hint = h
	return u
}

// SetUpsert sets the value for the Upsert field.
func (u *UserUpdateOneID) SetUpsert(b bool) *UserUpdateOneID {
	u.opts.Upsert = &b
	return u
}

// SetLet sets the value for the Let field.
func (u *UserUpdateOneID) SetLet(l interface{}) *UserUpdateOneID {
	u.opts.Let = l
	return u
}

func (u *UserUpdateOneID) Save(ctx context.Context) (*mongo.UpdateResult, error) {
	return u.cc.UpdateByID(ctx, u.filter, u.update, u.opts)
}

type UserReplaceOne struct {
	cc      *mongo.Collection
	filter  bson.D
	replace *UserData
	opts    *options.ReplaceOptions
}

func NewUserReplaceOne(cc *mongo.Collection) *UserReplaceOne {
	return &UserReplaceOne{cc: cc}
}

// SetFilter set filter
func (u *UserReplaceOne) SetFilter(f filters.Filter) *UserReplaceOne {
	u.filter = f.Document()
	return u
}

// SetReplace set replace doc
func (u *UserReplaceOne) SetReplace(replace *UserData) *UserReplaceOne {
	u.replace = replace
	return u
}

// SetBypassDocumentValidation sets the value for the BypassDocumentValidation field.
func (u *UserReplaceOne) SetBypassDocumentValidation(b bool) *UserReplaceOne {
	u.opts.BypassDocumentValidation = &b
	return u
}

// SetCollation sets the value for the Collation field.
func (u *UserReplaceOne) SetCollation(c *options.Collation) *UserReplaceOne {
	u.opts.Collation = c
	return u
}

// SetComment sets the value for the Comment field.
func (u *UserReplaceOne) SetComment(comment interface{}) *UserReplaceOne {
	u.opts.Comment = comment
	return u
}

// SetHint sets the value for the Hint field.
func (u *UserReplaceOne) SetHint(h interface{}) *UserReplaceOne {
	u.opts.Hint = h
	return u
}

// SetUpsert sets the value for the Upsert field.
func (u *UserReplaceOne) SetUpsert(b bool) *UserReplaceOne {
	u.opts.Upsert = &b
	return u
}

// SetLet sets the value for the Let field.
func (u *UserReplaceOne) SetLet(l interface{}) *UserReplaceOne {
	u.opts.Let = l
	return u
}

func (u *UserReplaceOne) Save(ctx context.Context) (*mongo.UpdateResult, error) {
	return u.cc.ReplaceOne(ctx, u.filter, u.replace, u.opts)
}
