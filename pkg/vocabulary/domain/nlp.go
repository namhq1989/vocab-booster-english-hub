package domain

import "github.com/namhq1989/vocab-booster-english-hub/core/appcontext"

type NlpRepository interface {
	AnalyzeSentence(ctx *appcontext.AppContext, sentence string) (*NlpSentenceAnalysisResult, error)
}

type NlpSentenceAnalysisResult struct {
	Translated   string
	PosTags      []PosTag
	Sentiment    Sentiment
	Dependencies []Dependency
	Verbs        []Verb
}
