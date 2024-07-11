package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type NlpRepository interface {
	AnalyzeSentence(ctx *appcontext.AppContext, sentence string) (*NlpSentenceAnalysisResult, error)
	TranslateDefinition(ctx *appcontext.AppContext, definition string) (*language.TranslatedLanguages, error)
	EvaluateSentence(ctx *appcontext.AppContext, sentence, tense string, vocabulary []string) (*NlpSentenceEvaluationResult, error)
	GrammarCheck(ctx *appcontext.AppContext, sentence string) ([]SentenceGrammarError, error)
}

type NlpSentenceAnalysisResult struct {
	Translated   language.TranslatedLanguages
	PosTags      []PosTag
	Sentiment    Sentiment
	Dependencies []Dependency
	Verbs        []Verb
	Level        SentenceLevel
}

type NlpSentenceEvaluationResult struct {
	IsEnglish           bool
	IsVocabularyCorrect bool
	IsTenseCorrect      bool
	Sentiment           Sentiment
	Translated          language.TranslatedLanguages
	Clauses             []SentenceClause
}
