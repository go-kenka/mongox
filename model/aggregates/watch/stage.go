package watch

import (
	"github.com/go-kenka/mongox/model/aggregates"
	"go.mongodb.org/mongo-driver/mongo"
)

type WatchPipe interface {
	Pipe() mongo.Pipeline
}

type watchPipe struct {
	stages []aggregates.WatchStage
}

func NewWatchPipe(s ...aggregates.WatchStage) *watchPipe {
	return &watchPipe{
		stages: s,
	}
}

func (p *watchPipe) Append(s ...aggregates.WatchStage) *watchPipe {
	p.stages = append(p.stages, s...)
	return p
}

func (p *watchPipe) Pipe() mongo.Pipeline {
	pipe := make(mongo.Pipeline, 0)
	for _, stage := range p.stages {
		pipe = append(pipe, stage.Document())
	}
	return pipe
}
