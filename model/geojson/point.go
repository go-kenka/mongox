package geojson

import (
	"github.com/go-kenka/mongox/bsonx"
)

type Point struct {
	Geometry
	coordinate Position
}

func NewPoint(coordinate Position) Point {
	return Point{
		coordinate: coordinate,
	}
}

func (p Point) GetType() GeoJsonObjectType {
	return GeoTypePoint
}

func (p Point) getCoordinates() Position {
	return p.coordinate
}

func (p Point) getPosition() Position {
	return p.coordinate
}

func (p Point) GetCoordinateReferenceSystem() CoordinateReferenceSystem {
	return p.coordinateReferenceSystem
}

func (p Point) Encode() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	b.Append("type", bsonx.String(p.GetType().typeName))
	coordinates := p.getPosition().Encode()
	b.Append("coordinates", coordinates)
	return b
}
