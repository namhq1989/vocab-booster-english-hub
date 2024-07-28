package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type Service interface {
	FindVocabulary(ctx *appcontext.AppContext, term string) (*Vocabulary, error)
	FindVocabularyExamples(ctx *appcontext.AppContext, vocabularyID string) ([]VocabularyExample, error)
}
