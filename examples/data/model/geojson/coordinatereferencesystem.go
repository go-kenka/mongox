package geojson

type CoordinateReferenceSystem interface {
	GetType() CoordinateReferenceSystemType
}
