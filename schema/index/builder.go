package index

import (
	"time"

	"github.com/go-kenka/mongox/types"
)

// Index 索引
func Index(key string) *indexBuilder {
	return &indexBuilder{&Descriptor{
		Background: true,
		StorageKey: key,
	}}
}

type indexBuilder struct {
	desc *Descriptor
}

func (b *indexBuilder) Keys(keys ...types.MapEntry) *indexBuilder {
	b.desc.Keys = keys
	return b
}

func (b *indexBuilder) Unique() *indexBuilder {
	b.desc.Unique = true
	return b
}

func (b *indexBuilder) Background() *indexBuilder {
	b.desc.Background = true
	return b
}

func (b *indexBuilder) Sparse() *indexBuilder {
	b.desc.Sparse = true
	return b
}

func (b *indexBuilder) ExpireTime(d time.Duration) *indexBuilder {
	b.desc.Expire = d
	return b
}
func (b *indexBuilder) Descriptor() *Descriptor {
	return b.desc
}
