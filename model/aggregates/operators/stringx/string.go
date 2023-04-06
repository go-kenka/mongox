// Package strings String expressions, with the exception of $concat, only have a well-defined behavior for strings of ASCII characters.
package stringx

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type stringOperator struct {
	doc bsonx.Bson
}

func (o stringOperator) Exp() bsonx.IBsonValue {
	return o.doc.Pro()
}

// Concat Concatenates strings and returns the concatenated string. $concat has
// the following syntax: { $concat: [ <expression1>, <expression2>, ... ] } The
// arguments can be any valid expression as long as they resolve to strings. For
// more information on expressions, see Expressions. If the argument resolves to
// a value of null or refers to a field that is missing, $concat returns null.
func Concat[T expression.StrExpression](ns []T) stringOperator {
	data := bsonx.Array()
	for _, t := range ns {
		data.Append(t)
	}
	return stringOperator{doc: bsonx.BsonDoc("$concat", data)}
}

// IndexOfBytes Searches a string for an occurrence of a substring and returns
// the UTF-8 byte index (zero-based) of the first occurrence. If the substring is
// not found, returns -1. $indexOfBytes has the following operator expression
// syntax: { $indexOfBytes: [ <string expression>, <substring expression>,
// <start>, <end> ] }
func IndexOfBytes[T expression.StrExpression, N expression.IntExpression](str, sub T, start, end N) stringOperator {
	data := bsonx.Array()
	data.Append(str)
	data.Append(sub)
	if start != nil {
		data.Append(start)
	}
	if end != nil {
		data.Append(end)
	}
	return stringOperator{doc: bsonx.BsonDoc("$indexOfBytes", data)}
}

// IndexOfCP Searches a string for an occurrence of a substring and returns the
// UTF-8 code point index (zero-based) of the first occurrence. If the substring
// is not found, returns -1. $indexOfCP has the following operator expression
// syntax: { $indexOfCP: [ <string expression>, <substring expression>, <start>,
// <end> ] }
func IndexOfCP[T expression.StrExpression, N expression.IntExpression](str, sub T, start, end N) stringOperator {
	data := bsonx.Array()
	data.Append(str)
	data.Append(sub)
	if start != nil {
		data.Append(start)
	}
	if end != nil {
		data.Append(end)
	}
	return stringOperator{doc: bsonx.BsonDoc("$indexOfCP", data)}
}

// Ltrim Removes whitespace characters, including null, or the specified
// characters from the beginning of a string. $ltrim has the following syntax: {
// $ltrim: { input: <string>, chars: <string> } }
func Ltrim[T expression.StrExpression](input, chars T) stringOperator {
	data := bsonx.BsonDoc("input", input)
	data.Append("chars", chars)
	return stringOperator{doc: bsonx.BsonDoc("$ltrim", data)}
}

// Rtrim Removes whitespace characters, including null, or the specified
// characters from the end of a string. $rtrim has the following syntax: {
// $rtrim: { input: <string>, chars: <string> } }
func Rtrim[T expression.StrExpression](input, chars T) stringOperator {
	data := bsonx.BsonDoc("input", input)
	data.Append("chars", chars)
	return stringOperator{doc: bsonx.BsonDoc("$rtrim", data)}
}

// Trim Removes whitespace characters, including null, or the specified
// characters from the beginning and end of a string. $trim has the following
// syntax: { $trim: { input: <string>, chars: <string> } }
func Trim[T expression.StrExpression](input, chars T) stringOperator {
	data := bsonx.BsonDoc("input", input)
	data.Append("chars", chars)
	return stringOperator{doc: bsonx.BsonDoc("$trim", data)}
}

// RegexFind New in version 4.2. Provides regular expression (regex) pattern
// matching capability in aggregation expressions. If a match is found, returns a
// document that contains information on the first match. If a match is not
// found, returns null. MongoDB uses Perl compatible regular expressions (i.e.
// "PCRE" ) version 8.41 with UTF-8 support. Prior to MongoDB 4.2, aggregation
// pipeline can only use the query operator $regex in the $match stage. For more
// information on using regex in a query, see $regex.
func RegexFind[T expression.StrExpression](input, regex T, options string) stringOperator {
	data := bsonx.BsonDoc("input", input)
	data.Append("regex", regex)
	if options != "" {
		data.Append("options", bsonx.String(options))
	}
	return stringOperator{doc: bsonx.BsonDoc("$regexFind", data)}
}

// RegexFindAll New in version 4.2. Provides regular expression (regex) pattern
// matching capability in aggregation expressions. The operator returns an array
// of documents that contains information on each match. If a match is not found,
// returns an empty array. MongoDB uses Perl compatible regular expressions (i.e.
// "PCRE" ) version 8.41 with UTF-8 support. Prior to MongoDB 4.2, aggregation
// pipeline can only use the query operator $regex in the $match stage. For more
// information on using regex in a query, see $regex.
func RegexFindAll[T expression.StrExpression](input, regex T, options string) stringOperator {
	data := bsonx.BsonDoc("input", input)
	data.Append("regex", regex)
	if options != "" {
		data.Append("options", bsonx.String(options))
	}
	return stringOperator{doc: bsonx.BsonDoc("$regexFindAll", data)}
}

// RegexMatch New in version 4.2. Performs a regular expression (regex) pattern
// matching and returns: true if a match exists. false if a match doesn't exist.
// MongoDB uses Perl compatible regular expressions (i.e. "PCRE" ) version 8.41
// with UTF-8 support. Prior to MongoDB 4.2, aggregation pipeline can only use
// the query operator $regex in the $match stage. For more information on using
// regex in a query, see $regex.
func RegexMatch[T expression.StrExpression](input, regex T, options string) stringOperator {
	data := bsonx.BsonDoc("input", input)
	data.Append("regex", regex)
	if options != "" {
		data.Append("options", bsonx.String(options))
	}
	return stringOperator{doc: bsonx.BsonDoc("$regexMatch", data)}
}

