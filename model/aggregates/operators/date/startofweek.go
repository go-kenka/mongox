package date

type StartOfWeek string

const (
	monday    StartOfWeek = "monday"
	tuesday   StartOfWeek = "tuesday"
	wednesday StartOfWeek = "wednesday"
	thursday  StartOfWeek = "thursday"
	friday    StartOfWeek = "friday"
	saturday  StartOfWeek = "saturday"
	sunday    StartOfWeek = "sunday"
)

func (w StartOfWeek) String() string {
	return string(w)
}
