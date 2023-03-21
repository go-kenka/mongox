package geojson

import "github.com/go-kenka/mongox/examples/data/bsonx"

type Polygon struct {
	Geometry
	coordinates PolygonCoordinates
}

func NewPolygon(exterior []Position, holes ...[]Position) Polygon {
	return Polygon{
		coordinates: NewPolygonCoordinates(
			exterior, holes,
		),
	}
}

func (p Polygon) GetType() GeoJsonObjectType {
	return GeoTypePolygon
}

func (p Polygon) getCoordinates() PolygonCoordinates {
	return p.coordinates
}

func (p Polygon) getExterior() []Position {
	return p.coordinates.GetExterior()
}

func (p Polygon) getHoles() [][]Position {
	return p.coordinates.GetHoles()
}

func (p Polygon) GetCoordinateReferenceSystem() CoordinateReferenceSystem {
	return p.coordinateReferenceSystem
}

func (p Polygon) Encode() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	b.Append("type", bsonx.NewBsonString(p.GetType().typeName))
	b.Append("coordinates", p.getCoordinates().Encode())
	return b
}
