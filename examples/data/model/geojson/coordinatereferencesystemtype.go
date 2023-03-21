package geojson

var (
	Name = CoordinateReferenceSystemType{typeName: "name"}
	Link = CoordinateReferenceSystemType{typeName: "link"}
)

type CoordinateReferenceSystemType struct {
	typeName string
}

func (t CoordinateReferenceSystemType) GetType() string {
	return t.typeName
}
