package geojson

import (
	"github.com/go-kenka/mongox/bsonx"
)

type GeometryCollection struct {
	Geometry
	geometries []Geometry
}

func NewGeometryCollection(geometries []Geometry) GeometryCollection {
	return GeometryCollection{
		geometries: geometries,
	}
}

func (p GeometryCollection) GetType() GeoJsonObjectType {
	return GeoTypeGeometryCollection
}

func (p GeometryCollection) GetGeometries() []Geometry {
	return p.geometries
}

func (p GeometryCollection) GetCoordinateReferenceSystem() CoordinateReferenceSystem {
	return p.coordinateReferenceSystem
}

func (p GeometryCollection) Encode() *bsonx.BsonDocument {
	b := bsonx.BsonEmpty()
	b.Append("type", bsonx.String(p.GetType().typeName))
	geometries := bsonx.Array()
	for _, v := range p.GetGeometries() {
		geometries.Append(v.Encode())
	}
	b.Append("geometries", geometries)
	return b
}
