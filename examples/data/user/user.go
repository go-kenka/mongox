package user

import (
	"context"

	"github.com/go-kenka/mongox/model/aggregates/watch"
	"github.com/go-kenka/mongox/model/bulks"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DatabaseName   = "db1"
	CollectionName = "user"
	FieldId        = "_id"
	FieldUserName  = "user_name"
)

type UserClient struct {
	cc *mongo.Collection
}

func NewUserClient(db *mongo.Client) *UserClient {
	return &UserClient{
		cc: db.Database(DatabaseName).Collection(CollectionName),
	}
}

//
// Double           MongoType = 0x01
// String           MongoType = 0x02
// EmbeddedDocument MongoType = 0x03
// Array            MongoType = 0x04
// Binary           MongoType = 0x05
// ObjectID         MongoType = 0x07
// Boolean          MongoType = 0x08
// DateTime         MongoType = 0x09
// Regex            MongoType = 0x0B
// DBPointer        MongoType = 0x0C
// JavaScript       MongoType = 0x0D
// Symbol           MongoType = 0x0E
// CodeWithScope    MongoType = 0x0F
// Int32            MongoType = 0x10
// Timestamp        MongoType = 0x11
// Int64            MongoType = 0x12
// Decimal128       MongoType = 0x13

// UserData .
type UserData struct {
	ObjectID         primitive.ObjectID     `bson:"_id"`
	Double           float64                `bson:"double"`
	String           string                 `bson:"string"`
	EmbeddedDocument *EmbeddedDocument_Data `bson:"embedded_document"`
	Array            []*Array_Data          `bson:"array"`
	Binary           primitive.Binary       `bson:"binary"`
	Boolean          bool                   `bson:"boolean"`
	DateTime         primitive.DateTime     `bson:"date_time"`
	Regex            primitive.Regex        `bson:"regex"`
	JavaScript       primitive.JavaScript   `bson:"java_script"`
	Int32            int32                  `bson:"int32"`
	Timestamp        primitive.Timestamp    `bson:"timestamp"`
	Int64            int64                  `bson:"int64"`
	Decimal128       primitive.Decimal128   `bson:"decimal128"`
}

func (d UserData) name() {

}

type EmbeddedDocument_Data struct {
	Double           float64                `bson:"double"`
	String           string                 `bson:"string"`
	EmbeddedDocument *EmbeddedDocument_Data `bson:"embedded_document"`
	Array            []*Array_Data          `bson:"array"`
	Binary           primitive.Binary       `bson:"binary"`
	Boolean          bool                   `bson:"boolean"`
	DateTime         primitive.DateTime     `bson:"date_time"`
	Regex            primitive.Regex        `bson:"regex"`
	JavaScript       primitive.JavaScript   `bson:"java_script"`
	Int32            int32                  `bson:"int32"`
	Timestamp        primitive.Timestamp    `bson:"timestamp"`
	Int64            int64                  `bson:"int64"`
	Decimal128       primitive.Decimal128   `bson:"decimal128"`
}

type Array_Data struct {
	Double           float64                `bson:"double"`
	String           string                 `bson:"string"`
	EmbeddedDocument *EmbeddedDocument_Data `bson:"embedded_document"`
	Array            []*Array_Data          `bson:"array"`
	Binary           primitive.Binary       `bson:"binary"`
	Boolean          bool                   `bson:"boolean"`
	DateTime         primitive.DateTime     `bson:"date_time"`
	Regex            primitive.Regex        `bson:"regex"`
	JavaScript       primitive.JavaScript   `bson:"java_script"`
	Int32            int32                  `bson:"int32"`
	Timestamp        primitive.Timestamp    `bson:"timestamp"`
	Int64            int64                  `bson:"int64"`
	Decimal128       primitive.Decimal128   `bson:"decimal128"`
}

func (c *UserClient) Query() *UserQuery {
	return NewUserQuery(c.cc)
}

func (c *UserClient) Create() *UserCreate {
	return NewUserCreate(c.cc)
}

func (c *UserClient) UpdateMany() *UserUpdateMany {
	return NewUserUpdateMany(c.cc)
}

func (c *UserClient) UpdateOne() *UserUpdateOne {
	return NewUserUpdateOne(c.cc)
}

func (c *UserClient) UpdateOneID(id primitive.ObjectID) *UserUpdateOneID {
	return NewUserUpdateOneID(id, c.cc)
}

func (c *UserClient) ReplaceOne() *UserReplaceOne {
	return NewUserReplaceOne(c.cc)
}

func (c *UserClient) DeleteMany() *UserDeleteMany {
	return NewUserDeleteMany(c.cc)
}

func (c *UserClient) DeleteOne() *UserDeleteOne {
	return NewUserDeleteOne(c.cc)
}

func (c *UserClient) Aggregate() *UserAggregate {
	return NewUserAggregate(c.cc)
}

func (c *UserClient) Drop(ctx context.Context) error {
	return c.cc.Drop(ctx)
}

func (c *UserClient) Watch(ctx context.Context, pipeline watch.WatchPipe, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return c.cc.Watch(ctx, pipeline.Pipe(), opts...)
}

func (c *UserClient) Indexes() mongo.IndexView {
	return c.cc.Indexes()
}

func (c *UserClient) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	return c.cc.Clone(opts...)
}

func (c *UserClient) Name() string {
	return c.cc.Name()
}

func (c *UserClient) Database() *mongo.Database {
	return c.cc.Database()
}

func (c *UserClient) BulkWrite(ctx context.Context, mds []bulks.BulkModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	var models []mongo.WriteModel
	for _, stage := range mds {
		models = append(models, stage.WriteModel())
	}
	return c.cc.BulkWrite(ctx, models, opts...)
}
