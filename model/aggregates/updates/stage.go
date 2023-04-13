package updates

import (
	"github.com/go-kenka/mongox/model/aggregates"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdatePipe interface {
	Document() any
	Update()
}

type updatePipe struct {
	stages []aggregates.UpdateStage
}

func NewUpdatePipe(stages ...aggregates.UpdateStage) *updatePipe {
	return &updatePipe{
		stages: stages,
	}
}

func (p *updatePipe) Append(stages ...aggregates.UpdateStage) *updatePipe {
	p.stages = append(p.stages, stages...)
	return p
}

func (p *updatePipe) Document() any {
	pipe := make(mongo.Pipeline, 0)
	for _, stage := range p.stages {
		pipe = append(pipe, stage.Document())
	}
	return pipe
}

func (p *updatePipe) Update() {
}
