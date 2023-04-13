package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:YioP9hcin2Pwe2pn@192.168.0.201:27017"

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// m := data.NewClient(client)
	//
	// m.WithTransaction(context.TODO(), func(sessCtx mongo.SessionContext) (interface{}, error) {
	// 	m.User.Update().UpdateOne(sessCtx, nil, nil)
	// 	return nil, nil
	// })

	data := map[string]interface{}{
		"min":       primitive.MinKey{},
		"max":       primitive.MaxKey{},
		"null":      primitive.Null{},
		"undefined": primitive.Undefined{},
		"timestamp": primitive.Timestamp{},
		"DBPointer": primitive.DBPointer{
			DB:      "aaa",
			Pointer: primitive.NewObjectID(),
		},
		"CodeWithScope": primitive.CodeWithScope{
			Code:  "aaa",
			Scope: map[string]interface{}{},
		},
	}

	_, err = client.Database("data_factory").Collection("logs").InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
}
