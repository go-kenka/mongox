// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package load

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-kenka/mongox"
	"github.com/go-kenka/mongox/schema/field"
	"github.com/go-kenka/mongox/schema/index"
	"github.com/go-kenka/mongox/types"
	"github.com/gobeam/stringy"
)

// Schema represents an ent.Schema that was loaded from a complied user package.
type Schema struct {
	Name    string   `json:"name,omitempty"`
	Fields  []*Field `json:"fields,omitempty"`
	Indexes []*Index `json:"indexes,omitempty"`
}

// Field represents an ent.Field that was loaded from a complied user package.
type Field struct {
	Tag       string          `json:"tag,omitempty"`
	Name      string          `json:"name,omitempty"`
	FieldType types.MongoType `json:"type,omitempty"`
	Optional  bool            `json:"optional,omitempty"`
	EmbedData []*Field        `json:"embed_data,omitempty"`
	ArrayType types.MongoType `json:"array_type,omitempty"`
}

// Index represents an ent.Index that was loaded from a complied user package.
type Index struct {
	StorageKey string           `json:"storage_key,omitempty"`
	Unique     bool             `json:"unique,omitempty"`
	Background bool             `json:"background,omitempty"`
	Sparse     bool             `json:"sparse,omitempty"`
	Expire     time.Duration    `json:"expire,omitempty"`
	Keys       []types.MapEntry `json:"keys,omitempty"`
}

type SubObject struct {
	Name   string
	Fields []*Field
}

// NewField creates a loaded field from field descriptor.
func NewField(fd *field.Descriptor) (*Field, error) {
	if fd.Err != nil {
		return nil, fmt.Errorf("field %q: %v", fd.Name, fd.Err)
	}

	embedData := make([]*Field, 0)

	for _, data := range fd.EmbedData {
		fs, err := NewField(data)
		if err != nil {
			return nil, err
		}
		embedData = append(embedData, fs)
	}

	sf := &Field{
		Tag:       fd.Tag,
		Name:      fd.Name,
		FieldType: fd.FieldType,
		Optional:  fd.Optional,
		EmbedData: embedData,
		ArrayType: fd.ArrayType,
	}

	return sf, nil
}

// NewIndex creates an loaded index from index descriptor.
func NewIndex(idx *index.Descriptor) *Index {
	ni := &Index{
		StorageKey: idx.StorageKey,
		Unique:     idx.Unique,
		Background: idx.Background,
		Sparse:     idx.Sparse,
		Keys:       idx.Keys,
	}
	return ni
}

// MarshalSchema encodes the mongox.Schema interface into a JSON
// that can be decoded into the Schema objects declared above.
func MarshalSchema(schema mongox.Interface) (b []byte, err error) {
	s := &Schema{
		Name: indirect(reflect.TypeOf(schema)).Name(),
	}

	if err := s.loadFields(schema); err != nil {
		return nil, fmt.Errorf("schema %q: %w", s.Name, err)
	}

	indexes, err := safeIndexes(schema)
	if err != nil {
		return nil, fmt.Errorf("schema %q: %w", s.Name, err)
	}
	for _, idx := range indexes {
		s.Indexes = append(s.Indexes, NewIndex(idx.Descriptor()))
	}

	return json.Marshal(s)
}

// UnmarshalSchema decodes the given buffer to a loaded schema.
func UnmarshalSchema(buf []byte) (*Schema, error) {
	s := &Schema{}
	if err := json.Unmarshal(buf, s); err != nil {
		return nil, err
	}
	return s, nil
}

// loadFields loads field to schema from ent.Interface.
func (s *Schema) loadFields(schema mongox.Interface) error {
	fields, err := safeFields(schema)
	if err != nil {
		return err
	}
	for _, f := range fields {
		sf, err := NewField(f.Descriptor())
		if err != nil {
			return err
		}
		s.Fields = append(s.Fields, sf)
	}
	return nil
}

// safeFields wraps the schema.Fields and mixin.Fields method with recover to ensure no panics in marshaling.
func safeFields(fd interface{ Fields() []mongox.Field }) (fields []mongox.Field, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("%T.Fields panics: %v", fd, v)
			fields = nil
		}
	}()
	return fd.Fields(), nil
}

// safeIndexes wraps the schema.Indexes method with recover to ensure no panics in marshaling.
func safeIndexes(schema interface{ Indexes() []mongox.Index }) (indexes []mongox.Index, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("schema.Indexes panics: %v", v)
			indexes = nil
		}
	}()
	return schema.Indexes(), nil
}

func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// PathList 获取路径
func (s *Schema) PathList() []string {
	path := getPath("", s.Fields)
	return path
}

// FieldNames 获取路径
func (s *Schema) FieldNames() []string {
	names := getFieldName("Field", s.Fields)
	return names
}

func getPath(prefix string, fs []*Field) []string {
	var path []string

	if len(fs) == 0 {
		return nil
	}

	for _, f := range fs {
		var current string
		if len(prefix) == 0 {
			current = strings.ToLower(f.Name)
		} else {
			current = strings.Join([]string{prefix, strings.ToLower(f.Name)}, ".")
		}
		path = append(path, current)
		children := getPath(current, f.EmbedData)
		path = append(path, children...)
	}

	return path
}

func getFieldName(prefix string, fs []*Field) []string {
	var names []string

	if len(fs) == 0 {
		return nil
	}

	for _, f := range fs {
		var current string
		if len(prefix) == 0 {
			current = stringy.New(f.Name).CamelCase()
		} else {
			current = prefix + stringy.New(f.Name).CamelCase()
		}
		names = append(names, current)
		children := getFieldName(current, f.EmbedData)
		names = append(names, children...)
	}

	return names
}

// SubObject 获取对象数组
func (s *Schema) SubObject() []*SubObject {

	root := SubObject{
		Name:   s.Name,
		Fields: []*Field{},
	}

	objects := getSubObject(&root, s.Fields)
	return objects
}

func getSubObject(parent *SubObject, fs []*Field) []*SubObject {
	var names []*SubObject

	if len(fs) == 0 {
		return nil
	}
	// 当前字段添加到父节点
	parent.Fields = append(parent.Fields, fs...)
	// 添加父节点
	names = append(names, parent)

	for _, f := range fs {
		current := parent.Name + stringy.New(f.Name).CamelCase()
		// 子对象的场合
		if f.FieldType == types.EmbeddedDocument {
			child := SubObject{
				Name:   current,
				Fields: []*Field{},
			}
			children := getSubObject(&child, f.EmbedData)
			names = append(names, children...)
		}
		// 数组对象的场合
		if f.FieldType == types.Array && f.ArrayType == types.Invalid {
			child := SubObject{
				Name:   current,
				Fields: []*Field{},
			}
			children := getSubObject(&child, f.EmbedData)
			names = append(names, children...)
		}
	}

	return names
}
