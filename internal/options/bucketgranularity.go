package options

type BucketGranularity string

const (
	R5        BucketGranularity = "R5"
	R10       BucketGranularity = "R10"
	R20       BucketGranularity = "R20"
	R40       BucketGranularity = "R40"
	R80       BucketGranularity = "R80"
	Series125 BucketGranularity = "1-2-5"
	E6        BucketGranularity = "E6"
	E12       BucketGranularity = "E12"
	E24       BucketGranularity = "E24"
	E48       BucketGranularity = "E48"
	E96       BucketGranularity = "E96"
	E192      BucketGranularity = "E192"
	Powersof2 BucketGranularity = "POWERSOF2"
)

func (g BucketGranularity) GetValue() string {
	return string(g)
}
