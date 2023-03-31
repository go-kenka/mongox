package updates

import (
	"github.com/go-kenka/mongox/model/aggregates"
)

type UpdateStage interface {
	aggregates.Stage
	aggregates.FieldsStage | aggregates.ReplaceStage | aggregates.ProjectStage
}
