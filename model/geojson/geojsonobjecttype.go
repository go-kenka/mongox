package geojson

var (
	GeoTypeInvalid            = NewGeoJsonObjectType("Invalid")
	GeoTypeGeometryCollection = NewGeoJsonObjectType("GeometryCollection")
	GeoTypeLineString         = NewGeoJsonObjectType("LineString")
	GeoTypeMultiLineString    = NewGeoJsonObjectType("MultiLineString")
	GeoTypeMultiPoint         = NewGeoJsonObjectType("MultiPoint")
	GeoTypeMultiPolygon       = NewGeoJsonObjectType("MultiPolygon")
	GeoTypePoint              = NewGeoJsonObjectType("Point")
	GeoTypePolygon            = NewGeoJsonObjectType("Polygon")
)

type GeoJsonObjectType struct {
	typeName string
}

func NewGeoJsonObjectType(typeName string) GeoJsonObjectType {
	return GeoJsonObjectType{
		typeName: typeName,
	}
}

func (t GeoJsonObjectType) GetType() string {
	return t.typeName
}
