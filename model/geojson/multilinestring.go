package geojson

import (
	"github.com/go-kenka/mongox/bsonx"
)

type MultiLineString struct {
	Geometry
	coordinates [][]Position
}

func NewMultiLineString(coordinates [][]Position) MultiLineString {
	return MultiLineString{
		coordinates: coordinates,
	}
}

func (p MultiLineString) GetType() GeoJsonObjectType {
	return GeoTypeMultiLineString
}

func (p MultiLineString) getCoordinates() [][]Position {
	return p.coordinates
}

func (p MultiLineString) GetCoordinateReferenceSystem() CoordinateReferenceSystem {
	return p.coordinateReferenceSystem
}

func (p MultiLineString) Encode() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	b.Append("type", bsonx.String(p.GetType().typeName))
	coordinates := bsonx.Array()
	for _, ps := range p.getCoordinates() {
		cds := bsonx.Array()
		for _, position := range ps {
			cds.Append(position.Encode())
		}
		coordinates.Append(cds)
	}
	b.Append("coordinates", coordinates)
	return b
}
