package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type ExternalApiRepository interface {
	SearchTermWithDatamuse(ctx *appcontext.AppContext, term string) (*DatamuseSearchTermResult, error)
}

type DatamuseSearchTermResult struct {
	Definitions   []DatamuseTermDefinition
	Frequency     float64
	Ipa           string
	PartsOfSpeech []string
	Synonyms      []string
	Antonyms      []string
}

type DatamuseTermDefinition struct {
	Pos        string
	Definition language.Multilingual
}