// ReplaceOne New in version 4.4. Replaces the first instance of a search string
// in an input string with a replacement string. If no occurrences are found,
// $replaceOne evaluates to the input string. $replaceOne is both case-sensitive
// and diacritic-sensitive, and ignores any collation present on a collection.
// The $replaceOne operator has the following operator expression syntax: {
// $replaceOne: { input: <expression>, find: <expression>, replacement:
// <expression> } }
func ReplaceOne[T expression.StrExpression](input, find, replacement T) stringOperator {
	data := bsonx.BsonDoc("input", input)
	data.Append("find", find)
	data.Append("replacement", replacement)
	return stringOperator{doc: bsonx.BsonDoc("$replaceOne", data)}
}

// ReplaceAll New in version 4.4. Replaces all instances of a search string in an
// input string with a replacement string. $replaceAll is both case-sensitive and
// diacritic-sensitive, and ignores any collation present on a collection. The
// $replaceAll operator has the following operator expression syntax: {
// $replaceAll: { input: <expression>, find: <expression>, replacement:
// <expression> } }
func ReplaceAll[T expression.StrExpression](input, find, replacement T) stringOperator {
	data := bsonx.BsonDoc("input", input)
	data.Append("find", find)
	data.Append("replacement", replacement)
	return stringOperator{doc: bsonx.BsonDoc("$replaceAll", data)}
}

// Split Divides a string into an array of substrings based on a delimiter.
// $split removes the delimiter and returns the resulting substrings as elements
// of an array. If the delimiter is not found in the string, $split returns the
// original string as the only element of an array. $split has the following
// operator expression syntax: { $split: [ <string expression>, <delimiter> ] }
func Split[T expression.StrExpression](str, delimiter T) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$split", bsonx.Array(str, delimiter))}
}

// StrLenBytes Returns the number of UTF-8 encoded bytes in the specified string.
// $strLenBytes has the following operator expression syntax: { $strLenBytes:
// <string expression> } The argument can be any valid expression as long as it
// resolves to a string. For more information on expressions, see Expressions. If
// the argument resolves to a value of null or refers to a missing field,
// $strLenBytes returns an error.
func StrLenBytes[T expression.StrExpression](str T) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$strLenBytes", str)}
}

// StrLenCP Returns the number of UTF-8 code points in the specified string.
// $strLenCP has the following operator expression syntax: { $strLenCP: <string
// expression> } The argument can be any valid expression that resolves to a
// string. If the argument resolves to a value of null or refers to a missing
// field, $strLenCP returns an error.
func StrLenCP[T expression.StrExpression](str T) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$strLenCP", str)}
}

// StrCaseCmp Performs case-insensitive comparison of two strings. Returns 1 if
// first string is "greater than" the second string. 0 if the two strings are
// equal. -1 if the first string is "less than" the second string. $strcasecmp
// has the following syntax: { $strcasecmp: [ <expression1>, <expression2> ] }
// The arguments can be any valid expression as long as they resolve to strings.
func StrCaseCmp[T expression.StrExpression](e1, e2 T) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$strcasecmp", bsonx.Array(e1, e2))}
}

// Substr Deprecated since version 3.4: $substr is now an alias for $substrBytes.
// Returns a substring of a string, starting at a specified index position and
// including the specified number of characters. The index is zero-based. $substr
// has the following syntax: { $substr: [ <string>, <start>, <length> ] } The
// arguments can be any valid expression as long as the first argument resolves
// to a string, and the second and third arguments resolve to integers
func Substr[T expression.StrExpression, N expression.IntExpression](str T, start, length N) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$substr", bsonx.Array(str, start, length))}
}

// SubstrBytes Returns the substring of a string. The substring starts with the
// character at the specified UTF-8 byte index (zero-based) in the string and
// continues for the number of bytes specified. $substrBytes has the following
// operator expression syntax: { $substrBytes: [ <string expression>, <byte
// index>, <byte count> ] }
func SubstrBytes[T expression.StrExpression, N expression.NumberExpression](str T, index, count N) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$substrBytes", bsonx.Array(str, index, count))}
}

// SubstrCP SubstrCPReturns the substring of a string. The substring starts with
// the character at the specified UTF-8 code point (CP) index (zero-based) in the
// string for the number of code points specified. $substrCP has the following
// operator expression syntax: { $substrCP: [ <string expression>, <code point
// index>, <code point count> ]
func SubstrCP[T expression.StrExpression, N expression.NumberExpression](str T, index, count N) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$substrCP", bsonx.Array(str, index, count))}
}

// ToLower Converts a string to lowercase, returning the result. $toLower has the
// following syntax: { $toLower: <expression> } The argument can be any
// expression as long as it resolves to a string.
func ToLower[T expression.StrExpression](str T) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$toLower", str)}
}

// ToUpper Converts a string to uppercase, returning the result. $toUpper has the
// following syntax: { $toUpper: <expression> } The argument can be any
// expression as long as it resolves to a string
func ToUpper[T expression.StrExpression](str T) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$toUpper", str)}
}

// ToString Converts a value to a string. If the value cannot be converted to a
// string, $toString errors. If the value is null or missing, $toString returns
// null. $toString has the following syntax:
//
//	{
//	  $toString: <expression>
//	}
//
// The $toString takes any valid expression.
// The $toString is a shorthand for the following $convert expression:
// { $convert: { input: <expression>, to: "string" } }
func ToString[T expression.AnyExpression](str T) stringOperator {
	return stringOperator{doc: bsonx.BsonDoc("$toString", str)}
}
