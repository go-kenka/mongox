package geojson

import (
	"github.com/go-kenka/mongox/bsonx"
	"go.mongodb.org/mongo-driver/bson"
)

type GeoNearConstructibleBson struct {
	base     bsonx.Bson
	appended *bsonx.Document
}

func NewGeoNearConstructibleBson(base bsonx.Bson, appended *bsonx.Document) GeoNearConstructibleBson {
	a := GeoNearConstructibleBson{
		base:     base,
		appended: EmptyDoc,
	}
	if !appended.IsEmpty() {
		a.appended = appended
	}
	return a
}

func (a GeoNearConstructibleBson) BsonDocument() *bsonx.BsonDocument {
	baseDoc := a.base.BsonDocument()
	if baseDoc.IsEmpty() && a.appended.IsEmpty() {
		return bsonx.BsonEmpty()
	}

	if a.appended.IsEmpty() {
		return baseDoc
	}

	return bsonx.NewMerged(baseDoc, a.appended.BsonDocument())
}

func (a GeoNearConstructibleBson) Document() bson.D {
	return a.BsonDocument().Document()
}

func (a GeoNearConstructibleBson) newAppended(name string, value any) GeoNearConstructibleBson {
	return a.newMutated(bsonx.Doc(name, value))
}

func (a GeoNearConstructibleBson) newMutated(d *bsonx.Document) GeoNearConstructibleBson {
	newAppended := bsonx.Empty()
	for _, v := range a.appended.Document() {
		newAppended.Append(v.Key, v.Value)
	}
	for _, v := range d.Document() {
		newAppended.Append(v.Key, v.Value)
	}

	return GeoNearConstructibleBson{
		base:     a.base,
		appended: newAppended,
	}
}
func (a GeoNearConstructibleBson) of(doc bsonx.Bson) GeoNearConstructibleBson {
	d, ok := doc.(GeoNearConstructibleBson)
	if ok {
		return d
	}
	return NewGeoNearConstructibleBson(doc, nil)
}
func (a GeoNearConstructibleBson) remove(key string) GeoNearConstructibleBson {
	doc := a.base.BsonDocument()
	doc.Remove(key)
	appended := a.appended
	appended.Remove(key)
	return GeoNearConstructibleBson{base: doc, appended: appended}
}

func (a GeoNearConstructibleBson) setOption(key string, value any) GeoNearOptions {
	return a.newAppended(key, value)
}
func (a GeoNearConstructibleBson) DistanceMultiplier(distanceMultiplier int64) GeoNearOptions {
	return a.setOption("distanceMultiplier", distanceMultiplier)
}
func (a GeoNearConstructibleBson) IncludeLocs(includeLocs string) GeoNearOptions {
	return a.setOption("includeLocs", includeLocs)
}
func (a GeoNearConstructibleBson) Key(key string) GeoNearOptions {
	return a.setOption("key", key)
}
func (a GeoNearConstructibleBson) MinDistance(minDistance int64) GeoNearOptions {
	return a.setOption("minDistance", minDistance)
}

func (a GeoNearConstructibleBson) MaxDistance(maxDistance int64) GeoNearOptions {
	return a.setOption("maxDistance", maxDistance)
}

func (a GeoNearConstructibleBson) Query(query bsonx.Document) GeoNearOptions {
	return a.setOption("query", query)
}

func (a GeoNearConstructibleBson) Spherical() GeoNearOptions {
	return a.setOption("spherical", true)
}
