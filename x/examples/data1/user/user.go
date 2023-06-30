// Code generated by mongox, DO NOT EDIT.
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
    CollectionName    = "user"
        ColumnFieldId = "_id"
        ColumnFieldString = "string"
        ColumnFieldBool = "bool"
        ColumnFieldBinary = "binary"
        ColumnFieldDouble = "double"
        ColumnFieldDataTime = "data_time"
        ColumnFieldCreateAt = "create_at"
        ColumnFieldPointer = "pointer"
        ColumnFieldDecimal = "decimal"
        ColumnFieldInt32 = "int32"
        ColumnFieldInt64 = "int64"
        ColumnFieldJs = "js"
        ColumnFieldJsScope = "js_scope"
        ColumnFieldRegular = "regular"
        ColumnFieldInt32s = "int32s"
        ColumnFieldInt64s = "int64s"
        ColumnFieldStrings = "strings"
        ColumnFieldFloats = "floats"
        ColumnFieldArrayobjectSimple = "arrayobject_simple"
        ColumnFieldArrayobjectSimpleId = "arrayobject_simple._id"
        ColumnFieldArrayobjectSimpleString = "arrayobject_simple.string"
        ColumnFieldArrayobjectSimpleBool = "arrayobject_simple.bool"
        ColumnFieldArrayobjectSimpleBinary = "arrayobject_simple.binary"
        ColumnFieldArrayobjectSimpleDouble = "arrayobject_simple.double"
        ColumnFieldArrayobjectSimpleDataTime = "arrayobject_simple.data_time"
        ColumnFieldArrayobjectSimpleCreateAt = "arrayobject_simple.create_at"
        ColumnFieldArrayobjectSimplePointer = "arrayobject_simple.pointer"
        ColumnFieldArrayobjectSimpleDecimal = "arrayobject_simple.decimal"
        ColumnFieldArrayobjectSimpleInt32 = "arrayobject_simple.int32"
        ColumnFieldArrayobjectSimpleInt64 = "arrayobject_simple.int64"
        ColumnFieldArrayobjectSimpleJs = "arrayobject_simple.js"
        ColumnFieldArrayobjectSimpleJsScope = "arrayobject_simple.js_scope"
        ColumnFieldArrayobjectSimpleRegular = "arrayobject_simple.regular"
        ColumnFieldArrayobjectSimpleInt32s = "arrayobject_simple.int32s"
        ColumnFieldArrayobjectSimpleInt64s = "arrayobject_simple.int64s"
        ColumnFieldArrayobjectSimpleStrings = "arrayobject_simple.strings"
        ColumnFieldArrayobjectSimpleFloats = "arrayobject_simple.floats"
        ColumnFieldObjectSimple = "object_simple"
        ColumnFieldObjectSimpleId = "object_simple._id"
        ColumnFieldObjectSimpleString = "object_simple.string"
        ColumnFieldObjectSimpleBool = "object_simple.bool"
        ColumnFieldObjectSimpleBinary = "object_simple.binary"
        ColumnFieldObjectSimpleDouble = "object_simple.double"
        ColumnFieldObjectSimpleDataTime = "object_simple.data_time"
        ColumnFieldObjectSimpleCreateAt = "object_simple.create_at"
        ColumnFieldObjectSimplePointer = "object_simple.pointer"
        ColumnFieldObjectSimpleDecimal = "object_simple.decimal"
        ColumnFieldObjectSimpleInt32 = "object_simple.int32"
        ColumnFieldObjectSimpleInt64 = "object_simple.int64"
        ColumnFieldObjectSimpleJs = "object_simple.js"
        ColumnFieldObjectSimpleJsScope = "object_simple.js_scope"
        ColumnFieldObjectSimpleRegular = "object_simple.regular"
        ColumnFieldObjectSimpleInt32s = "object_simple.int32s"
        ColumnFieldObjectSimpleInt64s = "object_simple.int64s"
        ColumnFieldObjectSimpleStrings = "object_simple.strings"
        ColumnFieldObjectSimpleFloats = "object_simple.floats")
