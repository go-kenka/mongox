package dsl

// A MongoType represents a field type.
type MongoType uint8

// List of field types.
const (
	TypeInvalid MongoType = iota
	TypeDouble
	TypeString
	TypeObject
	TypeArray
	TypeBinData
	TypeObjectID
	TypeBoolean
	TypeDate
	TypeInteger
	TypeTimestamp
	TypeDecimal
	endTypes
)
