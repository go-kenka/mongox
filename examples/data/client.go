package data

import (
	"context"

	"github.com/go-kenka/mongox/examples/data/user"
	"github.com/go-kenka/mongox/model/aggregates/watch"
	"github.com/go-kenka/mongox/model/filters"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client .
type Client struct {
	DB   *mongo.Client
	User *user.UserClient
}

// NewClient .
func NewClient(db *mongo.Client) *Client {
	return &Client{
		DB:   db,
		User: user.NewUserClient(db),
	}
}

// Open .
func Open(uri string) (*Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return NewClient(client), nil
}

func (c *Client) Watch(ctx context.Context, pipeline watch.WatchPipe, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return c.DB.Watch(ctx, pipeline.Pipe(), opts...)
}

func (c *Client) ListDatabaseNames(ctx context.Context, filter filters.Filter, opts ...*options.ListDatabasesOptions) ([]string, error) {
	return c.DB.ListDatabaseNames(ctx, filter.Document(), opts...)
}

func (c *Client) ListDatabases(ctx context.Context, filter filters.Filter, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error) {
	return c.DB.ListDatabases(ctx, filter.Document(), opts...)
}

func (c *Client) UseSession(ctx context.Context, opts *options.SessionOptions, fn func(sessionContext mongo.SessionContext) error) error {
	return c.DB.UseSessionWithOptions(ctx, opts, fn)
}

func (c *Client) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	session, err := c.DB.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	return session.WithTransaction(ctx, fn, opts...)
}