type UserData struct {
        Id primitive.ObjectID `bson:"_id" json:"aaaa"`
        String string `bson:"string" json:"s1"`
        Bool bool `bson:"bool" json:"s1"`
        Binary primitive.Binary `bson:"binary" json:"b1"`
        Double float64 `bson:"double" json:"d1"`
        DataTime primitive.DateTime `bson:"data_time" json:"aaaa"`
        CreateAt primitive.Timestamp `bson:"create_at" json:"aaaa"`
        Pointer primitive.DBPointer `bson:"pointer" json:"aaaa"`
        Decimal primitive.Decimal128 `bson:"decimal" json:"aaaa"`
        Int32 int32 `bson:"int32" json:"aaaa"`
        Int64 int64 `bson:"int64" json:"aaaa"`
        Js primitive.JavaScript `bson:"js" json:"aaaa"`
        JsScope primitive.CodeWithScope `bson:"js_scope" json:"aaaa"`
        Regular primitive.Regex `bson:"regular" json:"aaaa"`
            Int32s []int32 `bson:"int32s" json:"aaaa"`
            Int64s []int64 `bson:"int64s" json:"aaaa"`
            Strings []string `bson:"strings" json:"aaaa"`
            Floats []float64 `bson:"floats" json:"aaaa"`
            ArrayobjectSimpleList []*UserArrayobjectSimpleData `bson:"arrayobject_simple" json:"aaaa"`
        ObjectSimple *UserObjectSimpleData `bson:"object_simple" json:"aaaa"`
}
type UserArrayobjectSimpleData struct {
        Id primitive.ObjectID `bson:"_id" json:"aaaa"`
        String string `bson:"string" json:"s1"`
        Bool bool `bson:"bool" json:"s1"`
        Binary primitive.Binary `bson:"binary" json:"b1"`
        Double float64 `bson:"double" json:"d1"`
        DataTime primitive.DateTime `bson:"data_time" json:"aaaa"`
        CreateAt primitive.Timestamp `bson:"create_at" json:"aaaa"`
        Pointer primitive.DBPointer `bson:"pointer" json:"aaaa"`
        Decimal primitive.Decimal128 `bson:"decimal" json:"aaaa"`
        Int32 int32 `bson:"int32" json:"aaaa"`
        Int64 int64 `bson:"int64" json:"aaaa"`
        Js primitive.JavaScript `bson:"js" json:"aaaa"`
        JsScope primitive.CodeWithScope `bson:"js_scope" json:"aaaa"`
        Regular primitive.Regex `bson:"regular" json:"aaaa"`
            Int32s []int32 `bson:"int32s" json:"aaaa"`
            Int64s []int64 `bson:"int64s" json:"aaaa"`
            Strings []string `bson:"strings" json:"aaaa"`
            Floats []float64 `bson:"floats" json:"aaaa"`
}
type UserObjectSimpleData struct {
        Id primitive.ObjectID `bson:"_id" json:"aaaa"`
        String string `bson:"string" json:"s1"`
        Bool bool `bson:"bool" json:"s1"`
        Binary primitive.Binary `bson:"binary" json:"b1"`
        Double float64 `bson:"double" json:"d1"`
        DataTime primitive.DateTime `bson:"data_time" json:"aaaa"`
        CreateAt primitive.Timestamp `bson:"create_at" json:"aaaa"`
        Pointer primitive.DBPointer `bson:"pointer" json:"aaaa"`
        Decimal primitive.Decimal128 `bson:"decimal" json:"aaaa"`
        Int32 int32 `bson:"int32" json:"aaaa"`
        Int64 int64 `bson:"int64" json:"aaaa"`
        Js primitive.JavaScript `bson:"js" json:"aaaa"`
        JsScope primitive.CodeWithScope `bson:"js_scope" json:"aaaa"`
        Regular primitive.Regex `bson:"regular" json:"aaaa"`
            Int32s []int32 `bson:"int32s" json:"aaaa"`
            Int64s []int64 `bson:"int64s" json:"aaaa"`
            Strings []string `bson:"strings" json:"aaaa"`
            Floats []float64 `bson:"floats" json:"aaaa"`
}

func (d UserData) Document() any {
    return d
}
func (d UserData) Update() {}

type UserClient struct {
    c  *mongo.Client
    db string
}

func NewUserClient(db *mongo.Client) *UserClient {
    return &UserClient{
        c: db,
    }
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
