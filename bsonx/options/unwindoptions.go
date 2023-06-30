package options

type UnwindOptions struct {
	preserveNullAndEmptyArrays *bool
	includeArrayIndex          *string
}

func NewUnwindOptions(preserveNullAndEmptyArrays *bool, includeArrayIndex *string) UnwindOptions {
	return UnwindOptions{
		preserveNullAndEmptyArrays: preserveNullAndEmptyArrays,
		includeArrayIndex:          includeArrayIndex,
	}
}

func (u UnwindOptions) HasPreserveNullAndEmptyArrays() bool {
	return u.preserveNullAndEmptyArrays != nil
}
func (u UnwindOptions) PreserveNullAndEmptyArrays() bool {
	return *u.preserveNullAndEmptyArrays
}
func (u UnwindOptions) HasIncludeArrayIndex() bool {
	return u.includeArrayIndex != nil
}
func (u UnwindOptions) IncludeArrayIndex() string {
	return *u.includeArrayIndex
}
