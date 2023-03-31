package geojson

import (
	"github.com/go-kenka/mongox/bsonx"
)

type IGeometry interface {
	GetType() GeoJsonObjectType
	GetCoordinateReferenceSystem() CoordinateReferenceSystem
	Encode() *bsonx.BsonDocument
}

type Geometry struct {
	coordinateReferenceSystem CoordinateReferenceSystem
}

func NewGeometry(coordinateReferenceSystem CoordinateReferenceSystem) Geometry {
	return Geometry{
		coordinateReferenceSystem: coordinateReferenceSystem,
	}
}

func (g Geometry) GetType() GeoJsonObjectType {
	return GeoTypeInvalid
}

func (g Geometry) GetCoordinateReferenceSystem() CoordinateReferenceSystem {
	return g.coordinateReferenceSystem
}

func (g Geometry) Encode() *bsonx.BsonDocument {
	return bsonx.BsonEmpty()
}
