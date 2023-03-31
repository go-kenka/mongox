package filters

import (
	"github.com/go-kenka/mongox/internal/expression"
)

type FilterExpression interface {
	expression.Expression
	logicalFilter | comparisonFilter | arrayFilter | bitwiseFilter | elementFilter | evaluationFilter | geoFilter
}
