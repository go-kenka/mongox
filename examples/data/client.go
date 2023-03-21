package data

import (
	"context"
	"github.com/go-kenka/mongox/examples/data/user"
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
