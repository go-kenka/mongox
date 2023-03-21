package geojson

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
)

type GeoNearConstructibleBson struct {
	base     bsonx.Bson
	appended bsonx.Document
}

func NewGeoNearConstructibleBson(base bsonx.Bson, appended bsonx.Document) GeoNearConstructibleBson {
	a := GeoNearConstructibleBson{
		base:     base,
		appended: EmptyDoc,
	}
	if !appended.IsEmpty() {
		a.appended = appended
	}
	return a
}

func (a GeoNearConstructibleBson) ToBsonDocument() bsonx.BsonDocument {
	baseDoc := a.base.ToBsonDocument()
	if baseDoc.IsEmpty() && a.appended.IsEmpty() {
		return bsonx.NewEmptyDoc()
	}

	if a.appended.IsEmpty() {
		return baseDoc
	}

	return bsonx.NewMerged(baseDoc, a.appended.ToBsonDocument())
}

func (a GeoNearConstructibleBson) Document() bsonx.Document {
	return a.ToBsonDocument().Document()
}

func (a GeoNearConstructibleBson) newAppended(name string, value any) GeoNearConstructibleBson {
	return a.newMutated(bsonx.NewDocument(name, value))
}

func (a GeoNearConstructibleBson) newMutated(d bsonx.Document) GeoNearConstructibleBson {
	newAppended := bsonx.NewDoc()
	for k, v := range a.appended {
		newAppended.Append(k, v)
	}
	for k, v := range d {
		newAppended.Append(k, v)
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
	doc := a.base.ToBsonDocument()
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
