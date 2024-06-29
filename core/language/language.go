package language

type TranslatedLanguages struct {
	Vi string `json:"vi"` // Vietnamese
}

const (
	Vietnamese = "vi"
)

func (l TranslatedLanguages) GetLanguageValue(lang string) string {
	lang = toLanguage(lang)

	if lang == "vi" {
		return l.Vi
	}
	return l.Vi
}

func toLanguage(lang string) string {
	if lang != "vi" {
		return "vi"
	}
	return lang
}
