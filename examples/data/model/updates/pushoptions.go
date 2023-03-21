package updates

import (
	"fmt"
	"github.com/go-kenka/mongox/examples/data/bsonx"
)

type PushOptions struct {
	position     *int32
	slice        *int32
	sort         *int32
	sortDocument bsonx.Bson
}

func NewPushOptions() PushOptions {
	return PushOptions{}
}

func (b PushOptions) Position(position int32) PushOptions {
	b.position = &position
	return b
}

func (b PushOptions) HasPosition() bool {
	return b.position != nil
}

func (b PushOptions) GetPosition() int32 {
	return *b.position
}

func (b PushOptions) Slice(slice int32) PushOptions {
	b.slice = &slice
	return b
}

func (b PushOptions) HasSlice() bool {
	return b.slice != nil
}

func (b PushOptions) GetSlice() int32 {
	return *b.slice
}

func (b PushOptions) Sort(sort int32) PushOptions {
	b.sort = &sort
	return b
}

func (b PushOptions) HasSort() bool {
	return b.sort != nil
}

func (b PushOptions) GetSort() int32 {
	return *b.sort
}

func (b PushOptions) SortDocument(sortDocument bsonx.Bson) PushOptions {
	b.sortDocument = sortDocument
	return b
}

func (b PushOptions) GetSortDocument() bsonx.Bson {
	return b.sortDocument
}

func (b PushOptions) ToString() string {
	return fmt.Sprintf("%+v", b)
}
