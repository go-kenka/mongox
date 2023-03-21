package gen

import (
	"github.com/go-kenka/mongox/dsl"
)

type MongoType dsl.MongoType

const (
	TypeInvalid   = dsl.TypeInvalid
	TypeDouble    = dsl.TypeDouble
	TypeString    = dsl.TypeString
	TypeObject    = dsl.TypeObject
	TypeArray     = dsl.TypeArray
	TypeBinData   = dsl.TypeBinData
	TypeObjectID  = dsl.TypeObjectID
	TypeBoolean   = dsl.TypeBoolean
	TypeDate      = dsl.TypeDate
	TypeInteger   = dsl.TypeInteger
	TypeTimestamp = dsl.TypeTimestamp
	TypeDecimal   = dsl.TypeDecimal
)

var (
	TypeGoNames = [...]string{
		TypeInvalid:   "invalid",
		TypeDouble:    "float64",
		TypeString:    "string",
		TypeBinData:   "primitive.Binary",
		TypeObjectID:  "primitive.ObjectID",
		TypeBoolean:   "bool",
		TypeDate:      "primitive.DateTime",
		TypeInteger:   "int64",
		TypeTimestamp: "primitive.Timestamp",
		TypeDecimal:   "primitive.Decimal128",
	}

	TypeNames = [...]string{
		TypeInvalid:   "TypeInvalid",
		TypeDouble:    "TypeDouble",
		TypeString:    "TypeString",
		TypeObject:    "TypeObject",
		TypeArray:     "TypeArray",
		TypeBinData:   "TypeBinData",
		TypeObjectID:  "TypeObjectID",
		TypeBoolean:   "TypeBoolean",
		TypeDate:      "TypeDate",
		TypeInteger:   "TypeInteger",
		TypeTimestamp: "TypeTimestamp",
		TypeDecimal:   "TypeDecimal",
	}

	TypeNameMap = map[string]dsl.MongoType{
		"TypeInvalid":   TypeInvalid,
		"TypeDouble":    TypeDouble,
		"TypeString":    TypeString,
		"TypeObject":    TypeObject,
		"TypeArray":     TypeArray,
		"TypeBinData":   TypeBinData,
		"TypeObjectID":  TypeObjectID,
		"TypeBoolean":   TypeBoolean,
		"TypeDate":      TypeDate,
		"TypeInteger":   TypeInteger,
		"TypeTimestamp": TypeTimestamp,
		"TypeDecimal":   TypeDecimal,
	}
)

type Key struct {
	Key   string
	Value any // 1 升序 -1 降序 "2d" "2dsphere" "geoHaystack" "hashed" "text"
}

type Field struct {
	Name       string        // 字段名称
	Type       dsl.MongoType // 字段类型
	ObjectAttr []*Field      // 只在Object和Array类型时生效
	ArrayType  dsl.MongoType // Array类型时生效，如果是对象数组，则依赖ObjectAttr对象属性字段
	Comment    string        // 备注
}

type Index struct {
	Name       string // 索引名称
	Unique     bool   // 是否唯一
	Background bool   // 是否后台运行
	Sparse     bool   // 是否是稀疏索引（null不包含）
	Expire     string // 是否过期
	Keys       []*Key // 索引组合
}

type Collection struct {
	Database string   // 所属数据库
	Name     string   // 集合名称
	Desc     string   // 备注
	Fields   []*Field // 表字段集合
	Indexes  []*Index // 表字段集合
}
