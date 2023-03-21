package bsonx

import (
	"log"
	"reflect"
)

type Document map[string]any

func NewDoc() Document {
	return Document{}
}

func NewDocument(key string, value any) Document {
	return Document{
		key: value,
	}
}

func (d Document) ToBsonDocument() BsonDocument {
	bd := NewEmptyDoc()
	for k, v := range d {
		bd.Append(k, toBsonValue(v))
	}
	return bd
}

func (d Document) Document() Document {
	return d
}

func (d Document) Append(key string, value any) Document {
	d[key] = value
	return d
}
func (d Document) Remove(key string) Document {
	delete(d, key)
	return d
}
func (d Document) IsEmpty() bool {
	return len(d) == 0
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
		return BsonNull{}
	// 基础类型
	case reflect.String:
		return NewBsonString(v.String())
	case reflect.Bool:
		return NewBsonBoolean(v.Bool())
	case reflect.Float64, reflect.Float32:
		return NewBsonDouble(v.Float())
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		return NewBsonInt64(v.Int())
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		return NewBsonInt64(v.Int())
	case reflect.Array, reflect.Slice:
		return toBsonArray(val)
	case reflect.Complex64, reflect.Complex128, reflect.Uintptr, reflect.Pointer, reflect.UnsafePointer, reflect.Chan, reflect.Func, reflect.Struct, reflect.Interface:
		log.Fatal("cant support this type")
		return nil
	case reflect.Map:
		return toBsonDocument(v.MapRange())
	default:
		return BsonNull{}
	}
}

func toBsonArray(a any) BsonArray {
	aList := NewBsonArray()
	for _, d := range a.([]any) {
		aList.Append(toBsonValue(d))
	}
	return aList
}

func toBsonDocument(a *reflect.MapIter) BsonDocument {
	aList := NewEmptyDoc()
	for a.Next() {
		aList.Append(a.Key().String(), toBsonValue(a.Value().Interface()))
	}

	return aList
}
