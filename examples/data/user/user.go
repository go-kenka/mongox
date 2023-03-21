package user

import (
	"context"
	"github.com/go-kenka/mongox/examples/data/bsonx"
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

type UserData struct {
	Id         primitive.ObjectID   `bson:"_id"`
	UserName   string               `bson:"user_name"`
	DateTime   primitive.DateTime   `bson:"DateTime"`
	Decimal128 primitive.Decimal128 `bson:"Decimal128"`
	Timestamp  primitive.Timestamp  `bson:"Timestamp"`
	Binary     primitive.Binary     `bson:"Binary"`
}

func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		cc: c.cc,
	}
}

func (c *UserClient) Create() *UserCreate {
	return &UserCreate{
		cc: c.cc,
	}
}

func (c *UserClient) Update() *UserUpdate {
	return &UserUpdate{
		cc: c.cc,
	}
}

func (c *UserClient) Delete() *UserDelete {
	return &UserDelete{
		cc: c.cc,
	}
}

func (c *UserClient) Aggregate() *UserAggregate {
	return &UserAggregate{
		cc: c.cc,
	}
}

func (c *UserClient) Drop(ctx context.Context) error {
	return c.cc.Drop(ctx)
}

func (c *UserClient) Watch(ctx context.Context, pipeline bsonx.Bson, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return c.cc.Watch(ctx, pipeline.Document(), opts...)
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
