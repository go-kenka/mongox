package filters

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type bitwiseFilter struct {
	filter bsonx.Bson
}

func (f bitwiseFilter) Value() bsonx.IBsonValue {
	return f.filter.Pro()
}

func (f bitwiseFilter) Document() bson.D {
	return f.filter.Document()
}

// BitsAllClear $bitsAllClear matches documents where all of the bit positions
// given by the query are clear (i.e. 0) in field. { <field>: { $bitsAllClear:
// <numeric bitmask> } } { <field>: { $bitsAllClear: < BinData bitmask> } } {
// <field>: { $bitsAllClear: [ <position1>, <position2>, ... ] } } The field
// value must be either numeric or a BinData instance. Otherwise, $bitsAllClear
// will not match the current document.
func BitsAllClear(fieldName string, bitmask int64) bitwiseFilter {
	return bitwiseFilter{filter: newOperatorFilter("$bitsAllClear", fieldName, bsonx.Int64(bitmask))}
}

// BitsAllSet $bitsAllSet matches documents where all of the bit positions given by the
// query are set (i.e. 1) in field. { <field>: { $bitsAllSet: <numeric bitmask>
// } } { <field>: { $bitsAllSet: < BinData bitmask> } } { <field>: {
// $bitsAllSet: [ <position1>, <position2>, ... ] } } The field value must be
// either numeric or a BinData instance. Otherwise, $bitsAllSet will not match
// the current document.
func BitsAllSet(fieldName string, bitmask int64) bitwiseFilter {
	return bitwiseFilter{filter: newOperatorFilter("$bitsAllSet", fieldName, bsonx.Int64(bitmask))}
}

// BitsAnyClear $bitsAnyClear matches documents where any of the bit positions given by the
// query are clear (i.e. 0) in field. { <field>: { $bitsAnyClear: <numeric
// bitmask> } } { <field>: { $bitsAnyClear: < BinData bitmask> } } { <field>: {
// $bitsAnyClear: [ <position1>, <position2>, ... ] } } The field value must be
// either numeric or a BinData instance. Otherwise, $bitsAnyClear will not match
// the current document.
func BitsAnyClear(fieldName string, bitmask int64) bitwiseFilter {
	return bitwiseFilter{filter: newOperatorFilter("$bitsAnyClear", fieldName, bsonx.Int64(bitmask))}
}

// BitsAnySet $bitsAnySet matches documents where any of the bit positions given
// by the query are set (i.e. 1) in field. { <field>: { $bitsAnySet: <numeric
// bitmask> } } { <field>: { $bitsAnySet: < BinData bitmask> } } { <field>: {
// $bitsAnySet: [ <position1>, <position2>, ... ] } } The field value must be
// either numeric or a BinData instance. Otherwise, $bitsAnySet will not match
// the current document.
func BitsAnySet(fieldName string, bitmask int64) bitwiseFilter {
	return bitwiseFilter{filter: newOperatorFilter("$bitsAnySet", fieldName, bsonx.Int64(bitmask))}
}
