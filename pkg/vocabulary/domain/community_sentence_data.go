package domain

import (
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
)

type SentenceClause struct {
	Clause         string
	Tense          Tense
	IsPrimaryTense bool
}

func NewSentenceClause(clause, tense string, isPrimaryTense bool) (*SentenceClause, error) {
	dTense := ToTense(tense)
	if !dTense.IsValid() {
		return nil, apperrors.Common.InvalidTense
	}

	return &SentenceClause{
		Clause:         clause,
		Tense:          dTense,
		IsPrimaryTense: isPrimaryTense,
	}, nil
}

type SentenceGrammarError struct {
	Message     string
	Translated  language.TranslatedLanguages
	Segment     string
	Replacement string
}

func NewSentenceGrammarError(message, segment, replacement string, translated language.TranslatedLanguages) (*SentenceGrammarError, error) {
	return &SentenceGrammarError{
		Message:     message,
		Segment:     segment,
		Replacement: replacement,
		Translated:  translated,
	}, nil
}
