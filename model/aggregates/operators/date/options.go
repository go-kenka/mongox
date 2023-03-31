package date

import "github.com/go-kenka/mongox/bsonx"

type DateOptions struct {
	timezone string
}

type DateDiffOptions struct {
	timezone    string
	startOfWeek StartOfWeek
}

type DateFromPartsOptions struct {
	month        int32
	isoWeek      int32
	day          int32
	isoDayOfWeek int32
	hour         int32
	minute       int32
	second       int32
	millisecond  int32
	timezone     string
}

type DateFromStringOptions struct {
	format   string
	timezone string
	onError  bsonx.Bson
	onNull   bsonx.Bson
}

type DateToPartsOptions struct {
	timezone string
	iso8601  bool
}

type DateTruncOptions struct {
	timezone    string
	binSize     int64
	startOfWeek StartOfWeek
}
