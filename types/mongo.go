package types

// A MongoType represents a field type.
type MongoType uint8

// List of field types.
const (
	Invalid          MongoType = 0x00
	Double           MongoType = 0x01
	String           MongoType = 0x02
	EmbeddedDocument MongoType = 0x03
	Array            MongoType = 0x04
	Binary           MongoType = 0x05
	Undefined        MongoType = 0x06
	ObjectID         MongoType = 0x07
	Boolean          MongoType = 0x08
	DateTime         MongoType = 0x09
	Null             MongoType = 0x0A
	Regex            MongoType = 0x0B
	DBPointer        MongoType = 0x0C
	JavaScript       MongoType = 0x0D
	Symbol           MongoType = 0x0E
	CodeWithScope    MongoType = 0x0F
	Int32            MongoType = 0x10
	Timestamp        MongoType = 0x11
	Int64            MongoType = 0x12
	Decimal128       MongoType = 0x13
	MinKey           MongoType = 0xFF
	MaxKey           MongoType = 0x7F

	BinaryGeneric     byte = 0x00
	BinaryFunction    byte = 0x01
	BinaryBinaryOld   byte = 0x02
	BinaryUUIDOld     byte = 0x03
	BinaryUUID        byte = 0x04
	BinaryMD5         byte = 0x05
	BinaryEncrypted   byte = 0x06
	BinaryColumn      byte = 0x07
	BinaryUserDefined byte = 0x80
)
