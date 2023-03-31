package geojson

var (
	Epsg4326              = NewNamedCoordinateReferenceSystem("EPSG:4326")
	Crs84                 = NewNamedCoordinateReferenceSystem("urn:ogc:def:crs:OGC:1.3:CRS84")
	Epsg4326StrictWinding = NewNamedCoordinateReferenceSystem("urn:x-mongodb:crs:strictwinding:EPSG:4326")
)

type NamedCoordinateReferenceSystem struct {
	name string
}

func NewNamedCoordinateReferenceSystem(name string) NamedCoordinateReferenceSystem {
	return NamedCoordinateReferenceSystem{
		name: name,
	}
}

func (s NamedCoordinateReferenceSystem) GetType() CoordinateReferenceSystemType {
	return Name
}

func (s NamedCoordinateReferenceSystem) getName() string {
	return s.name
}
