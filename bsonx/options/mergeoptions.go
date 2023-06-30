package options

import (
	"github.com/go-kenka/mongox/bsonx/expression"
	"github.com/go-kenka/mongox/model/aggregates"
)

type WhenMatched uint8

const (
	WhenMatchedInvalid WhenMatched = iota
	WhenMatchedReplace
	WhenMatchedKeepExisting
	WhenMatchedMerge
	WhenMatchedFail
	WhenMatchedPipeline
)

type WhenNotMatched uint8

const (
	WhenNotMatchedInvalid WhenNotMatched = iota
	WhenNotMatchedInsert
	WhenNotMatchedDiscard
	WhenNotMatchedFail
)

var (
	WhenMatcheds = [...]string{
		"invalid",
		"merge",
		"replace",
		"keepExisting",
		"fail",
		"pipeline",
	}

	WhenNotMatcheds = [...]string{
		"invalid",
		"insert",
		"discard",
		"fail",
	}
)

type MergeOptions struct {
	uniqueIdentifier    []string
	whenMatched         WhenMatched
	variables           []aggregates.Variable[expression.AnyExpression]
	whenMatchedPipeline []bool
	whenNotMatched      WhenNotMatched
}

func (o MergeOptions) UniqueIdentifier() []string {
	return o.uniqueIdentifier
}

func (o MergeOptions) WhenMatched() WhenMatched {
	return o.whenMatched
}

func (o MergeOptions) Variables() []aggregates.Variable[expression.AnyExpression] {
	return o.variables
}

func (o MergeOptions) WhenMatchedPipeline() []bool {
	return o.whenMatchedPipeline
}

func (o MergeOptions) WhenNotMatched() WhenNotMatched {
	return o.whenNotMatched
}
