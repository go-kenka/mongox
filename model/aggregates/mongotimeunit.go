package aggregates

type MongoTimeUnit struct {
	value string
	fixed bool
}

func newMongoTimeUnit(value string, fixed bool) MongoTimeUnit {
	return MongoTimeUnit{
		value: value,
		fixed: fixed,
	}
}

var (
	Year        = newMongoTimeUnit("year", false)
	Quarter     = newMongoTimeUnit("quarter", false)
	Month       = newMongoTimeUnit("month", false)
	Week        = newMongoTimeUnit("week", true)
	Day         = newMongoTimeUnit("day", true)
	Hour        = newMongoTimeUnit("hour", true)
	Minute      = newMongoTimeUnit("minute", true)
	Second      = newMongoTimeUnit("second", true)
	Millisecond = newMongoTimeUnit("millisecond", true)
)

func (u MongoTimeUnit) GetValue() string {
	return u.value
}
func (u MongoTimeUnit) GetFixed() bool {
	return u.fixed
}
