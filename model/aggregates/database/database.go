package database

import (
	"github.com/go-kenka/mongox/model/aggregates"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabasePipe interface {
	Pipe() mongo.Pipeline
}

type databasePipe struct {
	stages []aggregates.DatabaseStage
}

func NewDatabasePipe(stages ...aggregates.DatabaseStage) *databasePipe {
	return &databasePipe{
		stages: stages,
	}
}

func (p *databasePipe) Append(stages ...aggregates.DatabaseStage) *databasePipe {
	p.stages = append(p.stages, stages...)
	return p
}

func (p *databasePipe) Pipe() mongo.Pipeline {
	pipe := make(mongo.Pipeline, 0)
	for _, stage := range p.stages {
		pipe = append(pipe, stage.Document())
	}
	return pipe
}
