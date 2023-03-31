package geojson

import (
	"github.com/go-kenka/mongox/bsonx"
)

type MultiPoint struct {
	Geometry
	coordinates []Position
}

func NewMultiPoint(coordinates []Position) MultiPoint {
	return MultiPoint{
		coordinates: coordinates,
	}
}

func (p MultiPoint) GetType() GeoJsonObjectType {
	return GeoTypeMultiPoint
}

func (p MultiPoint) getCoordinates() []Position {
	return p.coordinates
}

func (p MultiPoint) GetCoordinateReferenceSystem() CoordinateReferenceSystem {
	return p.coordinateReferenceSystem
}

func (p MultiPoint) Encode() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	b.Append("type", bsonx.String(p.GetType().typeName))
	coordinates := bsonx.Array()
	for _, ps := range p.getCoordinates() {
		coordinates.Append(ps.Encode())
	}
	b.Append("coordinates", coordinates)
	return b
}
