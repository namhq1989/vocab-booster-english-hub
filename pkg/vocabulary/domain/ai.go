package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type AIRepository interface {
	VocabularyExamples(ctx *appcontext.AppContext, vocabulary string, partsOfSpeech []string) ([]AIVocabularyExample, error)
	GrammarEvaluation(ctx *appcontext.AppContext, sentence string) ([]SentenceGrammarError, error)
	WordOfTheDay(ctx *appcontext.AppContext, country, date string) (*AIWordOfTheDay, error)
}

type AIVocabularyExample struct {
	Example string
	Word    string
}

type AIWordOfTheDay struct {
	Word        string
	Information language.Multilingual
}
