package filters

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"github.com/go-kenka/mongox/internal/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type evaluationFilter struct {
	filter bsonx.Bson
}

func (f evaluationFilter) Value() bsonx.IBsonValue {
	return f.filter.Pro()
}

func (f evaluationFilter) Document() bson.D {
	return f.filter.Document()
}

// Expr Allows the use of aggregation expressions within the query language.
// $expr has the following syntax: { $expr: { <expression> } }
func Expr[I expression.AnyExpression](expression I) evaluationFilter {
	return evaluationFilter{filter: newSimpleEncodingFilter("$expr", expression)}
}

// JsonSchema The $jsonSchema operator matches documents that satisfy the
// specified JSON Schema. The $jsonSchema operator expression has the following
// syntax: { $jsonSchema: <JSON Schema object> } Where the JSON Schema object is
// formatted according to draft 4 of the JSON Schema standard { <keyword1>:
// <value1>, ... }
func JsonSchema(schema bsonx.Bson) evaluationFilter {
	return evaluationFilter{filter: newSimpleEncodingFilter("$jsonSchema", schema.Pro())}
}

// Mod Select documents where the value of a field divided by a divisor has the
// specified remainder (i.e. perform a modulo operation to select documents). To
// specify a $mod expression, use the following syntax: { field: { $mod: [
// divisor, remainder ] } }
func Mod(fieldName string, divisor, remainder int64) evaluationFilter {
	return evaluationFilter{filter: newOperatorFilter("$mod", fieldName, bsonx.Array(bsonx.Int64(divisor), bsonx.Int64(remainder)))}
}

// Regex Provides regular expression capabilities for pattern matching strings in
// queries. MongoDB uses Perl compatible regular expressions (i.e. "PCRE" )
// version 8.42 with UTF-8 support. To use $regex , use one of the following
// syntaxes: { <field>: { $regex: /pattern/, $options: '<options>' } } {
// <field>: { $regex: 'pattern', $options: '<options>' } } { <field>: { $regex:
// /pattern/<options> } } In MongoDB, you can also use regular expression
// objects (i.e. /pattern/) to specify regular expressions: { <field>:
// /pattern/<options> }
func Regex(fieldName string, pattern, options string) evaluationFilter {
	return evaluationFilter{filter: newSimpleFilter(fieldName, bsonx.RegularExpression(primitive.Regex{
		Pattern: pattern,
		Options: options,
	}))}
}

// Text performs a text search on the content of the fields indexed with a text
// index. A $text expression has the following syntax:
//
//	{
//	 $text:
//	   {
//	     $search: <string>,
//	     $language: <string>,
//	     $caseSensitive: <boolean>,
//	     $diacriticSensitive: <boolean>
//	   }
//	}
func Text(search string, textSearchOptions options.TextSearchOptions) evaluationFilter {
	return evaluationFilter{filter: newTextFilter(search, textSearchOptions)}
}

type whereFilter struct {
	filter bsonx.Bson
}

func (f whereFilter) Value() bsonx.IBsonValue {
	return f.filter.Pro()
}

func (f whereFilter) Document() bson.D {
	return f.filter.Document()
}

// Where Use the $where operator to pass either a string containing a JavaScript
// expression or a full JavaScript function to the query system. The $where
// provides greater flexibility, but requires that the database processes the
// JavaScript expression or function for each document in the collection.
// Reference the document in the JavaScript expression or function using either
// this or obj . { $where: <string|JavaScript Code> }
func Where(javaScriptExpression string) whereFilter {
	return whereFilter{filter: bsonx.BsonDoc("$where", bsonx.String(javaScriptExpression))}
}
