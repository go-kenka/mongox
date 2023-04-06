package options

type TextSearchOptions struct {
	language           *string
	caseSensitive      *bool
	diacriticSensitive *bool
}

func NewTextSearchOptions() TextSearchOptions {
	return TextSearchOptions{}
}

func (t TextSearchOptions) Language(language string) TextSearchOptions {
	t.language = &language
	return t
}

func (t TextSearchOptions) HasLanguage() bool {
	return t.language != nil
}

func (t TextSearchOptions) GetLanguage() string {
	return *t.language
}

func (t TextSearchOptions) CaseSensitive(caseSensitive bool) TextSearchOptions {
	t.caseSensitive = &caseSensitive
	return t
}

func (t TextSearchOptions) HasCaseSensitive() bool {
	return t.caseSensitive != nil
}

func (t TextSearchOptions) GetCaseSensitive() bool {
	return *t.caseSensitive
}

func (t TextSearchOptions) DiacriticSensitive(diacriticSensitive bool) TextSearchOptions {
	t.diacriticSensitive = &diacriticSensitive
	return t
}

func (t TextSearchOptions) HasDiacriticSensitive() bool {
	return t.diacriticSensitive != nil
}

func (t TextSearchOptions) GetDiacriticSensitive() bool {
	return *t.diacriticSensitive
}
