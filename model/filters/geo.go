package filters

import (
	"github.com/go-kenka/mongox/bsonx"
	"github.com/go-kenka/mongox/model/geojson"
	"go.mongodb.org/mongo-driver/bson"
)

type geoFilter struct {
	filter bsonx.Bson
}

func (f geoFilter) Value() bsonx.IBsonValue {
	return f.filter.ToBsonDocument()
}

func (f geoFilter) Document() bson.D {
	return f.filter.Document()
}

// GeoWithin Selects documents with geospatial data that exists entirely within a specified shape.
// The specified shape can be either a GeoJSON Polygon (either single-ringed or multi-ringed), a GeoJSON MultiPolygon, or a shape defined by legacy coordinate pairs. The
// $geoWithin
// operator uses the $geometry operator to specify the GeoJSON object.
func GeoWithin(fieldName string, geometry geojson.IGeometry) geoFilter {
	return geoFilter{filter: newGeometryOperatorFilter("$geoWithin", fieldName,
		geometry.Encode(), 0, 0)}
}

// GeoWithinBson Selects documents with geospatial data that exists entirely within a specified shape.
// The specified shape can be either a GeoJSON Polygon (either single-ringed or multi-ringed), a GeoJSON MultiPolygon, or a shape defined by legacy coordinate pairs. The
// $geoWithin
// operator uses the $geometry operator to specify the GeoJSON object.
func GeoWithinBson(fieldName string, geometry bsonx.Bson) geoFilter {
	return geoFilter{filter: newGeometryOperatorFilter("$geoWithin", fieldName,
		geometry.ToBsonDocument(), 0, 0)}
}

// GeoWithinBox $geoWithin The available shape operators are $box
func GeoWithinBox(fieldName string, lowerLeftX, lowerLeftY, upperRightX, upperRightY float64) geoFilter {
	box := bsonx.BsonDoc("$box", bsonx.Array(
		bsonx.Array(
			bsonx.Double(lowerLeftX),
			bsonx.Double(lowerLeftY),
		),
		bsonx.Array(
			bsonx.Double(upperRightX),
			bsonx.Double(upperRightY),
		)))

	return geoFilter{filter: newGeometryOperatorFilter("$geoWithin", fieldName,
		box, 0, 0)}
}

// GeoWithinPolygon $geoWithin The available shape operators are $polygon
func GeoWithinPolygon(fieldName string, points [][]float64) geoFilter {
	pointsArray := bsonx.Array()
	for _, point := range points {
		pointsArray.Append(bsonx.Array(bsonx.Double(point[0]), bsonx.Double(point[1])))
	}
	polygon := bsonx.BsonDoc("$polygon", pointsArray)
	return geoFilter{filter: newOperatorFilter("$geoWithin", fieldName, polygon)}
}

// GeoWithinCenter $geoWithin The available shape operators are $center
func GeoWithinCenter(fieldName string, x, y, radius float64) geoFilter {
	center := bsonx.BsonDoc("$center", bsonx.Array(
		bsonx.Array(bsonx.Double(x), bsonx.Double(y)),
		bsonx.Double(radius),
	))
	return geoFilter{filter: newOperatorFilter("$geoWithin", fieldName, center)}
}

// GeoWithinCenterSphere $geoWithin The available shape operators are $centerSphere
func GeoWithinCenterSphere(fieldName string, x, y, radius float64) geoFilter {
	center := bsonx.BsonDoc("$centerSphere", bsonx.Array(
		bsonx.Array(bsonx.Double(x), bsonx.Double(y)),
		bsonx.Double(radius),
	))
	return geoFilter{filter: newOperatorFilter("$geoWithin", fieldName, center)}
}

