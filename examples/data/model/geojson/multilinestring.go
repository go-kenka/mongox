package geojson

import "github.com/go-kenka/mongox/examples/data/bsonx"

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

func (p MultiLineString) Encode() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	b.Append("type", bsonx.NewBsonString(p.GetType().typeName))
	coordinates := bsonx.NewBsonArray()
	for _, ps := range p.getCoordinates() {
		cds := bsonx.NewBsonArray()
		for _, position := range ps {
			cds.Append(position.Encode())
		}
		coordinates.Append(cds)
	}
	b.Append("coordinates", coordinates)
	return b
}
