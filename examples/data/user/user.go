package user

import (
	"context"
	"errors"

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
	c  *mongo.Client
	db string
}

func NewUserClient(db *mongo.Client) *UserClient {
	return &UserClient{
		c: db,
	}
}

// UserData .
type UserData struct {
	ID               primitive.ObjectID     `bson:"_id"`
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

func (d UserData) Document() any {
	return d
}
func (d UserData) Update() {

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

func (c *UserClient) DBName(database string) *UserClient {
	c.db = database
	return c
}

func (c *UserClient) collection() *mongo.Collection {
	if len(c.db) == 0 {
		panic(errors.New("db not set"))
	}

	return c.c.Database(c.db).Collection(CollectionName)
}

func (c *UserClient) Query() *UserQuery {
	return NewUserQuery(c.collection())
}

func (c *UserClient) Create() *UserCreate {
	return NewUserCreate(c.collection())
}

func (c *UserClient) UpdateMany() *UserUpdateMany {
	return NewUserUpdateMany(c.collection())
}

func (c *UserClient) UpdateOne() *UserUpdateOne {
	return NewUserUpdateOne(c.collection())
}

func (c *UserClient) UpdateOneID(id primitive.ObjectID) *UserUpdateOneID {
	return NewUserUpdateOneID(id, c.collection())
}

func (c *UserClient) ReplaceOne() *UserReplaceOne {
	return NewUserReplaceOne(c.collection())
}

func (c *UserClient) DeleteMany() *UserDeleteMany {
	return NewUserDeleteMany(c.collection())
}

func (c *UserClient) DeleteOne() *UserDeleteOne {
	return NewUserDeleteOne(c.collection())
}

func (c *UserClient) Aggregate() *UserAggregate {
	return NewUserAggregate(c.collection())
}

func (c *UserClient) Drop(ctx context.Context) error {
	return c.collection().Drop(ctx)
}

func (c *UserClient) Watch(ctx context.Context, pipeline watch.WatchPipe, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return c.collection().Watch(ctx, pipeline.Pipe(), opts...)
}

func (c *UserClient) Indexes() mongo.IndexView {
	return c.collection().Indexes()
}

func (c *UserClient) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	return c.collection().Clone(opts...)
}

func (c *UserClient) Name() string {
	return c.collection().Name()
}

func (c *UserClient) Database() *mongo.Database {
	return c.collection().Database()
}

func (c *UserClient) BulkWrite(ctx context.Context, mds []bulks.BulkModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	var models []mongo.WriteModel
	for _, stage := range mds {
		models = append(models, stage.WriteModel())
	}
	return c.collection().BulkWrite(ctx, models, opts...)
}
