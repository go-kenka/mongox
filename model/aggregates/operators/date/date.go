// Package date returns date objects or components of a date object
package date

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/internal/filter"
	"github.com/go-kenka/mongox/internal/options"
	"go.mongodb.org/mongo-driver/bson"
)

type dateOperator struct {
	doc bsonx.Bson
}

func (o dateOperator) Exp() bsonx.IBsonValue {
	return o.doc.BsonDocument()
}

// DateAdd NewDefaultStage in version 5.0.
// Increments a Date object by a specified number of time units.
// The $dateAdd expression has the following syntax:
//
//	{
//	  $dateAdd: {
//	     startDate: <Expression>,
//	     unit: <Expression>,
//	     amount: <Expression>,
//	     timezone: <tzExpression>
//	  }
//	}
//
// Returns a Date. The startDate can be any expression that resolves to type Date, Timestamp or ObjectId. No matter which data type is used as input, the value returned will be a Date object.
func DateAdd[T expression.DateExpression, N expression.NumberExpression](startDate T, unit options.MongoTimeUnit, amount N, options DateOptions) dateOperator {
	doc := bsonx.BsonDoc("startDate", startDate)
	doc.Append("unit", bsonx.String(unit.GetValue()))
	doc.Append("amount", amount)
	if options.timezone != "" {
		doc.Append("timezone", bsonx.String(options.timezone))
	}
	return dateOperator{doc: filter.NewSimpleFilter("$dateAdd", doc)}
}

// DateDiff NewDefaultStage in version 5.0.
// Returns the difference between two dates.
// The $dateDiff expression has this syntax:
//
//	{
//	  $dateDiff: {
//	     startDate: <Expression>,
//	     endDate: <Expression>,
//	     unit: <Expression>,
//	     timezone: <tzExpression>,
//	     startOfWeek: <String>
//	  }
//	}
//
// Subtracts startDate from endDate. Returns an integer in the specified unit.
func DateDiff[T expression.DateExpression, N expression.NumberExpression](startDate, endDate T, unit options.MongoTimeUnit, amount N, options DateDiffOptions) dateOperator {
	doc := bsonx.BsonDoc("startDate", startDate)
	doc.Append("endDate", endDate)
	doc.Append("unit", bsonx.String(unit.GetValue()))
	doc.Append("amount", amount)
	if options.timezone != "" {
		doc.Append("timezone", bsonx.String(options.timezone))
	}
	if options.startOfWeek != "" {
		doc.Append("startOfWeek", bsonx.String(options.startOfWeek.String()))
	}
	return dateOperator{doc: filter.NewSimpleFilter("$dateDiff", doc)}
}

// DateFromParts Constructs and returns a Date object given the date's constituent properties.
//
// The $dateFromParts expression has the following syntax:
//
//	{
//	   $dateFromParts : {
//	       'year': <year>, 'month': <month>, 'day': <day>,
//	       'hour': <hour>, 'minute': <minute>, 'second': <second>,
//	       'millisecond': <ms>, 'timezone': <tzExpression>
//	   }
//	}
func DateFromParts(year int32, options DateFromPartsOptions) dateOperator {
	doc := bsonx.BsonDoc("year", bsonx.Int32(year))
	if options.month > 0 {
		doc.Append("month", bsonx.Int32(options.month))
	}
	if options.day > 0 {
		doc.Append("day", bsonx.Int32(options.day))
	}
	if options.hour > 0 {
		doc.Append("hour", bsonx.Int32(options.hour))
	}
	if options.minute > 0 {
		doc.Append("minute", bsonx.Int32(options.minute))
	}
	if options.second > 0 {
		doc.Append("second", bsonx.Int32(options.second))
	}
	if options.minute > 0 {
		doc.Append("millisecond", bsonx.Int32(options.millisecond))
	}
	if options.timezone != "" {
		doc.Append("timezone", bsonx.String(options.timezone))
	}
	return dateOperator{doc: filter.NewSimpleFilter("$dateFromParts", doc)}
}

// DateFromPartsIso Constructs and returns a Date object given the date's constituent properties.
// ISO week date format using the following syntax:
//
//	{
//	   $dateFromParts : {
//	       'isoWeekYear': <year>, 'isoWeek': <week>, 'isoDayOfWeek': <day>,
//	       'hour': <hour>, 'minute': <minute>, 'second': <second>,
//	       'millisecond': <ms>, 'timezone': <tzExpression>
//	   }
//	}
func DateFromPartsIso(isoWeekYear int32, options DateFromPartsOptions) dateOperator {
	doc := bsonx.BsonDoc("isoWeekYear", bsonx.Int32(isoWeekYear))
	if options.isoWeek > 0 {
		doc.Append("isoWeek", bsonx.Int32(options.isoWeek))
	}
	if options.isoDayOfWeek > 0 {
		doc.Append("isoDayOfWeek", bsonx.Int32(options.isoDayOfWeek))
	}
	if options.hour > 0 {
		doc.Append("hour", bsonx.Int32(options.hour))
	}
	if options.minute > 0 {
		doc.Append("minute", bsonx.Int32(options.minute))
	}
	if options.second > 0 {
		doc.Append("second", bsonx.Int32(options.second))
	}
	if options.minute > 0 {
		doc.Append("millisecond", bsonx.Int32(options.millisecond))
	}
	if options.timezone != "" {
		doc.Append("timezone", bsonx.String(options.timezone))
	}
	return dateOperator{doc: filter.NewSimpleFilter("$dateFromParts", doc)}
}

