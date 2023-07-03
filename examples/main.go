package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kenka/mongox/bsonx"
	data "github.com/go-kenka/mongox/examples/simple"
	"github.com/go-kenka/mongox/examples/simple/user"
	"github.com/go-kenka/mongox/model/aggregates"
	"github.com/go-kenka/mongox/model/filters"
	"github.com/go-kenka/mongox/model/updates"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const uri = "mongodb://root:YioP9hcin2Pwe2pn@192.168.0.201:27017"

func main() {

	client, err := data.Open(uri)
	if err != nil {
		panic(err)
	}

	logid, _ := primitive.ObjectIDFromHex("6347be588b4468d85edfc314")

	n, err := client.User.Database("data_factory").Create().InsertOne(context.Background(), user.UserData{
		Id:     primitive.NewObjectID(),
		String: "xxxx01",
		Bool:   false,
		Binary: primitive.Binary{
			Subtype: bsontype.BinaryUUID,
			Data:    uuid.New().NodeID(),
		},
		Double:   10.22,
		DataTime: primitive.NewDateTimeFromTime(time.Now()),
		CreateAt: primitive.Timestamp{},
		Pointer: primitive.DBPointer{
			DB:      "logs",
			Pointer: logid,
		},
		Decimal: primitive.NewDecimal128(1205, 0),
		Int32:   151412,
		Int64:   112121,
		Js:      "{}",
		JsScope: primitive.CodeWithScope{
			Code:  "{}",
			Scope: map[string]string{},
		},
		Regular: primitive.Regex{
			Pattern: "/",
			Options: "m",
		},
		Int32s:                []int32{12, 132, 156},
		Int64s:                []int64{122, 1322, 1562},
		Strings:               []string{"12", "132", "156"},
		Floats:                []float64{122.0023, 1322.12, 1562.2},
		ArrayAny:              []any{122.0023, "aaa", 12, 1322.12, 1562.2, primitive.NewDateTimeFromTime(time.Now())},
		ArrayobjectSimpleList: nil,
		ObjectSimple:          nil,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(n.Id.Hex())
	doc := filters.And(
		filters.Eq("string", bsonx.String("xxxx01")),
	)

	// 多条查询
	result, err := client.User.Database("data_factory").Query().SetFilter(doc).Find(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("多条", result)
	// 单条查询
	one, err := client.User.Database("data_factory").Query().SetFilter(doc).FindOne(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("单条", one)

	// 单条更新(id)
	_, err = client.User.Database("data_factory").UpdateOneID(one.Id).
		SetUpdate(updates.Combine(updates.Set(user.FieldBool, bsonx.Boolean(true)))).
		Save(context.Background())
	if err != nil {
		panic(err)
	}
	// 单条更新
	_, err = client.User.Database("data_factory").UpdateOne().SetFilter(doc).
		SetUpdate(updates.Combine(updates.Set(user.FieldBool, bsonx.Boolean(false)))).
		Save(context.Background())
	if err != nil {
		panic(err)
	}
	// 多条更新
	_, err = client.User.Database("data_factory").UpdateMany().SetFilter(doc).
		SetUpdate(updates.Combine(updates.Set(user.FieldBool, bsonx.Boolean(true)))).
		Save(context.Background())
	if err != nil {
		panic(err)
	}

	// 聚合查询
	var userData []*user.UserData
	err = client.User.Database("data_factory").Aggregate().SetPipe(aggregates.NewPipe(
		aggregates.Match(doc),
	)).Find(context.Background(), &userData)
	if err != nil {
		panic(err)
	}

	// 删除一条
	_, err = client.User.Database("data_factory").DeleteOne().SetFilter(doc).
		Save(context.Background())
	if err != nil {
		panic(err)
	}

	// 删除多条
	_, err = client.User.Database("data_factory").DeleteMany().SetFilter(doc).
		Save(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(userData)
}
