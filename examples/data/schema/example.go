package schema

import (
	"time"

	"github.com/go-kenka/mongox"
	"github.com/go-kenka/mongox/schema/field"
	"github.com/go-kenka/mongox/schema/index"
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

func (User) Indexes() []mongox.Index {
	return []mongox.Index{
		index.Index("string_-1").Keys(types.MapEntry{
			Key:   "string",
			Value: types.TypeDesc,
		}).Background().Unique().Sparse().ExpireTime(time.Second),
	}
}