// DateFromString Converts a date/time string to a date object.
//
// The $dateFromString expression has the following syntax:
//
//	{ $dateFromString: {
//	    dateString: <dateStringExpression>,
//	    format: <formatStringExpression>,
//	    timezone: <tzExpression>,
//	    onError: <onErrorExpression>,
//	    onNull: <onNullExpression>
//	} }
func DateFromString(dateString string, options DateFromStringOptions) dateOperator {
	doc := bsonx.BsonDoc("dateString", bsonx.String(dateString))
	if options.format != "" {
		doc.Append("format", bsonx.String(options.format))
	}
	if options.timezone != "" {
		doc.Append("timezone", bsonx.String(options.timezone))
	}
	if options.onError != nil {
		doc.Append("onError", options.onError.BsonDocument())
	}
	if options.onNull != nil {
		doc.Append("onNull", options.onNull.BsonDocument())
	}
	return dateOperator{doc: filter.NewSimpleFilter("$dateFromString", doc)}
}

// DateSubtract NewDefaultStage in version 5.0.
// Decrements a Date object by a specified number of time units.
// The $dateSubtract expression has the following syntax:
//
//	{
//	  $dateSubtract: {
//	     startDate: <Expression>,
//	     unit: <Expression>,
//	     amount: <Expression>,
//	     timezone: <tzExpression>
//	  }
//	}
//
// Returns a Date. The startDate can be any expression that resolves to type Date, Timestamp or ObjectId. No matter which data type is used as input, the value returned will be a Date object.
func DateSubtract[T expression.DateExpression, N expression.NumberExpression](startDate T, unit options.MongoTimeUnit, amount N, options DateOptions) dateOperator {
	doc := bsonx.BsonDoc("startDate", startDate)
	doc.Append("unit", bsonx.String(unit.GetValue()))
	doc.Append("amount", amount)
	if options.timezone != "" {
		doc.Append("timezone", bsonx.String(options.timezone))
	}
	return dateOperator{doc: filter.NewSimpleFilter("$dateSubtract", doc)}
}

// DateToParts Returns a document that contains the constituent parts of a given BSON Date value as individual properties. The properties returned are year, month, day, hour, minute, second and millisecond.
// You can set the iso8601 property to true to return the parts representing an
// ISO week date
// instead. This will return a document where the properties are isoWeekYear, isoWeek, isoDayOfWeek, hour, minute, second and millisecond.
// The $dateToParts expression has the following syntax:
//
//	{
//	   $dateToParts: {
//	       'date' : <dateExpression>,
//	       'timezone' : <timezone>,
//	       'iso8601' : <boolean>
//	   }
//	}
func DateToParts[T expression.DateExpression](date T, options DateToPartsOptions) dateOperator {
	doc := bsonx.BsonDoc("date", date)
	doc.Append("iso8601", bsonx.Boolean(options.iso8601))
	if options.timezone != "" {
		doc.Append("timezone", bsonx.String(options.timezone))
	}
	return dateOperator{doc: filter.NewSimpleFilter("$dateSubtract", doc)}
}

type DateToStringOptions struct {
	timezone string
	format   string
	onNull   bsonx.Bson
}

// DateToString Converts a date object to a string according to a user-specified format.
// The $dateToString expression has the following operator expression syntax:
//
//	{ $dateToString: {
//	   date: <dateExpression>,
//	   format: <formatString>,
//	   timezone: <tzExpression>,
//	   onNull: <expression>
//	} }
func DateToString[T expression.DateExpression](date T, options DateToStringOptions) dateOperator {
	doc := bsonx.BsonDoc("date", date)
	if options.format != "" {
		doc.Append("format", bsonx.String(options.format))
	}
	if options.timezone != "" {
		doc.Append("timezone", bsonx.String(options.timezone))
	}
	if options.onNull != nil {
		doc.Append("onNull", options.onNull.BsonDocument())
	}
	return dateOperator{doc: filter.NewSimpleFilter("$dateToString", doc)}
}

// DateTrunc NewDefaultStage in version 5.0.
// Truncates a date.
// $dateTrunc syntax:
//
//	{
//	  $dateTrunc: {
//	     date: <Expression>,
//	     unit: <Expression>,
//	     binSize: <Expression>,
//	     timezone: <tzExpression>,
//	     startOfWeek: <Expression>
//	  }
//	}
func DateTrunc[T expression.DateExpression](date T, unit options.MongoTimeUnit, options DateTruncOptions) dateOperator {
	doc := bsonx.BsonDoc("date", date)
	doc.Append("unit", bsonx.String(unit.GetValue()))
	if options.binSize > 0 {
		doc.Append("binSize", bsonx.Int64(options.binSize))
	}
	if options.timezone != "" {
		doc.Append("timezone", bsonx.String(options.timezone))
	}
	if options.startOfWeek != "" {
		doc.Append("startOfWeek", bsonx.String(options.startOfWeek.String()))
	}
	return dateOperator{doc: filter.NewSimpleFilter("$dateTrunc", doc)}
}

