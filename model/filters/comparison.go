package filters

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/internal/expression"
	"go.mongodb.org/mongo-driver/bson"
)

type comparisonFilter struct {
	filter bsonx.Bson
}

func (f comparisonFilter) Value() bsonx.IBsonValue {
	return f.filter.ToBsonDocument()
}

func (f comparisonFilter) Document() bson.D {
	return f.filter.Document()
}

// Eq Matches values that are equal to a specified value.
func Eq[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: newSimpleFilter(fieldName, value)}
}

// Ne Matches all values that are not equal to a specified value.
func Ne[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: newOperatorFilter("$ne", fieldName, value)}
}

// Gt Matches values that are greater than a specified value.
func Gt[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: newOperatorFilter("$gt", fieldName, value)}
}

// Lt Matches values that are less than a specified value.
func Lt[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: newOperatorFilter("$lt", fieldName, value)}
}

// Gte Matches values that are greater than or equal to a specified value.
func Gte[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: newOperatorFilter("$gte", fieldName, value)}
}

// Lte Matches values that are less than or equal to a specified value.
func Lte[I expression.AnyExpression](fieldName string, value I) comparisonFilter {
	return comparisonFilter{filter: newOperatorFilter("$lte", fieldName, value)}
}

// In Matches any of the values specified in an array.
func In[I expression.AnyExpression](fieldName string, values ...I) comparisonFilter {
	return comparisonFilter{filter: newIterableOperatorFilter(fieldName, "$in", values)}
}

// Nin Matches none of the values specified in an array.
func Nin[I expression.AnyExpression](fieldName string, values ...I) comparisonFilter {
	return comparisonFilter{filter: newIterableOperatorFilter(fieldName, "$nin", values)}
}
