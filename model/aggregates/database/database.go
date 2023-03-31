package database

import "github.com/go-kenka/mongox/model/aggregates"

type DatabaseStage interface {
	aggregates.Stage
	aggregates.DocumentsStage
}