// DayOfMonth Returns the day of the month for a date as a number between 1 and 31.
// The $dayOfMonth expression has the following operator expression syntax:
// { $dayOfMonth: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func DayOfMonth[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$dayOfMonth", date, options)}
}

// DayOfWeek Returns the day of the week for a date as a number between 1 (Sunday) and 7 (Saturday).
// The $dayOfWeek expression has the following operator expression syntax:
// { $dayOfWeek: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func DayOfWeek[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$dayOfWeek", date, options)}
}

// DayOfYear Returns the day of the year for a date as a number between 1 and 366.
// The $dayOfYear expression has the following operator expression syntax:
// { $dayOfYear: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func DayOfYear[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$dayOfYear", date, options)}
}

// Hour Returns the hour portion of a date as a number between 0 and 23.
// The $hour expression has the following operator expression syntax:
// { $hour: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func Hour[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$hour", date, options)}
}

// IsoDayOfWeek Returns the weekday number in ISO 8601 format, ranging from 1 (for Monday) to 7 (for Sunday).
// The $isoDayOfWeek expression has the following operator expression syntax:
// { $isoDayOfWeek: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func IsoDayOfWeek[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$isoDayOfWeek", date, options)}
}

// IsoWeek Returns the week number in ISO 8601 format, ranging from 1 to 53. Week numbers start at 1 with the week (Monday through Sunday) that contains the year's first Thursday.
// The $isoWeek expression has the following operator expression syntax:
// { $isoWeek: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func IsoWeek[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$isoWeek", date, options)}
}

// IsoWeekYear Returns the year number in ISO 8601 format. The year starts with the Monday of week 1 and ends with the Sunday of the last week.
// The $isoWeekYear expression has the following operator expression syntax:
// { $isoWeekYear: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func IsoWeekYear[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$isoWeekYear", date, options)}
}

// Millisecond Returns the millisecond portion of a date as an integer between 0 and 999.
// The $millisecond expression has the following operator expression syntax:
// { $millisecond: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func Millisecond[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$millisecond", date, options)}
}

// Minute Returns the minute portion of a date as a number between 0 and 59.
// The $minute expression has the following operator expression syntax:
// { $minute: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func Minute[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$minute", date, options)}
}

// Month Returns the second portion of a date as a number between 0 and 59, but can be 60 to account for leap seconds.
// The $second expression has the following operator expression syntax:
// { $second: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func Month[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$month", date, options)}
}

// Second Returns the second portion of a date as a number between 0 and 59, but can be 60 to account for leap seconds.
// The $second expression has the following operator expression syntax:
// { $second: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func Second[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$second", date, options)}
}

// Week Returns the week of the year for a date as a number between 0 and 53.
// Weeks begin on Sundays, and week 1 begins with the first Sunday of the year. Days preceding the first Sunday of the year are in week 0. This behavior is the same as the "%U" operator to the strftime standard library function.
// The $week expression has the following operator expression syntax:
// { $week: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func Week[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$week", date, options)}
}

// Year Returns the year portion of a date.
// The $year expression has the following operator expression syntax:
// { $year: <dateExpression> }
// The argument can be:
// An expression that resolves to a Date, a Timestamp, or an ObjectID.
// A document with this format:
// { date: <dateExpression>, timezone: <tzExpression> }
func Year[T expression.DateExpression](date T, options DateOptions) dateOperator {
	return dateOperator{doc: NewSimpleFormatDate("$year", date, options)}
}

// ToDate Converts a value to a date. If the value cannot be converted to a date,
// $toDate errors. If the value is null or missing, $toDate returns null.
// $toDate  has the following syntax:
//
//	{
//	  $toDate: <expression>
//	}
//
// The $toDate takes any valid expression.
// The $toDate is a shorthand for the following $convert expression:
// { $convert: { input: <expression>, to: "date" } }
func ToDate[T expression.DateExpression](date T) dateOperator {
	return dateOperator{doc: bsonx.BsonDoc("$toDate", date)}
}

type SimpleFormatDate[T expression.DateExpression] struct {
	name    string
	date    T
	options DateOptions
}

func NewSimpleFormatDate[T expression.DateExpression](name string, date T, options DateOptions) SimpleFormatDate[T] {
	return SimpleFormatDate[T]{
		name:    name,
		date:    date,
		options: options,
	}
}

func (s SimpleFormatDate[T]) BsonDocument() *bsonx.BsonDocument {
	format := bsonx.Doc("date", s.date)
	if s.options.timezone != "" {
		format.Append("timezone", s.options.timezone)
	}
	return format.BsonDocument()
}

func (s SimpleFormatDate[T]) Document() bson.D {
	return s.BsonDocument().Document()
}
