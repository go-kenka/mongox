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
	YEAR        = newMongoTimeUnit("year", false)
	QUARTER     = newMongoTimeUnit("quarter", false)
	MONTH       = newMongoTimeUnit("month", false)
	WEEK        = newMongoTimeUnit("week", true)
	DAY         = newMongoTimeUnit("day", true)
	HOUR        = newMongoTimeUnit("hour", true)
	MINUTE      = newMongoTimeUnit("minute", true)
	SECOND      = newMongoTimeUnit("second", true)
	MILLISECOND = newMongoTimeUnit("millisecond", true)
)

func (u MongoTimeUnit) GetValue() string {
	return u.value
}
func (u MongoTimeUnit) GetFixed() bool {
	return u.fixed
}
