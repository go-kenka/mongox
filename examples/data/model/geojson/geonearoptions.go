package geojson

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
)

var (
	EmptyDoc              = bsonx.NewDoc()
	DefaultGeoNearOptions = GeoNearConstructibleBson{}.of(EmptyDoc)
)

type GeoNearOptions interface {
	bsonx.Bson
	DistanceMultiplier(distanceMultiplier int64) GeoNearOptions
	IncludeLocs(includeLocs string) GeoNearOptions
	Key(key string) GeoNearOptions
	MinDistance(minDistance int64) GeoNearOptions
	MaxDistance(maxDistance int64) GeoNearOptions
	Query(query bsonx.Document) GeoNearOptions
	Spherical() GeoNearOptions
}
