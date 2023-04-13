package expression

import (
	"github.com/go-kenka/mongox/bsonx"
)

type Expression interface {
	Exp() bsonx.IBsonValue
}

// AnyExpression 整数表达式
type AnyExpression interface {
	Expression
	bsonx.IBsonValue
}

type defaultExpression interface {
	// 未解析的表达式 | 对象表达式 | 文字、路径、系统变量
	LiteralExpression | ObjExpression | Variable
}

// Stringer 字符串类型
type Stringer interface {
	// 默认表达式 | 字符串
	defaultExpression | bsonx.BsonString
}

// StrExpression 字符串表达式
type StrExpression interface {
	Expression
	bsonx.IBsonValue
	Stringer
}

// Integer 整数类型
type Integer interface {
	// 默认表达式 | 64位整数 | 32位整数
	defaultExpression | bsonx.BsonInt64 | bsonx.BsonInt32
}

// IntExpression 整数表达式
type IntExpression interface {
	Expression
	bsonx.IBsonValue
	Integer
}

// IntStrExpression 整数或者字符串表达式
type IntStrExpression interface {
	Expression
	bsonx.IBsonValue
	Integer | Stringer
}

// Number 数值类型
type Number interface {
	// 默认表达式 | 64位整数 | 32位整数 | 浮点数 | decimal数
	defaultExpression | bsonx.BsonInt64 | bsonx.BsonInt32 | bsonx.BsonDouble | bsonx.BsonDecimal128
}

// NumberExpression 数值表达式
type NumberExpression interface {
	Expression
	bsonx.IBsonValue
	Number
}

// Date 日期类型
type Date interface {
	// 默认表达式 | 日期 | 时间戳
	defaultExpression | bsonx.BsonDateTime | bsonx.BsonTimestamp | bsonx.BsonObjectId
}

// DateExpression 日期表达式
type DateExpression interface {
	Expression
	bsonx.IBsonValue
	Date
}

// DateNumberExpression 数字或日期表达式
type DateNumberExpression interface {
	Expression
	bsonx.IBsonValue
	Date | Number
}

// Array 日期类型
type Array interface {
	// 默认表达式 | 数组
	defaultExpression | bsonx.BsonArray
}

// ArrayExpression 数组表达式
type ArrayExpression interface {
	Expression
	bsonx.IBsonValue
	Array
}

// Boolean 布尔类型
type Boolean interface {
	// 默认表达式 | 布尔
	defaultExpression | bsonx.BsonBoolean
}

// BooleanExpression 数组表达式
type BooleanExpression interface {
	Expression
	bsonx.IBsonValue
	Boolean
}

// Document 文档对象类型
type Document interface {
	// 默认表达式 | 文档
	defaultExpression | bsonx.BsonDocument
}

// DocumentExpression 文档对象表达式
type DocumentExpression interface {
	Expression
	bsonx.IBsonValue
	Object
}

// Object 对象类型
type Object interface {
	// 默认表达式 | 文档| NULL | 数组对象 | undefined
	defaultExpression | bsonx.BsonDocument | bsonx.BsonNull | bsonx.BsonArray | bsonx.BsonUndefined
}

// ObjectExpression 对象表达式
type ObjectExpression interface {
	Expression
	bsonx.IBsonValue
	Object
}

type Binary interface {
	// 默认表达式 | 字节
	defaultExpression | bsonx.BsonString | bsonx.BsonBinary
}

// BinaryExpression byte数组表达式
type BinaryExpression interface {
	Expression
	bsonx.IBsonValue
	Binary
}
