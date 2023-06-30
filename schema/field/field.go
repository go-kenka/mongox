package field

import (
	"github.com/go-kenka/mongox/types"
)

// A Descriptor for field configuration.
type Descriptor struct {
	Tag       string          // struct tag.
	Name      string          // field name.
	FieldType types.MongoType // field type
	Optional  bool            // nullable field in database.
	EmbedData []*Descriptor   // 内嵌结构
	ArrayType types.MongoType // 数组结构
	Err       error
}
