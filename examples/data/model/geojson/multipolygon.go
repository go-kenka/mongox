package geojson

import "github.com/go-kenka/mongox/examples/data/bsonx"

type MultiPolygon struct {
	Geometry
	coordinates []PolygonCoordinates
}

func NewMultiPolygon(coordinates []PolygonCoordinates) MultiPolygon {
	return MultiPolygon{
		coordinates: coordinates,
	}
}

func (p MultiPolygon) GetType() GeoJsonObjectType {
	return GeoTypeMultiPolygon
}

func (p MultiPolygon) getCoordinates() []PolygonCoordinates {
	return p.coordinates
}

func (p MultiPolygon) GetCoordinateReferenceSystem() CoordinateReferenceSystem {
	return p.coordinateReferenceSystem
}

func (p MultiPolygon) Encode() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	b.Append("type", bsonx.NewBsonString(p.GetType().typeName))
	coordinates := bsonx.NewBsonArray()
	for _, ps := range p.getCoordinates() {
		coordinates.Append(ps.Encode())
	}
	b.Append("coordinates", coordinates)
	return b
}
