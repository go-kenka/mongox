package geojson

import (
	"github.com/go-kenka/mongox/bsonx"
)

type PolygonCoordinates struct {
	exterior []Position
	holes    [][]Position
}

func NewPolygonCoordinates(exterior []Position, holes [][]Position) PolygonCoordinates {
	return PolygonCoordinates{
		exterior: exterior,
		holes:    holes,
	}
}

func (p PolygonCoordinates) GetExterior() []Position {
	return p.exterior
}

func (p PolygonCoordinates) GetHoles() [][]Position {
	return p.holes
}
func (p PolygonCoordinates) Encode() *bsonx.BsonArray {
	a := bsonx.Array()
	exterior := bsonx.Array()
	for _, position := range p.GetExterior() {
		exterior.Append(position.Encode())
	}

	a.Append(exterior)
	holes := bsonx.Array()
	for _, positions := range p.GetHoles() {
		ps := bsonx.Array()
		for _, position := range positions {
			ps.Append(position.Encode())
		}
		holes.Append(ps)
	}
	return a
}
