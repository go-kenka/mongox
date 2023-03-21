package densify

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/aggregates"
	"time"
)

type DensifyRange interface {
	bsonx.Bson
}

func FullRangeWithStep(step int) NumberDensifyRange {
	return NewDensifyConstructibleBson(bsonx.NewDocument("bounds", "full").
		Append("step", step), nil)
}

func PartitionRangeWithStep(step int) NumberDensifyRange {
	return NewDensifyConstructibleBson(bsonx.NewDocument("bounds", "partition").
		Append("step", step), nil)
}

func RangeWithStep(l, u, step int) NumberDensifyRange {
	return NewDensifyConstructibleBson(bsonx.NewDocument("bounds", []int{l, u}).
		Append("step", step), nil)
}

func DateFullRangeWithStep(step int64, unit aggregates.MongoTimeUnit) DateDensifyRange {
	return NewDensifyConstructibleBson(bsonx.NewBsonDocument("bounds", bsonx.NewBsonString("partition")).
		Append("step", bsonx.NewBsonInt64(step)).
		Append("unit", bsonx.NewBsonString(unit.GetValue())), nil)
}

func DateRangeWithStep(l, u time.Duration, step int64, unit aggregates.MongoTimeUnit) NumberDensifyRange {
	return NewDensifyConstructibleBson(bsonx.NewDocument("bounds", []any{l, u}).
		Append("step", step).
		Append("unit", unit.GetValue()), nil)
}
