package bsonx

import (
	"log"
	"reflect"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Document struct {
	data bson.D
}

func Empty() *Document {
	return &Document{
		data: bson.D{},
	}
}

func Doc(key string, value any) *Document {
	return &Document{
		data: bson.D{{key, value}},
	}
}

func (d *Document) ToBsonDocument() *BsonDocument {
	bd := BsonEmpty()
	for _, v := range d.data {
		bd.Append(v.Key, toBsonValue(v.Value))
	}
	return bd
}

func (d *Document) Document() bson.D {
	return d.data
}

func (d *Document) Append(key string, value any) *Document {
	d.data = append(d.data, bson.E{Key: key, Value: value})
	return d
}
func (d *Document) Remove(key string) *Document {
	lo.DropWhile(d.data, func(item primitive.E) bool {
		return item.Key == key
	})
	return d
}
func (d *Document) IsEmpty() bool {
	return d != nil && len(d.data) == 0
}

func toBsonValue(val any) IBsonValue {
	switch v := val.(type) {
	case IBsonValue:
		return v
	default:
		return doReflect(v)
	}
}

func doReflect(val any) IBsonValue {

	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.Invalid:
		return &BsonNull{}
	// 基础类型
	case reflect.String:
		return String(v.String())
	case reflect.Bool:
		return Boolean(v.Bool())
	case reflect.Float64, reflect.Float32:
		return Double(v.Float())
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		return Int64(v.Int())
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		return Int64(v.Int())
	case reflect.Array, reflect.Slice:
		return toBsonArray(val)
	case reflect.Complex64, reflect.Complex128, reflect.Uintptr, reflect.Pointer, reflect.UnsafePointer, reflect.Chan, reflect.Func, reflect.Struct, reflect.Interface:
		log.Fatal("cant support this type")
		return nil
	case reflect.Map:
		return toBsonDocument(v.MapRange())
	default:
		return &BsonNull{}
	}
}

func toBsonArray(a any) *BsonArray {
	aList := Array()
	for _, d := range a.([]any) {
		aList.Append(toBsonValue(d))
	}
	return aList
}

func toBsonDocument(a *reflect.MapIter) *BsonDocument {
	aList := BsonEmpty()
	for a.Next() {
		aList.Append(a.Key().String(), toBsonValue(a.Value().Interface()))
	}

	return aList
}
