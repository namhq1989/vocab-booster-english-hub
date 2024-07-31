package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type Service interface {
	SearchVocabulary(ctx *appcontext.AppContext, performerID, term string) (*Vocabulary, []string, error)
	FindVocabulary(ctx *appcontext.AppContext, term string) (*Vocabulary, error)
	FindVocabularyExamples(ctx *appcontext.AppContext, vocabularyID string) ([]VocabularyExample, error)
}
