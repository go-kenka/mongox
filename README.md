# `Mongox`是一个简单的mongodb代码生成工具

通过内置的`dsl`语法，实现数据库`mongo`和实体之间的映射和相关`CRUD`的实现
实现本身通过解析`dsl`语法树，解析数据表的相关定义，然后使用`template`实现模板的生成。

# 快速开始
```shell
go install github.com/go-kenka/mongox/cmd/mongox@latest
```

# 生成schema文件
```shell
mongox init user --target ./data
```

## 示例
```go
package schema

import (
	"time"

	"github.com/go-kenka/mongox"
	"github.com/go-kenka/mongox/schema/field"
	"github.com/go-kenka/mongox/types"
)

type User struct {
	mongox.Schema
}

func (User) Fields() []mongox.Field {
	return []mongox.Field{
		field.ObjectId("_id").Tag(`json:"aaaa"`).Optional(),
		field.String("string").Tag(`json:"s1"`).Optional(),
		field.Bool("bool").Tag(`json:"s1"`).Optional(),
		field.Bytes("binary").Tag(`json:"b1"`).Optional(),
		field.Float("double").Tag(`json:"d1"`).Optional(),
		field.Time("data_time").Tag(`json:"aaaa"`).Optional(),
		field.Timestamp("create_at").Tag(`json:"aaaa"`).Optional(),
		field.DbPointer("pointer").Tag(`json:"aaaa"`).Optional(),
		field.Decimal128("decimal").Tag(`json:"aaaa"`).Optional(),
		field.Int32("int32").Tag(`json:"aaaa"`).Optional(),
		field.Int64("int64").Tag(`json:"aaaa"`).Optional(),
		field.JavaScript("js").Tag(`json:"aaaa"`).Optional(),
		field.JavaScriptScope("js_scope").Tag(`json:"aaaa"`).Optional(),
		field.Regular("regular").Tag(`json:"aaaa"`).Optional(),
		field.Int32s("int32s").Tag(`json:"aaaa"`).Optional(),
		field.Int64s("int64s").Tag(`json:"aaaa"`).Optional(),
		field.Strings("strings").Tag(`json:"aaaa"`).Optional(),
		field.Floats("floats").Tag(`json:"aaaa"`).Optional(),
		field.Arrays("array_any").Tag(`json:"aaaa"`),
		field.Arrays("arrayobject_simple").Tag(`json:"aaaa"`).Attributes(
			field.ObjectId("_id").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.String("string").Tag(`json:"s1"`).Optional().Descriptor(),
			field.Bool("bool").Tag(`json:"s1"`).Optional().Descriptor(),
			field.Bytes("binary").Tag(`json:"b1"`).Optional().Descriptor(),
			field.Float("double").Tag(`json:"d1"`).Optional().Descriptor(),
			field.Time("data_time").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Timestamp("create_at").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.DbPointer("pointer").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Decimal128("decimal").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Int32("int32").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Int64("int64").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.JavaScript("js").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.JavaScriptScope("js_scope").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Regular("regular").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Int32s("int32s").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Int64s("int64s").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Strings("strings").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Floats("floats").Tag(`json:"aaaa"`).Optional().Descriptor(),
		).Optional(),
		field.Object("object_simple").Tag(`json:"aaaa"`).Attributes(
			field.ObjectId("_id").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.String("string").Tag(`json:"s1"`).Optional().Descriptor(),
			field.Bool("bool").Tag(`json:"s1"`).Optional().Descriptor(),
			field.Bytes("binary").Tag(`json:"b1"`).Optional().Descriptor(),
			field.Float("double").Tag(`json:"d1"`).Optional().Descriptor(),
			field.Time("data_time").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Timestamp("create_at").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.DbPointer("pointer").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Decimal128("decimal").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Int32("int32").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Int64("int64").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.JavaScript("js").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.JavaScriptScope("js_scope").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Regular("regular").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Int32s("int32s").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Int64s("int64s").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Strings("strings").Tag(`json:"aaaa"`).Optional().Descriptor(),
			field.Floats("floats").Tag(`json:"aaaa"`).Optional().Descriptor(),
		).Optional(),
	}
}
```


# 生成CRUD文件
```shell
mongox gen ./data/schema --target ./data
```

## 示例结果

### 目录结构
```text
│  client.go
    │
    ├─schema
    │      example.go
    │
    └─user
            user.go
            user_aggregate.go
            user_create.go
            user_delete.go
            user_query.go
            user_update.go

```

### client.go

```go
// Code generated by mongox, DO NOT EDIT.
package simple

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-kenka/mongox/examples/simple/user"
	"github.com/go-kenka/mongox/model/aggregates/watch"
	"github.com/go-kenka/mongox/model/filters"
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
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Println("[MONGOX]", evt.Command)
		},
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri).SetMonitor(cmdMonitor))
	if err != nil {
		return nil, err
	}
	return NewClient(client), nil
}

// Watch .
func (c *Client) Watch(ctx context.Context, pipeline watch.WatchPipe, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return c.DB.Watch(ctx, pipeline.Pipe(), opts...)
}

// ListDatabaseNames .
func (c *Client) ListDatabaseNames(ctx context.Context, filter filters.Filter, opts ...*options.ListDatabasesOptions) ([]string, error) {
	return c.DB.ListDatabaseNames(ctx, filter.Document(), opts...)
}

// ListDatabases .
func (c *Client) ListDatabases(ctx context.Context, filter filters.Filter, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error) {
	return c.DB.ListDatabases(ctx, filter.Document(), opts...)
}

// UseSession .
func (c *Client) UseSession(ctx context.Context, opts *options.SessionOptions, fn func(sessionContext mongo.SessionContext) error) error {
	return c.DB.UseSessionWithOptions(ctx, opts, fn)
}

// WithTransaction .
func (c *Client) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	session, err := c.DB.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	return session.WithTransaction(ctx, fn, opts...)
}


```

### 使用示例
```go
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
```