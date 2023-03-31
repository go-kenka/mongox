package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/filters"
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

	coll := client.Database("data_factory").Collection("movies")

	// doc := bsonx.BsonDoc("user", bsonx.BsonDoc("id", bsonx.String("01")).
	// 	Append("name", bsonx.String("xiaowu")).
	// 	Append("avatar", bsonx.String("https://www.mongodb.com/")).
	// 	Append("sex", bsonx.Boolean(true)),
	// ).Append("profile", bsonx.String("hhhhhh"))
	// doc := bsonx.BsonDoc("user", bsonx.BsonDoc("id", bsonx.String("03")).
	// 	Append("name", bsonx.String("xiaowu2")).
	// 	Append("avatar", bsonx.String("https://www.mongodb.com/")).
	// 	Append("sex", bsonx.Boolean(true)),
	// )

	doc := filters.And(
		filters.Eq("user", bsonx.BsonDoc("id", bsonx.String("02")).
			Append("name", bsonx.String("xiaowu1")).
			Append("avatar", bsonx.String("https://www.mongodb.com/")).
			Append("sex", bsonx.Boolean(true)),
		),
	)

	data := doc.Exp().AsDocument().Document()

	query, _ := json.Marshal(data)

	fmt.Println(string(query))

	var user map[string]any

	err = coll.FindOne(context.TODO(), data).Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
