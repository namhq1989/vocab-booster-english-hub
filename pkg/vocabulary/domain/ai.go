package domain

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
)

type AIRepository interface {
	GetVocabularyData(ctx *appcontext.AppContext, term string) (*VocabularyData, error)
	GetVocabularyExamples(ctx *appcontext.AppContext, vocabularyID, term string) ([]VocabularyExample, error)
}
