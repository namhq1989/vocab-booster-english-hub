package domain

import (
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-utilities/language"
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
	Message     language.Multilingual
	Segment     string
	Replacement string
}

func NewSentenceGrammarError(message language.Multilingual, segment, replacement string) (*SentenceGrammarError, error) {
	return &SentenceGrammarError{
		Message:     message,
		Segment:     segment,
		Replacement: replacement,
	}, nil
}
