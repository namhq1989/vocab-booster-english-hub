package mapping

import "github.com/namhq1989/vocab-booster-english-hub/core/language"

type SentenceClause struct {
	Clause         string `json:"clause"`
	Tense          string `json:"tense"`
	IsPrimaryTense bool   `json:"is_primary_tense"`
}

type SentenceGrammarError struct {
	Message     string                       `json:"message"`
	Translated  language.TranslatedLanguages `json:"translated"`
	Segment     string                       `json:"segment"`
	Replacement string                       `json:"replacement"`
}
