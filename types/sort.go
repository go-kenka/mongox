package types

type SortType string

const (
	TypeAsc         SortType = "1"
	TypeDesc        SortType = "-1"
	Type2D          SortType = "2d"
	Type2dSphere    SortType = "2dsphere"
	TypeGeoHaystack SortType = "geoHaystack"
	TypeHashed      SortType = "hashed"
	TypeText        SortType = "text"
)

type MapEntry struct {
	Key   string
	Value SortType
}
