package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/projections"
)

type ProjectStage Stage

// Project Passes along the documents with the requested fields to the next stage in the
// pipeline. The specified fields can be existing fields from the input
// documents or newly computed fields. The $project stage has the following
// prototype form: { $project: { <specification(s)> } } The $project takes a
// document that can specify the inclusion of fields, the suppression of the _id
// field, the addition of new fields, and the resetting of the values of
// existing fields. Alternatively, you may specify the exclusion of fields.
func Project(projection projections.Projections) ProjectStage {
	return NewStage(bsonx.BsonDoc("$project", projection.Pro()))
}
