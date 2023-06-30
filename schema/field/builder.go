package field

import (
	"github.com/go-kenka/mongox/types"
)

// String 字符串
func String(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.String}}
}

// Bool 布尔
func Bool(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Boolean}}
}

// Time  日期
func Time(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.DateTime}}
}

// Float Float
func Float(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Double}}
}

// Int32 Int32
func Int32(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Int32}}
}

// Int64 Int64
func Int64(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Int64}}
}

// JavaScript JavaScript
func JavaScript(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.JavaScript}}
}

// JavaScriptScope JavaScriptScope
func JavaScriptScope(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.CodeWithScope}}
}

// ObjectId ObjectId
func ObjectId(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.ObjectID}}
}

// Symbol 标志
func Symbol(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Symbol}}
}

// Timestamp 时间戳
func Timestamp(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Timestamp}}
}

// DbPointer 坐标
func DbPointer(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.DBPointer}}
}

// Decimal128 Decimal128
func Decimal128(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Decimal128}}
}

// Bytes 二进制
func Bytes(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Binary}}
}

// Regular 正则表达式
func Regular(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Regex}}
}

// Int32s int32数组
func Int32s(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Array, ArrayType: types.Int32}}
}

// Int64s int64数组
func Int64s(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Array, ArrayType: types.Int64}}
}

// Strings 字符串数组
func Strings(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Array, ArrayType: types.String}}
}

// Floats 浮点数组
func Floats(name string) *builder {
	return &builder{&Descriptor{Name: name, FieldType: types.Array, ArrayType: types.Double}}
}

// Arrays 对象数组
func Arrays(name string) *arrayBuilder {
	return &arrayBuilder{&Descriptor{Name: name, FieldType: types.Array}}
}

// Object 内嵌对象
func Object(name string) *embedBuilder {
	return &embedBuilder{&Descriptor{Name: name, FieldType: types.EmbeddedDocument}}
}

type builder struct {
	desc *Descriptor
}

func (b *builder) Tag(tag string) *builder {
	b.desc.Tag = tag
	return b
}

func (b *builder) Optional() *builder {
	b.desc.Optional = true
	return b
}
func (b *builder) Descriptor() *Descriptor {
	return b.desc
}

type embedBuilder struct {
	desc *Descriptor
}

func (b *embedBuilder) Tag(tag string) *embedBuilder {
	b.desc.Tag = tag
	return b
}

func (b *embedBuilder) Optional() *embedBuilder {
	b.desc.Optional = true
	return b
}

func (b *embedBuilder) Attributes(f ...*Descriptor) *embedBuilder {
	b.desc.EmbedData = append(b.desc.EmbedData, f...)
	return b
}

func (b *embedBuilder) Descriptor() *Descriptor {
	return b.desc
}

type arrayBuilder struct {
	desc *Descriptor
}

func (b *arrayBuilder) Tag(tag string) *arrayBuilder {
	b.desc.Tag = tag
	return b
}

func (b *arrayBuilder) Optional() *arrayBuilder {
	b.desc.Optional = true
	return b
}

func (b *arrayBuilder) ArrayType(mongoType types.MongoType) *arrayBuilder {
	b.desc.ArrayType = mongoType
	return b
}

func (b *arrayBuilder) Attributes(f ...*Descriptor) *arrayBuilder {
	b.desc.EmbedData = append(b.desc.EmbedData, f...)
	return b
}

func (b *arrayBuilder) Descriptor() *Descriptor {
	return b.desc
}
