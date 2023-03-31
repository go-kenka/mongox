package densify

import (
	"time"

	bsonx2 "github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/aggregates"
)

type DensifyRange interface {
	bsonx2.Bson
}

func FullRangeWithStep(step int) NumberDensifyRange {
	return NewDensifyConstructibleBson(bsonx2.Doc("bounds", "full").
		Append("step", step), nil)
}

func PartitionRangeWithStep(step int) NumberDensifyRange {
	return NewDensifyConstructibleBson(bsonx2.Doc("bounds", "partition").
		Append("step", step), nil)
}

func RangeWithStep(l, u, step int) NumberDensifyRange {
	return NewDensifyConstructibleBson(bsonx2.Doc("bounds", []int{l, u}).
		Append("step", step), nil)
}

func DateFullRangeWithStep(step int64, unit aggregates.MongoTimeUnit) DateDensifyRange {
	return NewDensifyConstructibleBson(bsonx2.BsonDoc("bounds", bsonx2.String("partition")).
		Append("step", bsonx2.Int64(step)).
		Append("unit", bsonx2.String(unit.GetValue())), nil)
}

func DateRangeWithStep(l, u time.Duration, step int64, unit aggregates.MongoTimeUnit) NumberDensifyRange {
	return NewDensifyConstructibleBson(bsonx2.Doc("bounds", []any{l, u}).
		Append("step", step).
		Append("unit", unit.GetValue()), nil)
}
