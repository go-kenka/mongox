package filters

import (
	"github.com/go-kenka/mongox/bsonx"
)

// type Filter bsonx.Bson

type MatchFilter Filter

// func And(filters ...Filter) MatchFilter {
// 	return filter.NewAddFilter(filters)
// }
// func Or(filters ...Filter) MatchFilter {
// 	return filter.NewOrNorFilter(filter.OR, filters)
// }
// func Nor(filters ...Filter) MatchFilter {
// 	return filter.NewOrNorFilter(filter.NOR, filters)
// }
//
// func Not(fieldName string, f Filter) MatchFilter {
// 	return filter.NewNotFilter(fieldName, f)
// }
//
// func Exists(fieldName string, value bool) MatchFilter {
// 	return filter.NewOperatorFilter("$exists", fieldName, bsonx.Boolean(value))
// }
// func Type(fieldName string, value bsonx.BsonType) MatchFilter {
// 	return filter.NewOperatorFilter("$type", fieldName, bsonx.Int32(value.Value()))
// }

func Empty() MatchFilter {
	return bsonx.BsonEmpty()
}
