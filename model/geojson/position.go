package geojson

import (
	"github.com/go-kenka/mongox/bsonx"
)

type Position struct {
	values []float64
}

func NewPosition(values []float64) Position {
	return Position{
		values: values,
	}
}

func (p Position) GetValues() []float64 {
	return p.values
}

func (p Position) Encode() *bsonx.BsonArray {
	coordinates := bsonx.Array()
	for _, v := range p.GetValues() {
		coordinates.Append(bsonx.Double(v))
	}
	return coordinates
}
