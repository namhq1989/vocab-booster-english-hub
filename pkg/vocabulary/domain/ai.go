package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type AIRepository interface {
	VocabularyExamples(ctx *appcontext.AppContext, vocabulary string, partsOfSpeech []string) ([]AIVocabularyExample, error)
	GrammarEvaluation(ctx *appcontext.AppContext, sentence string) ([]SentenceGrammarError, error)
}

type AIVocabularyExample struct {
	Example string
	Word    string
}
