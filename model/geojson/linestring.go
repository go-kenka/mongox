package geojson

import (
	"github.com/go-kenka/mongox/bsonx"
)

type LineString struct {
	Geometry
	coordinates []Position
}

func NewLineString(coordinates []Position) LineString {
	return LineString{
		coordinates: coordinates,
	}
}

func (p LineString) GetType() GeoJsonObjectType {
	return GeoTypeLineString
}

func (p LineString) GetCoordinates() []Position {
	return p.coordinates
}

func (p LineString) GetCoordinateReferenceSystem() CoordinateReferenceSystem {
	return p.coordinateReferenceSystem
}

func (p LineString) Encode() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	b.Append("type", bsonx.String(p.GetType().typeName))
	coordinates := bsonx.Array()
	for _, position := range p.GetCoordinates() {
		coordinates.Append(position.Encode())
	}
	b.Append("coordinates", coordinates)
	return b
}
