package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ExternalApiRepository interface {
	SearchTermWithDatamuse(ctx *appcontext.AppContext, term string) (*DatamuseSearchTermResult, error)
}

type DatamuseSearchTermResult struct {
	Definitions   []VocabularyDefinition
	Frequency     float64
	Ipa           string
	PartsOfSpeech []string
	Synonyms      []string
	Antonyms      []string
}
