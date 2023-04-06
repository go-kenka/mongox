package aggregates

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/geojson"
	"go.mongodb.org/mongo-driver/bson"
)

type GeoNearStage Stage

// GeoNear Outputs documents in order of nearest to farthest from a specified point. The
// $geoNear stage has the following prototype form: { $geoNear: { <geoNear
// options> } } The $geoNear operator accepts a document that contains the
// following $geoNear options. Specify all distances in the same units as those
// of the processed documents' coordinate system:
func GeoNear(near geojson.Point, distanceField string, options geojson.GeoNearOptions) GeoNearStage {
	return NewGeoNearStage(near, distanceField, options)
}

type geoNearStage struct {
	near          geojson.Point
	distanceField string
	options       geojson.GeoNearOptions
}

func (f geoNearStage) Bson() bsonx.Bson {
	return f.Pro()
}

func NewGeoNearStage(
	near geojson.Point,
	distanceField string,
	options geojson.GeoNearOptions,
) geoNearStage {
	return geoNearStage{
		near:          near,
		distanceField: distanceField,
		options:       options,
	}
}

func (f geoNearStage) Pro() *bsonx.BsonDocument {
	doc := bsonx.BsonEmpty()

	geoNear := bsonx.BsonEmpty()
	geoNear.Append("near", f.near.Encode())
	geoNear.Append("distanceField", bsonx.String(f.distanceField))

	for _, v := range f.options.Pro().Data() {
		geoNear.Append(v.Key, v.Value)
	}

	doc.Append("$geoNear", geoNear)
	return doc
}

func (f geoNearStage) Document() bson.D {
	return f.Pro().Document()
}
