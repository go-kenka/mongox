package gen

import (
	"github.com/go-kenka/mongox/types"
)

const (
	Invalid          = types.Invalid
	Double           = types.Double
	String           = types.String
	EmbeddedDocument = types.EmbeddedDocument
	Array            = types.Array
	Binary           = types.Binary
	Undefined        = types.Undefined
	ObjectID         = types.ObjectID
	Boolean          = types.Boolean
	DateTime         = types.DateTime
	Null             = types.Null
	Regex            = types.Regex
	DBPointer        = types.DBPointer
	JavaScript       = types.JavaScript
	Symbol           = types.Symbol
	CodeWithScope    = types.CodeWithScope
	Int32            = types.Int32
	Timestamp        = types.Timestamp
	Int64            = types.Int64
	Decimal128       = types.Decimal128
	MinKey           = types.MinKey
	MaxKey           = types.MaxKey
)

var (
	TypeGoNames = [...]string{
		Invalid:          "invalid",
		Double:           "float64",
		String:           "string",
		EmbeddedDocument: "struct",
		Array:            "[]struct",
		Binary:           "primitive.Binary",
		Undefined:        "primitive.Undefined",
		ObjectID:         "primitive.ObjectID",
		Boolean:          "bool",
		DateTime:         "primitive.DateTime",
		Null:             "primitive.Null",
		Regex:            "primitive.Regex",
		DBPointer:        "primitive.DBPointer",
		JavaScript:       "primitive.JavaScript",
		Symbol:           "primitive.Symbol",
		CodeWithScope:    "primitive.CodeWithScope",
		Int32:            "int32",
		Timestamp:        "primitive.Timestamp",
		Int64:            "int64",
		Decimal128:       "primitive.Decimal128",
		MinKey:           "primitive.MinKey",
		MaxKey:           "primitive.MaxKey",
	}

	TypeNames = [...]string{
		Invalid:          "Invalid",
		Double:           "Double",
		String:           "String",
		EmbeddedDocument: "EmbeddedDocument",
		Array:            "Array",
		Binary:           "Binary",
		Undefined:        "Undefined",
		ObjectID:         "ObjectID",
		Boolean:          "Boolean",
		DateTime:         "DateTime",
		Null:             "Null",
		Regex:            "Regex",
		DBPointer:        "DBPointer",
		JavaScript:       "JavaScript",
		Symbol:           "Symbol",
		CodeWithScope:    "CodeWithScope",
		Int32:            "Int32",
		Timestamp:        "Timestamp",
		Int64:            "Int64",
		Decimal128:       "Decimal128",
		MinKey:           "MinKey",
		MaxKey:           "MaxKey",
	}
)
