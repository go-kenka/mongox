package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/projections"
	"go.mongodb.org/mongo-driver/bson"
)

type ProjectStage struct {
	stage bsonx.Bson
}

func (s ProjectStage) Bson() bsonx.Bson {
	return s.stage
}

func (s ProjectStage) Document() bson.D {
	return s.stage.Document()
}

func (s ProjectStage) Watch() {
}

func (s ProjectStage) Update() {
}

// Project Passes along the documents with the requested fields to the next DefaultStage in the
// pipeline. The specified fields can be existing fields from the input
// documents or newly computed fields. The $project DefaultStage has the following
// prototype form: { $project: { <specification(s)> } } The $project takes a
// document that can specify the inclusion of fields, the suppression of the _id
// field, the addition of new fields, and the resetting of the values of
// existing fields. Alternatively, you may specify the exclusion of fields.
func Project(projection projections.Projections) ProjectStage {
	return ProjectStage{stage: bsonx.BsonDoc("$project", projection.Pro())}
}
