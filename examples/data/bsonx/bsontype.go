package bsonx

type BsonType uint8

const (
	END_OF_DOCUMENT       BsonType = 0x00
	DOUBLE                BsonType = 0x01
	STRING                BsonType = 0x02
	DOCUMENT              BsonType = 0x03
	ARRAY                 BsonType = 0x04
	BINARY                BsonType = 0x05
	UNDEFINED             BsonType = 0x06
	OBJECT_ID             BsonType = 0x07
	BOOLEAN               BsonType = 0x08
	DATE_TIME             BsonType = 0x09
	NULL                  BsonType = 0x0a
	REGULAR_EXPRESSION    BsonType = 0x0b
	DB_POINTER            BsonType = 0x0c
	JAVASCRIPT            BsonType = 0x0d
	SYMBOL                BsonType = 0x0e
	JAVASCRIPT_WITH_SCOPE BsonType = 0x0f
	INT32                 BsonType = 0x10
	TIMESTAMP             BsonType = 0x11
	INT64                 BsonType = 0x12
	DECIMAL128            BsonType = 0x13
	MIN_KEY               BsonType = 0xff
	MAX_KEY               BsonType = 0x7f
)

func (t BsonType) Value() int32 {
	return int32(t)
}
