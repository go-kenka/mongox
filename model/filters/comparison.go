package filters

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
)

type comparisonFilter struct {
	filter bsonx.Bson
}

func (f comparisonFilter) Exp() bsonx.IBsonValue {
	return f.filter.ToBsonDocument()
}

// Eq Matches values that are equal to a specified value.
func Eq[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: NewSimpleFilter(fieldName, value)}
}

// Ne Matches all values that are not equal to a specified value.
func Ne[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: NewOperatorFilter("$ne", fieldName, value)}
}

// Gt Matches values that are greater than a specified value.
func Gt[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: NewOperatorFilter("$gt", fieldName, value)}
}

// Lt Matches values that are less than a specified value.
func Lt[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: NewOperatorFilter("$lt", fieldName, value)}
}

// Gte Matches values that are greater than or equal to a specified value.
func Gte[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: NewOperatorFilter("$gte", fieldName, value)}
}

// Lte Matches values that are less than or equal to a specified value.
func Lte[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: NewOperatorFilter("$lte", fieldName, value)}
}

// In Matches any of the values specified in an array.
func In[I expression.AnyExpression](fieldName string, values ...I) comparisonFilter {
	return comparisonFilter{filter: NewIterableOperatorFilter(fieldName, "$in", values)}
}

// Nin Matches none of the values specified in an array.
func Nin[I expression.AnyExpression](fieldName string, values ...I) comparisonFilter {
	return comparisonFilter{filter: NewIterableOperatorFilter(fieldName, "$nin", values)}
}
