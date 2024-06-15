package domain

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
)

type AIRepository interface {
	GetVocabularyData(ctx *appcontext.AppContext, vocabulary string) (*AIVocabularyData, error)
}

type AIVocabularyData struct {
	PosTags  []string
	IPA      string
	Synonyms []string
	Antonyms []string
	Examples []AIVocabularyExample
}

type AIVocabularyExample struct {
	Example    string
	Word       string
	Pos        string
	Definition string
}