// GeoIntersects Selects documents whose geospatial data intersects with a specified GeoJSON
// object; i.e. where the intersection of the data and the specified object is
// non-empty. The $geoIntersects operator uses the $geometry operator to specify
// the GeoJSON object. To specify a GeoJSON polygons or multipolygons using the
// default coordinate reference system (CRS),
func GeoIntersects(fieldName string, geometry geojson.IGeometry) geoFilter {
	return geoFilter{filter: newGeometryOperatorFilter("$geoIntersects", fieldName,
		geometry.Encode(), 0, 0)}
}

// GeoIntersectsBson Selects documents whose geospatial data intersects with a specified GeoJSON
// object; i.e. where the intersection of the data and the specified object is
// non-empty. The $geoIntersects operator uses the $geometry operator to specify
// the GeoJSON object. To specify a GeoJSON polygons or multipolygons using the
// default coordinate reference system (CRS),
func GeoIntersectsBson(fieldName string, geometry bsonx.Bson) geoFilter {
	return geoFilter{filter: newGeometryOperatorFilter("$geoIntersects", fieldName,
		geometry.ToBsonDocument(), 0, 0)}
}

// Near Specifies a point for which a geospatial query returns the documents
// from nearest to farthest. The $near operator can specify either a GeoJSON
// point or legacy coordinate point.
func Near(fieldName string, geometry geojson.Point, maxDistance, minDistance float64) geoFilter {
	return geoFilter{filter: newGeometryOperatorFilter("$near", fieldName, geometry.Encode(), maxDistance, minDistance)}
}

// NearBson Specifies a point for which a geospatial query returns the documents
// from nearest to farthest. The $near operator can specify either a GeoJSON
// point or legacy coordinate point.
func NearBson(fieldName string, geometry bsonx.Bson, maxDistance, minDistance float64) geoFilter {
	return geoFilter{filter: newGeometryOperatorFilter("$near", fieldName, geometry.ToBsonDocument(), maxDistance, minDistance)}
}

// NearCoordinate Specifies a point for which a geospatial query returns the documents
// from nearest to farthest. The $near operator can specify either a GeoJSON
// point or legacy coordinate point.
func NearCoordinate(fieldName string, x, y float64, maxDistance, minDistance float64) geoFilter {
	return geoFilter{filter: createNearFilterDocument(fieldName, x, y, maxDistance, minDistance, "$near")}
}

// NearSphere Specifies a point for which a geospatial query returns the
// documents from nearest to farthest. MongoDB calculates distances for
// $nearSphere using spherical geometry.
func NearSphere(fieldName string, geometry geojson.Point, maxDistance, minDistance float64) geoFilter {
	return geoFilter{filter: newGeometryOperatorFilter("$nearSphere", fieldName, geometry.Encode(), maxDistance, minDistance)}
}

// NearSphereBson Specifies a point for which a geospatial query returns the
// documents from nearest to farthest. MongoDB calculates distances for
// $nearSphere using spherical geometry.
func NearSphereBson(fieldName string, geometry bsonx.Bson, maxDistance, minDistance float64) geoFilter {
	return geoFilter{filter: newGeometryOperatorFilter("$nearSphere", fieldName, geometry.ToBsonDocument(), maxDistance, minDistance)}
}

// NearSphereCoordinate Specifies a point for which a geospatial query returns the
// documents from nearest to farthest. MongoDB calculates distances for
// $nearSphere using spherical geometry.
func NearSphereCoordinate(fieldName string, x, y float64, maxDistance, minDistance float64) geoFilter {
	return geoFilter{filter: createNearFilterDocument(fieldName, x, y, maxDistance, minDistance, "$nearSphere")}
}

func createNearFilterDocument(fieldName string, x, y float64, maxDistance, minDistance float64, operator string) bsonx.Bson {
	nearFilter := bsonx.BsonDoc(operator, bsonx.Array(bsonx.Double(x), bsonx.Double(y)))
	if maxDistance > 0 {
		nearFilter.Append("$maxDistance", bsonx.Double(maxDistance))
	}
	if minDistance > 0 {
		nearFilter.Append("$minDistance", bsonx.Double(minDistance))
	}
	return bsonx.BsonDoc(fieldName, nearFilter)
}
