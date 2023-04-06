package date

type StartOfWeek string

const (
	Monday    StartOfWeek = "monday"
	Tuesday   StartOfWeek = "tuesday"
	Wednesday StartOfWeek = "wednesday"
	Thursday  StartOfWeek = "thursday"
	Friday    StartOfWeek = "friday"
	Saturday  StartOfWeek = "saturday"
	Sunday    StartOfWeek = "sunday"
)

func (w StartOfWeek) String() string {
	return string(w)
}
