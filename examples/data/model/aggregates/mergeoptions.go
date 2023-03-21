package aggregates

import (
	"github.com/go-kenka/mongox/examples/data/model/expressions"
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
	variables           []Variable[expressions.TExpression]
	whenMatchedPipeline []bool
	whenNotMatched      WhenNotMatched
}
