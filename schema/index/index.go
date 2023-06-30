package index

import (
	"time"

	"github.com/go-kenka/mongox/types"
)

// A Descriptor for index configuration.
type Descriptor struct {
	StorageKey string           // 索引名称
	Unique     bool             // 是否唯一
	Background bool             // 是否后台运行
	Sparse     bool             // 是否是稀疏索引（null不包含）
	Expire     time.Duration    // 是否过期
	Keys       []types.MapEntry // 索引组合
	Err        error
}
