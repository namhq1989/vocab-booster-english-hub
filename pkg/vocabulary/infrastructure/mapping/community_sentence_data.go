package mapping

import "github.com/namhq1989/vocab-booster-utilities/language"

type SentenceClause struct {
	Clause         string `json:"clause"`
	Tense          string `json:"tense"`
	IsPrimaryTense bool   `json:"is_primary_tense"`
}

type SentenceGrammarError struct {
	Message     language.Multilingual `json:"message"`
	Segment     string                `json:"segment"`
	Replacement string                `json:"replacement"`
}
