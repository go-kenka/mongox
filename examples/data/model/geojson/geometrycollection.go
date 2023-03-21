package geojson

import "github.com/go-kenka/mongox/examples/data/bsonx"

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

func (p GeometryCollection) Encode() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	b.Append("type", bsonx.NewBsonString(p.GetType().typeName))
	geometries := bsonx.NewBsonArray()
	for _, v := range p.GetGeometries() {
		geometries.Append(v.Encode())
	}
	b.Append("geometries", geometries)
	return b
}
