package domain

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
)

type NlpRepository interface {
	AnalyzeSentence(ctx *appcontext.AppContext, sentence string) (*NlpSentenceAnalysisResult, error)
	TranslateDefinition(ctx *appcontext.AppContext, definition string) (*language.TranslatedLanguages, error)
}

type NlpSentenceAnalysisResult struct {
	Translated   language.TranslatedLanguages
	PosTags      []PosTag
	Sentiment    Sentiment
	Dependencies []Dependency
	Verbs        []Verb
}
